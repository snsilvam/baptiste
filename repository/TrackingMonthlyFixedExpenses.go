package repository

import (
	"context"

	"baptiste.com/models"
)

type TrackingMonthlyFixedExpensesRepository interface {
	InsertTrackingMonthlyFixedExpense(ctx context.Context, object *models.TrackingMonthlyFixedExpensesInsert) error
	GetTrackingMonthlyFixedExpense(ctx context.Context, id string) (*models.TrackingMonthlyFixedExpenses, error)
	GetAllTrackingMonthlyFixedExpenses(ctx context.Context) (*[]models.TrackingMonthlyFixedExpenses, error)
	UpdateTrackingMonthlyFixedExpense(ctx context.Context, object *models.TrackingMonthlyFixedExpenses) error
}

var trackingMonthlyFixedExpenseImplementation TrackingMonthlyFixedExpensesRepository

func SetTrackingMonthlyFixedExpensesRepository(repository TrackingMonthlyFixedExpensesRepository) {
	trackingMonthlyFixedExpenseImplementation = repository
}

func InsertTrackingMonthlyFixedExpense(ctx context.Context, object *models.TrackingMonthlyFixedExpensesInsert) error {
	return trackingMonthlyFixedExpenseImplementation.InsertTrackingMonthlyFixedExpense(ctx, object)
}

func GetTrackingMonthlyFixedExpense(ctx context.Context, id string) (*models.TrackingMonthlyFixedExpenses, error) {
	return trackingMonthlyFixedExpenseImplementation.GetTrackingMonthlyFixedExpense(ctx, id)
}

func GetAllTrackingMonthlyFixedExpenses(ctx context.Context) (*[]models.TrackingMonthlyFixedExpenses, error) {
	return trackingMonthlyFixedExpenseImplementation.GetAllTrackingMonthlyFixedExpenses(ctx)
}

func UpdateTrackingMonthlyFixedExpense(ctx context.Context, object *models.TrackingMonthlyFixedExpenses) error {
	return trackingMonthlyFixedExpenseImplementation.UpdateTrackingMonthlyFixedExpense(ctx, object)
}
