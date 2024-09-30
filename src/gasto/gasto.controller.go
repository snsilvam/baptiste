package gasto

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"baptiste.com/handlers"
	gastos "baptiste.com/src/gasto/models"
	"github.com/gin-gonic/gin"
)

type GastoController struct {
	service GastoService
}

func ConstructorGastoController(service GastoService) *GastoController {
	return &GastoController{
		service: service,
	}
}

func (gc *GastoController) CreateGasto(c *gin.Context) {
	var gastos gastos.Gastos
	if err := c.BindJSON(&gastos); err != nil {
		log.Println("error al recibir el gasto: ", err)

		c.JSON(404, handlers.MessageError{
			Message: "error al recibir el gasto: " + err.Error(),
			Url:     "Gastos",
		})
		return
	}

	hoy := time.Now()
	fmt.Println("Fecha de hoy:", hoy)
	gastos.FechaDelGasto = hoy

	err := gc.service.CreateGasto(c, &gastos)
	if err != nil {
		c.JSON(500, handlers.MessageError{
			Message: err.Error(),
			Url:     "error creando el usuario en la base de datos.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Gastos almacenados en la base de datos.")
}

func (gc *GastoController) GetGastoByID(c *gin.Context) {
	idGastoStr := c.Param("id")

	idGasto, err := strconv.Atoi(idGastoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.MessageError{
			Message: "ID inválido, debe ser un número.",
			Url:     "/id",
		})
		return
	}

	gasto, err := gc.service.GetGastoByID(c, idGasto)
	if err != nil {
		c.JSON(500, handlers.MessageError{
			Message: "error consultando el gasto: " + err.Error(),
			Url:     "/id",
		})
		return
	}

	c.JSON(http.StatusOK, gasto)
}

func (gc *GastoController) GetAllGastos(c *gin.Context) {
	gastos, err := gc.service.GetAllGastos(c)
	if err != nil {
		c.JSON(500, handlers.MessageError{
			Message: err.Error(),
			Url:     "error consultando la tabla gastos.",
		})
		return
	}

	c.JSON(http.StatusOK, gastos)
}

func (gc *GastoController) UpdateGasto(c *gin.Context) {
	var gastos gastos.Gastos
	if err := c.BindJSON(&gastos); err != nil {
		log.Println("error al recibir el gasto: ", err)

		c.JSON(404, handlers.MessageError{
			Message: "error al recibir el gasto: " + err.Error(),
			Url:     "Gastos",
		})
		return
	}

	err := gc.service.UpdateGasto(c, &gastos)
	if err != nil {
		c.JSON(500, handlers.MessageError{
			Message: err.Error(),
			Url:     "error creando el usuario en la base de datos.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Gastos actualizados en la base de datos.")
}
