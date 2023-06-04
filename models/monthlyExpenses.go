package models

type MonthlyExpenses struct {
	//Nombre del asto fijo mensual.
	NameFixedExpense string `json:"nameFixedExpense"`
	//Fecha de pago del gasto fijo mensual.
	DueDate string `json:"dueDate"`
}
