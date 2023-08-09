package handlers

import (
	"fmt"
	"net/http"

	"baptiste.com/models"
	"baptiste.com/repository"
	"github.com/gin-gonic/gin"
)

func PostTrackingMonthlyFixedExpense(c *gin.Context) {
	var trackingMonthlyFixedExpense models.TrackingMonthlyFixedExpensesInsert

	if err := c.BindJSON(&trackingMonthlyFixedExpense); err != nil {
		fmt.Println("error al recibir el objeto(trackingMonthlyFixedExpense) en el request", err)

		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/tracking-monthly-fixed-expenses",
		})
		return
	}

	fmt.Println("este fue el objeto(trackingMonthlyFixedExpense) enviado en el body", trackingMonthlyFixedExpense)

	err := repository.InsertTrackingMonthlyFixedExpense(c, &trackingMonthlyFixedExpense)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/tracking-monthly-fixed-expenses",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "TrackingMonthlyFixedExpense creado con exito.")
}

func GetTrackingMonthlyFixedExpense(c *gin.Context) {
	idDoc := c.Param("id")
	trackingMonthlyFixedExpense, err := repository.GetTrackingMonthlyFixedExpense(c, idDoc)

	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/tracking-monthly-fixed-expenses/" + idDoc,
		})
		return
	}

	c.JSON(http.StatusOK, trackingMonthlyFixedExpense)
}

func GetTrackingMonthlyFixedExpenses(c *gin.Context) {
	trackingMonthlyFixedExpenses, err := repository.GetAllTrackingMonthlyFixedExpenses(c)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/tracking-monthly-fixed-expenses/",
		})
		return
	}

	c.JSON(http.StatusOK, trackingMonthlyFixedExpenses)
}

func UpdateTrackingMonthlyFixedExpense(c *gin.Context) {
	var trackingMonthlyFixedExpenses models.TrackingMonthlyFixedExpenses

	if err := c.BindJSON(&trackingMonthlyFixedExpenses); err != nil {
		fmt.Println("error al recibir el objeto(UserUpdated) en el request", err)

		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/tracking-monthly-fixed-expenses",
		})
		return
	}

	fmt.Println("este fue el objeto(UserUpdated) enviado en el request", trackingMonthlyFixedExpenses)

	if err := repository.UpdateTrackingMonthlyFixedExpense(c, &trackingMonthlyFixedExpenses); err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/tracking-monthly-fixed-expenses",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "trackingMonthlyFixedExpenses actualizado con exito.")
}
