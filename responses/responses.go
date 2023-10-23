package responses

import "github.com/wasimkhandot01/assessment/models"

type MinCostResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type HighEarnerResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    []models.HighEarner `json:"data"`
}

type Transaction struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Execution string `json:"execution"`
}

type CustomDataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
