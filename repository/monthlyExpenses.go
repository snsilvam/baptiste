package repository

import (
	"context"

	"baptiste.com/models"
)

type MonthlyExpensesRepository interface {
	InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyExpensesModel) error
	GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyExpensesModel, error)
}

var implementation MonthlyExpensesRepository

func SetRepository(repository MonthlyExpensesRepository) {
	implementation = repository
}

func InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyExpensesModel) error {
	return implementation.InsertMonthlyExpenses(ctx, monthlyExpenses)
}

func GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyExpensesModel, error) {
	return implementation.GetMonthlyExpense(ctx, id)
}
