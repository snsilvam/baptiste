package handlers

import (
	"fmt"
	"net/http"

	"baptiste.com/models"
	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
)

type MessageError struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

func GetMonthlyFixedExpensesByIDHandler(c *gin.Context) {
	idDoc := c.Param("id")
	monthlyExpense, err := repository.GetMonthlyExpense(c, idDoc)

	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses/:id",
		})
		return
	}

	c.JSON(http.StatusOK, monthlyExpense)
}

func PostMonthlyFixedExpensesHandler(c *gin.Context) {
	var object models.MonthlyFixedExpensesModelInsert
	if err := c.BindJSON(&object); err != nil {
		fmt.Println("error al recibir el objeto en el request", err)
		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses",
		})
		return
	}

	fmt.Println("este fue el objeto enviado en el body", object)
	err := repository.InsertMonthlyExpenses(c, &object)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Monthly Expense creado con exito.")
}

func PatchMonthlyExpenseHandler(c *gin.Context) {
	var object models.MonthlyFixedExpensesModelUpdate
	if err := c.BindJSON(&object); err != nil {
		fmt.Println("error al recibir el objeto en el request", err)
		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses",
		})
		return
	}

	fmt.Println("este fue el objeto enviado en el request", object)
	if err := repository.UpdateMonthlyExpense(c, &object); err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Monthly Expense actualizado con exito.")
}
