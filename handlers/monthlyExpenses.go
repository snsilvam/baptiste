package handlers

import (
	"net/http"

	"baptiste.com/models"
	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type MessageError struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

func GetMonthlyExpensesByIDHandler(c *gin.Context) {
	idDoc := c.Param("id")
	monthlyExpense, err := repository.GetMonthlyExpense(c, idDoc)

	if err != nil {
		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses/:id",
		})
		return
	}

	c.JSON(http.StatusOK, monthlyExpense)
}

func PostMonthlyExpensesHandler(c *gin.Context) {
	object := &models.MonthlyExpensesModel{
		NameFixedExpense: "internet",
		DueDate:          "15 de cada mes",
	}

	err := repository.InsertMonthlyExpenses(c, object)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	}

	c.IndentedJSON(http.StatusOK, "goood")
}
