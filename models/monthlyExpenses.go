package models

type MonthlyExpensesModel struct {
	ID string `json:"id"`
	//Nombre del gasto fijo mensual.
	NameFixedExpense string `json:"nameFixedExpense"`
	//Fecha de pago del gasto fijo mensual.
	DueDate string `json:"dueDate"`
}
