package services

import (
	"github.com/wasimkhandot01/assessment/config"
	"github.com/wasimkhandot01/assessment/models"
	"github.com/wasimkhandot01/assessment/responses"
	"github.com/wasimkhandot01/assessment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTablesTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {

		var tool models.Transaction
		if err := c.ShouldBindJSON(&tool); err != nil {
			c.JSON(http.StatusBadRequest, responses.Transaction{
				Status:    http.StatusBadRequest,
				Message:   "error",
				Execution: "Clint Side issue in request body.",
			})
			return
		}

		if validationErr := validateData.Struct(tool); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Transaction{
				Status:    http.StatusBadRequest,
				Message:   "error",
				Execution: "Client Side issue in request body.",
			})
			return
		}

		db, err := config.NewDBFactory()
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.HighEarnerResponse{
				Status:  http.StatusBadRequest,
				Message: "error1",
				Data:    []models.HighEarner{},
			})
			return
		}

		if tool.Tool == "slice" {

			err := db.TablesCreationTransaction(utils.CreateTableQueriesSlice)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.Transaction{
					Status:    http.StatusBadRequest,
					Message:   "error",
					Execution: "Query Failed to create table",
				})
				return
			}

		} else if tool.Tool == "map" {
			err := db.TablesCreationTransaction(utils.CreateTableQueriesMap)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.Transaction{
					Status:    http.StatusBadRequest,
					Message:   "error",
					Execution: "Query Failed to create table",
				})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, responses.Transaction{
				Status:    http.StatusBadRequest,
				Message:   "error",
				Execution: "Undefined tool in request body",
			})
			return
		}

		c.JSON(http.StatusOK, responses.Transaction{
			Status:    http.StatusOK,
			Message:   "success",
			Execution: "Tables are successfully created.",
		})
	}
}
