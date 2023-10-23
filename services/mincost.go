package services

import (
	"github.com/wasimkhandot01/assessment/models"
	"github.com/wasimkhandot01/assessment/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validateData = validator.New()

func FindMinCostSteps() gin.HandlerFunc {
	return func(c *gin.Context) {

		var mincost models.MinCost
		if err := c.ShouldBindJSON(&mincost); err != nil {
			c.JSON(http.StatusBadRequest, responses.MinCostResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		if validationErr := validateData.Struct(mincost); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.MinCostResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()},
			})
			return
		}

		minCostStairs(&mincost)
		c.JSON(http.StatusOK, responses.MinCostResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"mincost": mincost.Mincost_found},
		})
	}
}

func minCostStairs(mincost *models.MinCost) {
	var cost []int
	cost = mincost.Cost

	totalSteps := len(cost)
	if totalSteps <= 1 {
		mincost.Mincost_found = 0
	}
	singleStep := cost[0]
	doubleStep := cost[1]

	if totalSteps <= 2 {
		mincost.Mincost_found = min(singleStep, doubleStep)
	}

	for i := 2; i < totalSteps; i++ {
		currentStep := cost[i] + min(singleStep, doubleStep)
		singleStep, doubleStep = doubleStep, currentStep
	}

	mincost.Mincost_found = minCost(singleStep, doubleStep)
}

func minCost(a, b int) int {
	if a < b {
		return a
	}
	return b
}
