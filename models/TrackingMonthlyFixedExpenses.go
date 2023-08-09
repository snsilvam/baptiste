package models

import "time"

type TrackingMonthlyFixedExpenses struct {
	//Identificador del objeto en la base de datos.
	ID string `json:"ID"`
	//Identificador del usuario que realiza el seguimiento.
	IDUser string `json:"IDUser"`
	//Identificador del gasto fijo mensual, al que se realiza el seguimiento.
	IDFixedExpense string `json:"IDFixedExpense"`
	//Fecha de creacion.
	CreatedAt time.Time `json:"createdAt"`
	//Fecha de actualizacion.
	UpdatedAt time.Time `json:"updatedAt"`
	//Mes en el que se realiza el seguimiento.
	Monthly string `json:"monthly"`
	//Dinero a pagar en el mes actual para el gasto fijo mensual.
	MonthlyCost int64 `json:"monthlyCost"`
	//Dinero a pagar en el mes actual, para cada usuario.
	MonthlyCostForEachUser int64 `json:"monthlyCostForEachUser"`
	//Estado del pago.
	PaymentStatus string `json:"paymentStatus"`
	//Estado del seguimiento en la plataforma, si es falso, quiere decir que esta eliminado o desactivado.
	Status bool `json:"status"`
}

type TrackingMonthlyFixedExpensesInsert struct {
	//Identificador del usuario que realiza el seguimiento.
	IDUser string `json:"IDUser"`
	//Identificador del gasto fijo mensual, al que se realiza el seguimiento.
	IDFixedExpense string `json:"IDFixedExpense"`
	//Fecha de creacion.
	CreatedAt time.Time `json:"createdAt"`
	//Fecha de actualizacion.
	UpdatedAt time.Time `json:"updatedAt"`
	//Mes en el que se realiza el seguimiento.
	Monthly string `json:"monthly"`
	//Dinero a pagar en el mes actual para el gasto fijo mensual.
	MonthlyCost int64 `json:"monthlyCost"`
	//Dinero a pagar en el mes actual, para cada usuario.
	MonthlyCostForEachUser int64 `json:"monthlyCostForEachUser"`
	//Estado del pago.
	PaymentStatus string `json:"paymentStatus"`
	//Estado del seguimiento en la plataforma, si es falso, quiere decir que esta eliminado o desactivado.
	Status bool `json:"status"`
}
