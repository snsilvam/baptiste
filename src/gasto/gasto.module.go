package gasto

import (
	postgres "baptiste.com/database"
	"github.com/gin-gonic/gin"
)

type GastoModule struct {
	Controller *GastoController
}

func ConstructorGastoModule(database postgres.DatabasePostgres) *GastoModule {
	service := ConstructorGastoService(database)

	gasto := &GastoModule{
		Controller: ConstructorGastoController(*service),
	}
	return gasto
}

func (gm *GastoModule) RegisterGastoRoutes(router *gin.Engine) {
	gastoRoutes := router.Group("/gastos")
	{
		gastoRoutes.POST("", gm.Controller.CreateGasto)
		gastoRoutes.GET("", gm.Controller.Hello)
	}
}
