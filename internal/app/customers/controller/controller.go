package controller

import (
	"log"
	"net/http"

	"github.com/MarianoArias/ApiGo/internal/app/customers/model"
	"github.com/gin-gonic/gin"
)

type cgetResponse struct {
	Customers []model.Customer `json:"customers"`
	Total     int64            `json:"total"`
}

// @Summary Get customers
// @Description Get customers from Elasticsearch
// @Accept json
// @Produce json
// @Success 200 {object} cgetResponse
// @Failure 500
// @Router /customers/ [get]
func CgetHandler(c *gin.Context) {
	customers, total, err := model.GetCustomers()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	response := cgetResponse{
		Customers: customers,
		Total:     total,
	}

	c.JSON(http.StatusOK, response)
}
