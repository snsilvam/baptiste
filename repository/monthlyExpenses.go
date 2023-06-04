package repository

import (
	"context"

	"baptiste.com/models"
)

type MonthlyExpenses interface {
	InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyExpensesModel) error
	GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyExpensesModel, error)
}

var implementation MonthlyExpenses

func SetRepository(repository MonthlyExpenses) {
	implementation = repository
}

func InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyExpensesModel) error {
	return implementation.InsertMonthlyExpenses(ctx, monthlyExpenses)
}

func GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyExpensesModel, error) {
	return implementation.GetMonthlyExpense(ctx, id)
}
