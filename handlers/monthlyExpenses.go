package handlers

import (
	"net/http"

	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func GetMonthlyExpensesByIDHandler(c *gin.Context) {
	monthlyExpense, err := repository.GetMonthlyExpense(c, "0LcjDmB1O0ldHjd2obBT")

	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.IndentedJSON(http.StatusOK, monthlyExpense)
}
