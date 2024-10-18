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

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor
	router.Run(":8080")
}
