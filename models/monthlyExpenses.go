package models

type MonthlyExpensesModel struct {
	ID string `json:"id"`
	//Nombre del gasto fijo mensual.
	NameFixedExpense string `json:"nameFixedExpense"`
	//Fecha de pago del gasto fijo mensual.
	DueDate string `json:"dueDate"`
}

type MonthlyExpensesModelInsert struct {
	//Nombre del gasto fijo mensual.
	NameFixedExpense string `json:"nameFixedExpense"`
	//Fecha de pago del gasto fijo mensual.
	DueDate string `json:"dueDate"`
	//Status define si el objeto esta activo o inactivo.
	Status string `json:"status"`
}

type MonthlyExpensesModelUpdate struct {
	ID string `json:"id"`
	//Nombre del gasto fijo mensual.
	NameFixedExpense *string `json:"nameFixedExpense"`
	//Fecha de pago del gasto fijo mensual.
	DueDate *string `json:"dueDate"`
	//Status define si el objeto esta activo o inactivo.
	Status string `json:"status"`
}
