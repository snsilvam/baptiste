package handlers

import (
	"fmt"
	"net/http"

	"baptiste.com/models"
	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
)

func PostMonthlyFixedExpense(c *gin.Context) {
	var monthlyFixedExpensesInsert models.MonthlyFixedExpensesInsert

	if err := c.BindJSON(&monthlyFixedExpensesInsert); err != nil {
		fmt.Println("error al recibir el objeto(MonthlyFixedExpensesInsert) en el request", err)

		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/monthly-fixed-expenses",
		})
		return
	}

	fmt.Println("este fue el objeto(monthlyFixedExpensesInsert) enviado en el body", monthlyFixedExpensesInsert)

	err := repository.InsertMonthlyFixedExpense(c, &monthlyFixedExpensesInsert)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-expenses",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Monthly Fixed Expense creado con exito.")
}

func GetMonthlyFixedExpense(c *gin.Context) {
	idDoc := c.Param("id")
	monthlyFixedExpense, err := repository.GetMonthlyFixedExpense(c, idDoc)

	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-fixed-expenses/" + idDoc,
		})
		return
	}

	c.JSON(http.StatusOK, monthlyFixedExpense)
}

func GetAllMonthlyFixedExpenses(c *gin.Context) {
	monthlyFixedExpenses, err := repository.GetAllMonthlyFixedExpenses(c)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-fixed-expenses/",
		})
		return
	}

	c.JSON(http.StatusOK, monthlyFixedExpenses)
}

func UpdateMonthlyFixedExpense(c *gin.Context) {
	var monthlyFixedExpense models.MonthlyFixedExpenses
	if err := c.BindJSON(&monthlyFixedExpense); err != nil {
		fmt.Println("error al recibir el objeto(MonthlyFixedExpense) en el request", err)

		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/monthly-fixed-expenses",
		})
		return
	}

	fmt.Println("este fue el objeto(MonthlyFixedExpense) enviado en el request", monthlyFixedExpense)
	if err := repository.UpdateMonthlyFixedExpense(c, &monthlyFixedExpense); err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/monthly-fixed-expenses",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Monthly Expense actualizado con exito.")
}
