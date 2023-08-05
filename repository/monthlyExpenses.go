package repository

import (
	"context"

	"baptiste.com/models"
)

type MonthlyExpensesRepository interface {
	InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyFixedExpensesModelInsert) error
	GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyFixedExpensesModel, error)
	UpdateMonthlyExpense(ctx context.Context, monthlyExpense *models.MonthlyFixedExpensesModelUpdate) error
}

var implementation MonthlyExpensesRepository

func SetRepository(repository MonthlyExpensesRepository) {
	implementation = repository
}

func InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyFixedExpensesModelInsert) error {
	return implementation.InsertMonthlyExpenses(ctx, monthlyExpenses)
}

func GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyFixedExpensesModel, error) {
	return implementation.GetMonthlyExpense(ctx, id)
}

func UpdateMonthlyExpense(ctx context.Context, monthlyExpense *models.MonthlyFixedExpensesModelUpdate) error {
	return implementation.UpdateMonthlyExpense(ctx, monthlyExpense)
}
