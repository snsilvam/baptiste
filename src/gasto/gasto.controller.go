package gasto

import (
	"fmt"
	"log"
	"net/http"
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
	hoy := time.Now()
	fmt.Println("Fecha de hoy:", hoy)
	if err := c.BindJSON(&gastos); err != nil {
		log.Println("error al recibir el gasto: ", err)

		c.JSON(404, handlers.MessageError{
			Message: "error al recibir el gasto: " + err.Error(),
			Url:     "Gastos",
		})
	}

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
