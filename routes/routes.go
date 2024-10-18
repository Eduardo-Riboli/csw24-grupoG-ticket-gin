package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grupoG/csw24-grupoG-ticket-gin/controllers"
)

func SetupRouter(
	sampleController *controllers.SampleController,
	tenantController *controllers.TenantController,
	userController *controllers.UserController,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/samples", sampleController.GetAllSamples)
		api.POST("/samples", sampleController.CreateSample)
		api.GET("/tenants", tenantController.GetAllTenants)
		api.GET("/tenants/:id", tenantController.GetTenantByID)
		api.POST("/tenants", tenantController.CreateTenant)
		api.GET("/users", userController.GetAllUsers)
		api.GET("/users/:id", userController.GetUserByID)
		api.POST("/users", userController.CreateUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}

	return router
}
