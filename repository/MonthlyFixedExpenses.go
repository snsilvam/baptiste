package repository

import (
	"context"

	"baptiste.com/models"
)

type MonthlyFixedExpensesRepository interface {
	InsertMonthlyFixedExpense(ctx context.Context, monthlyFixedExpense *models.MonthlyFixedExpensesInsert) error
	GetMonthlyFixedExpense(ctx context.Context, id string) (*models.MonthlyFixedExpenses, error)
	GetAllMonthlyFixedExpenses(ctx context.Context) (*[]models.MonthlyFixedExpenses, error)
	UpdateMonthlyFixedExpense(ctx context.Context, monthlyFixedExpense *models.MonthlyFixedExpenses) error
}

var implementation MonthlyFixedExpensesRepository

func SetRepository(repository MonthlyFixedExpensesRepository) {
	implementation = repository
}

func InsertMonthlyFixedExpense(ctx context.Context, monthlyFixedExpense *models.MonthlyFixedExpensesInsert) error {
	return implementation.InsertMonthlyFixedExpense(ctx, monthlyFixedExpense)
}

func GetMonthlyFixedExpense(ctx context.Context, id string) (*models.MonthlyFixedExpenses, error) {
	return implementation.GetMonthlyFixedExpense(ctx, id)
}

func GetAllMonthlyFixedExpenses(ctx context.Context) (*[]models.MonthlyFixedExpenses, error) {
	return implementation.GetAllMonthlyFixedExpenses(ctx)
}

func UpdateMonthlyFixedExpense(ctx context.Context, monthlyFixedExpense *models.MonthlyFixedExpenses) error {
	return implementation.UpdateMonthlyFixedExpense(ctx, monthlyFixedExpense)
}
