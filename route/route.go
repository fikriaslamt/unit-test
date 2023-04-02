package route

import (
	"sesi_8/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, handler handler.HttpServer) {
	api := r.Group("/books") // prefix
	{
		api.GET("", handler.GetAllBook)
		api.POST("", handler.AddBook)          // /employees
		api.GET("/:id", handler.GetBookById)   // /employee/:id
		api.PUT("/:id", handler.UpdateBook)    // /employee/:id
		api.DELETE("/:id", handler.DeleteBook) // /employee/:id
	}
}
