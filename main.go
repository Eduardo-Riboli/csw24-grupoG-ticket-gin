package main

import (
	"github.com/grupoG/csw24-grupoG-ticket-gin/configs"
	"github.com/grupoG/csw24-grupoG-ticket-gin/controllers"
	_ "github.com/grupoG/csw24-grupoG-ticket-gin/docs"
	"github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
	"github.com/grupoG/csw24-grupoG-ticket-gin/routes"
	"github.com/grupoG/csw24-grupoG-ticket-gin/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Configurar o banco de dados
	db := configs.SetupDatabase()

	// Inicializar camadas
	sampleRepository := repositories.NewSampleRepository(db)
	sampleService := services.NewSampleService(sampleRepository)
	sampleController := controllers.NewSampleController(sampleService)
	tenantRepositoru := repositories.NewTenantRepository(db)
	tenantService := services.NewTenantService(tenantRepositoru)
	tenantController := controllers.NewTenantController(tenantService)

	// Configurar as rotas
	router := routes.SetupRouter(sampleController, tenantController)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor
	router.Run(":8080")
}
