package main

import (
	"log"

	"github.com/grupoG/csw24-grupoG-ticket-gin/configs"
	"github.com/grupoG/csw24-grupoG-ticket-gin/controllers"
	"github.com/grupoG/csw24-grupoG-ticket-gin/docs" // Adicione essa linha para importar o pacote docs
	"github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
	"github.com/grupoG/csw24-grupoG-ticket-gin/routes"
	"github.com/grupoG/csw24-grupoG-ticket-gin/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
    // stdout and stderr are sent to AWS CloudWatch Logs
    log.Printf("Gin cold start")

    // Configurar o banco de dados
    db := configs.SetupDatabase()

    // Inicializar camadas
    sampleRepository := repositories.NewSampleRepository(db)
    sampleService := services.NewSampleService(sampleRepository)
    sampleController := controllers.NewSampleController(sampleService)
    tenantRepository := repositories.NewTenantRepository(db)
    tenantService := services.NewTenantService(tenantRepository)
    tenantController := controllers.NewTenantController(tenantService)
    userRepository := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepository)
    userController := controllers.NewUserController(userService)
    eventRepository := repositories.NewEventRepository(db)
    eventService := services.NewEventService(eventRepository)
    eventController := controllers.NewEventController(eventService)
    ticketRepository := repositories.NewTicketRepository(db)
    ticketService := services.NewTicketService(ticketRepository)
    ticketController := controllers.NewTicketController(ticketService)
    transactionRepository := repositories.NewTransactionRepository(db)
    transactionService := services.NewTransactionService(transactionRepository)
    transactionController := controllers.NewTransactionController(transactionService)
    notificationPreferencesRepository := repositories.NewNotificationPreferencesRepository(db)
    notificationPreferencesService := services.NewNotificationPreferencesService(notificationPreferencesRepository)
    notificationPreferencesController := controllers.NewNotificationPreferencesController(notificationPreferencesService)

    // Configurar as rotas
    router := routes.SetupRouter(sampleController,
        tenantController,
        userController,
        eventController,
        ticketController,
        transactionController,
        notificationPreferencesController)

    docs.SwaggerInfo.BasePath = "/api" // Acessar a variável BasePath corretamente

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.POST("/tickets/sell", ticketController.SellTicket)

    ginLambda = ginadapter.NewV2(router)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
    return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
    lambda.Start(Handler)
}