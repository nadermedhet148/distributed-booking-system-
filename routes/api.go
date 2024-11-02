package routes

import (
	"os"

	"github.com/coroo/go-starter/app/console"
	"github.com/coroo/go-starter/app/deliveries"
	"github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/app/usecases"
	"github.com/gin-gonic/gin"
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
		AvailableSlotRepository repositories.AvailableSlotRepository = repositories.NewAvailableSlotRepository()
		AvailableSlotService    usecases.AvailableSlotService        = usecases.NewAvailableSlotService(AvailableSlotRepository)
	)
	deliveries.NewAvailableSlotController(router, API_PREFIX, AvailableSlotService)

	console.Schedule()
	router.Run(":" + os.Getenv("MAIN_PORT"))
}
