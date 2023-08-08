package models

import "time"

type MonthlyFixedExpenses struct {
	//Identificador del objeto(MonthlyFixedExpenses) en la base de datos.
	ID string `json:"ID"`
	//Fecha de creacion.
	CreatedAt time.Time `json:"createdAt"`
	//Fecha de actualizacion.
	UpdatedAt time.Time `json:"updatedAt"`
	//Nombre del gasto fijo mensual.
	Name string `json:"name"`
	//Fecha en la que se debe realizar el pago.
	DueDate string `json:"dueDate"`
	//Estado del gasto fijo mensual en la plataforma, si es falso, quiere decir que eliminado o desactivado.
	Status bool `json:"status"`
}

type MonthlyFixedExpensesInsert struct {
	//Fecha de creacion.
	CreatedAt time.Time `json:"createdAt"`
	//Fecha de actualizacion.
	UpdatedAt time.Time `json:"updatedAt"`
	//Nombre del gasto fijo mensual.
	Name string `json:"name"`
	//Fecha en la que se debe realizar el pago.
	DueDate string `json:"dueDate"`
	//Estado del gasto fijo mensual en la plataforma, si es falso, quiere decir que eliminado o desactivado.
	Status bool `json:"status"`
}
