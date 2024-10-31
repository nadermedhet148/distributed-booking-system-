package routes

import (
	"os"

	"github.com/coroo/go-starter/app/console"
	"github.com/coroo/go-starter/app/deliveries"
	"github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/app/usecases"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Api() {
	router := gin.Default()
	// router.Use(middlewares.BasicAuth())

	API_PREFIX := os.Getenv("API_PREFIX")

	router.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": os.Getenv("MAIN_DESCRIPTION"),
		})
	})
	var (
		gopayLinkingRepository repositories.GopayLinkingRepository = repositories.NewGopayLinkingRepository()
		gopayLinkingService    usecases.GopayLinkingService        = usecases.NewGopayLinkingService(gopayLinkingRepository)
	)
	deliveries.NewGopayLinkingController(router, API_PREFIX, gopayLinkingService)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	console.Schedule()
	router.Run(":" + os.Getenv("MAIN_PORT"))
}
