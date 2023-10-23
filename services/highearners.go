package services

import (
	"fmt"
	"github.com/wasimkhandot01/assessment/config"
	"github.com/wasimkhandot01/assessment/models"
	"github.com/wasimkhandot01/assessment/responses"
	"github.com/wasimkhandot01/assessment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TopEarners() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := config.NewDBFactory()
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.HighEarnerResponse{
				Status:  http.StatusBadRequest,
				Message: "error1",
				Data:    []models.HighEarner{},
			})
			return
		}

		err2 := db.PrepareDB()
		if err2 != nil {
			c.JSON(http.StatusBadRequest, responses.HighEarnerResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    []models.HighEarner{},
			})
			return
		}

		fmt.Println("DB is Prepared!")

		data, err := db.TopThreeHighEarners(utils.Query)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.HighEarnerResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    []models.HighEarner{},
			})
			return
		}

		fmt.Println("Query Executed")
		highEarners := []models.HighEarner{}
		for _, row := range data {
			highEarner := models.HighEarner{
				Department: row.Department,
				Employee:   row.Employee,
				Salary:     row.Salary,
			}
			highEarners = append(highEarners, highEarner)
		}

		c.JSON(http.StatusOK, responses.HighEarnerResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    highEarners,
		})
	}
}
