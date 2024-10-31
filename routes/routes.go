package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grupoG/csw24-grupoG-ticket-gin/controllers"
)

func SetupRouter(
	sampleController *controllers.SampleController,
	tenantController *controllers.TenantController,
	userController *controllers.UserController,
	eventController *controllers.EventController,
	ticketController *controllers.TicketController,
	transactionController *controllers.TransactionController,
	notificationPreferencesController *controllers.NotificationPreferencesController,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/samples", sampleController.GetAllSamples)
		api.POST("/samples", sampleController.CreateSample)
		api.GET("/tenants", tenantController.GetAllTenants)
		api.GET("/tenants/:id", tenantController.GetTenantByID)
		api.POST("/tenants", tenantController.CreateTenant)
		api.PUT("/tenants/:id", tenantController.UpdateTenant)
		api.DELETE("/tenants/:id", tenantController.DeleteTenant)
		api.GET("/users", userController.GetAllUsers)
		api.GET("/users/:id", userController.GetUserByID)
		api.POST("/users", userController.CreateUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
		api.GET("/events", eventController.GetAllEvents)
		api.GET("/events/:id", eventController.GetEventByID)
		api.POST("/events", eventController.CreateEvent)
		api.PUT("/events/:id", eventController.UpdateEvent)
		api.DELETE("/events/:id", eventController.DeleteEvent)
		api.GET("/tickets", ticketController.GetAllTickets)
		api.GET("/tickets/:id", ticketController.GetTicketByID)
		api.POST("/tickets", ticketController.CreateTicket)
		api.PUT("/tickets/:id", ticketController.UpdateTicket)
		api.DELETE("/tickets/:id", ticketController.DeleteTicket)
		api.POST("/tickets/purchase", ticketController.PurchaseTicket)
		api.POST("/tickets/sell", ticketController.SellTicket)
		api.POST("/tickets/authenticate", ticketController.AuthenticateTicket)
		api.POST("/tickets/refund", ticketController.RefundTicket)
		api.GET("/transactions", transactionController.GetAllTransactions)
		api.GET("/transactions/:id", transactionController.GetTransactionByID)
		api.POST("/transactions", transactionController.CreateTransaction)
		api.PUT("/transactions/:id", transactionController.UpdateTransaction)
		api.DELETE("/transactions/:id", transactionController.DeleteTransaction)
		api.GET("/notification-preferences", notificationPreferencesController.GetAllPreferences)
		api.GET("/notification-preferences/:id", notificationPreferencesController.GetPreferenceByID)
		api.POST("/notification-preferences", notificationPreferencesController.CreatePreference)
		api.PUT("/notification-preferences/:id", notificationPreferencesController.UpdatePreference)
		api.DELETE("/notification-preferences/:id", notificationPreferencesController.DeletePreference)
	}

	return router
}
