package handlers

import (
	"fmt"
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
	var object models.MonthlyExpensesModelInsert
	if err := c.BindJSON(&object); err != nil {
		fmt.Println("error al recibir el objeto en el request", err)
		return
	}

	fmt.Println("este fue el objeto enviado en el body", object)
	err := repository.InsertMonthlyExpenses(c, &object)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.IndentedJSON(http.StatusOK, "Monthly Expense creado con exito.")
}
