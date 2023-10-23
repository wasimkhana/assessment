package routes

import (
	"github.com/wasimkhandot01/assessment/services"
	"github.com/wasimkhandot01/assessment/utils"

	"github.com/gin-gonic/gin"
)

var rateLimiter = utils.NewIPRateLimiter()

func TasksRoute(router *gin.Engine) {
	router.POST("/task/oneandsix/mincoststairs", rateLimiter.LimitRequest(), services.FindMinCostSteps())
	router.POST("/task/four/transaction", services.CreateTablesTransaction())
	router.POST("/task/three/custombinding", services.CustomBinding())
	router.GET("/task/five/highearners", services.TopEarners())
}
