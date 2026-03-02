package routes

import (
	"Drawer/controllers"
	"Drawer/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	// 使用日志中间件和恢复中间件
	r.Use(middlewares.GinLogger(), gin.Recovery())

	api := r.Group("/api")
	{
		apiBill := api.Group("/bills")
		{
			apiTransactions := apiBill.Group("/transactions")
			{
				apiTransactions.POST("", controllers.Controller.CreateTransaction)
				apiTransactions.GET("", controllers.Controller.QueryTransactions)
				apiTransactions.PUT("/:id", controllers.Controller.UpdateTransaction)
				apiTransactions.DELETE("/:id", controllers.Controller.DeleteTransaction)
			}

			apiCategories := apiBill.Group("/categories")
			{
				apiCategories.POST("", controllers.Controller.CreateCategory)
				apiCategories.GET("", controllers.Controller.GetCategories)
				apiCategories.GET("/:id", controllers.Controller.GetCategory)
				apiCategories.PUT("/:id", controllers.Controller.UpdateCategory)
				apiCategories.DELETE("/:id", controllers.Controller.DeleteCategory)
			}

			apiPayees := apiBill.Group("/payees")
			{
				apiPayees.POST("", controllers.Controller.CreatePayee)
				apiPayees.GET("", controllers.Controller.GetPayees)
				apiPayees.GET("/:id", controllers.Controller.GetPayee)
				apiPayees.PUT("/:id", controllers.Controller.UpdatePayee)
				apiPayees.DELETE("/:id", controllers.Controller.DeletePayee)
			}

			apiAccounts := apiBill.Group("/accounts")
			{
				apiAccounts.POST("", controllers.Controller.CreateAccount)
				apiAccounts.GET("", controllers.Controller.GetAccounts)
				apiAccounts.GET("/:id", controllers.Controller.GetAccount)
				apiAccounts.PUT("/:id", controllers.Controller.UpdateAccount)
				apiAccounts.DELETE("/:id", controllers.Controller.DeleteAccount)
			}
		}
	}

	return r
}
