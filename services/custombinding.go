package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wasimkhandot01/assessment/models"
	"github.com/wasimkhandot01/assessment/responses"
	"io/ioutil"
	"net/http"
)

func CustomBinding() gin.HandlerFunc {
	return func(c *gin.Context) {
		var customdata models.CustomData

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CustomDataResponse{
				Status:  http.StatusBadRequest,
				Message: "Error in custom binding",
				Error:   err.Error(),
			})
			return
		}

		if err := json.Unmarshal(body, &customdata); err != nil {
			c.JSON(http.StatusBadRequest, responses.CustomDataResponse{
				Status:  http.StatusBadRequest,
				Message: "Error in custom binding",
				Error:   err.Error(),
			})
			return
		}

		if validationErr := validateData.Struct(customdata); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CustomDataResponse{
				Status:  http.StatusBadRequest,
				Message: "Failed validating custom binding",
				Error:   validationErr.Error(),
			})
			return
		}

		response := fmt.Sprintf("Received data: BoolField=%v, IntField=%d, StringField=%s, PointerField=%s, Nested (Head=%d, Tail=%d)",
			customdata.BoolField, customdata.IntField, customdata.StringField, *customdata.PointerField, customdata.ExtraField.Head, customdata.ExtraField.Tail)

		c.JSON(http.StatusOK, responses.CustomDataResponse{
			Status:  http.StatusOK,
			Message: response,
			Error:   "nil",
		})
	}
}
