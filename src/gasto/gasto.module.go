package gasto

import (
	database "baptiste.com/database"
	"github.com/gin-gonic/gin"
)

type GastoModule struct {
	Controller *GastoController
}

func ConstructorGastoModule(database database.Database) *GastoModule {
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
		gastoRoutes.GET("", gm.Controller.GetAllGastos)
		gastoRoutes.PUT("", gm.Controller.UpdateGasto)
	}
}
