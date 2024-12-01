provider "aws" {
  region = "us-east-1"
}

resource "aws_vpc" "main_vpc" {
  cidr_block = "10.0.0.0/16"

  enable_dns_support = true  
  enable_dns_hostnames = true 

  tags = {
    Name = "Main VPC"
  }
}

resource "aws_subnet" "subnet1" {
  vpc_id                  = aws_vpc.main_vpc.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "us-east-1a"
  map_public_ip_on_launch = true
}

resource "aws_subnet" "subnet2" {
  vpc_id                  = aws_vpc.main_vpc.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "us-east-1b"
  map_public_ip_on_launch = true
}

resource "aws_security_group" "ecs_security_group" {
  vpc_id      = aws_vpc.main_vpc.id
  name        = "ecs-security-group"
  description = "Allow access to ECS containers"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "rds_security_group" {
  vpc_id      = aws_vpc.main_vpc.id
  name        = "rds-security-group"
  description = "Allow database access"

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    security_groups = [aws_security_group.ecs_security_group.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "go_rds" {
  identifier              = "rds-t1"
  engine                  = "postgres"
  engine_version          = "16.3"  # Versão do PostgreSQL
  instance_class          = "db.t3.micro"
  db_name                 = "godb"
  username                = "postgres"
  password                = "minhasenha123"
  publicly_accessible     = true
  skip_final_snapshot     = true
  allocated_storage       = 20
  vpc_security_group_ids  = [aws_security_group.rds_security_group.id]
  db_subnet_group_name    = aws_db_subnet_group.db_subnet_group.name
  tags = {
    Name = "PostgresRDS"
  }
}

resource "aws_db_subnet_group" "db_subnet_group" {
  name        = "rds-db-subnet-group"
  subnet_ids  = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]
  description = "Subnet group for NestJS RDS database"

  tags = {
    Name = "RDS Subnet Group"
  }
}

resource "aws_ecs_cluster" "ecs_cluster" {
  name = "go-ecs-cluster"
}

resource "aws_ecs_task_definition" "task_definition" {
  family                   = "go-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"

  execution_role_arn = "arn:aws:iam::807496334534:role/LabRole" 

  container_definitions = jsonencode([
    {
      name = "go-container"
      image = "807496334534.dkr.ecr.us-east-1.amazonaws.com/t1-repository:b28ad98a3eeb5118ad7319a04a1e70ae533e0f1d"

      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      environment = [
        {
          name  = "DATABASE_URL"
          value = "postgresql://${aws_db_instance.go_rds.username}:${aws_db_instance.go_rds.password}@${aws_db_instance.go_rds.endpoint}/${aws_db_instance.go_rds.db_name}"
        }
      ]
    }
  ])
}

resource "aws_lb" "ecs_alb" {
  name               = "ecs-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.ecs_security_group.id]
  subnets            = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]

  enable_deletion_protection = false
  enable_http2               = true

  tags = {
    Name = "ECS ALB"
  }
}


resource "aws_lb_listener" "alb_listener" {
  load_balancer_arn = aws_lb.ecs_alb.arn
  port              = 8080
  protocol          = "HTTP"

  default_action {
    type = "forward"
    target_group_arn = aws_lb_target_group.ecs_target_group.arn
  }
}

resource "aws_lb_target_group" "ecs_target_group" {
  name        = "ecs-target-group"
  port        = 8080
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main_vpc.id
  target_type = "ip"

    health_check {
    path                = "/health"
    interval            = 30
    timeout             = 10
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main_vpc.id

  tags = {
    Name = "Main Internet Gateway"
  }
}


resource "aws_route_table" "main_route_table" {
  vpc_id = aws_vpc.main_vpc.id
}

resource "aws_route" "internet_route" {
  route_table_id         = aws_route_table.main_route_table.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.main.id
}

resource "aws_route_table_association" "subnet1_route_association" {
  subnet_id      = aws_subnet.subnet1.id
  route_table_id = aws_route_table.main_route_table.id
}

resource "aws_route_table_association" "subnet2_route_association" {
  subnet_id      = aws_subnet.subnet2.id
  route_table_id = aws_route_table.main_route_table.id
}

resource "aws_ecs_service" "ecs_service" {
  depends_on = [
    aws_lb_listener.alb_listener
  ]
  name            = "go-service"
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.task_definition.arn
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = [aws_subnet.subnet1.id]
    security_groups = [aws_security_group.ecs_security_group.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.ecs_target_group.arn
    container_name   = "go-container"
    container_port   = 8080
  }

  desired_count = 1
}

  output "ecs_cluster_name" {
    value = aws_ecs_cluster.ecs_cluster.name
  }

  output "alb_dns" {
    value = aws_lb.ecs_alb.dns_name
  }
  output "rds_endpoint" {
    value = aws_db_instance.go_rds.endpoint
  }

resource "aws_instance" "teste_t1" {
    ami = "ami-0866a3c8686eaeeba"
    instance_type = "t2.micro"

    vpc_security_group_ids = [aws_security_group.api_access.id]

    # key_name = aws_key_pair.my_key_pair.key_name
    tags = {
      Name = "ec2-t1"
    }

      user_data = <<-EOF
        #!/bin/bash
        sudo apt-get update -y
        sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
        curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
        sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
        sudo apt-get update -y
        sudo apt-get install -y docker-ce
        sudo usermod -aG docker ubuntu
        sudo systemctl enable docker
        sudo systemctl start docker
    EOF
}

resource "aws_ecr_repository" "myRepository" {
  name                 = "t1-repository"  # Nome do seu repositório
  image_tag_mutability = "MUTABLE"            # Controla se as tags de imagem podem ser alteradas
}


resource "aws_security_group" "api_access" {
  name        = "API-security-group-T1"
  description = "Security group para permitir SSH, HTTP e HTTPS"

  # Regra de entrada para SSH (porta 22)
  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Regra de entrada para HTTP (porta 80)
  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Regra de entrada para HTTPS (porta 443)
  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "api port"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "pg port"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

    ingress {
    description = "pgAdmin port"
    from_port   = 5050
    to_port     = 5050
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Regra de saída que permite todo o tráfego
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
