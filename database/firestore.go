package database

import (
	"context"
	"fmt"

	"baptiste.com/models"
	"cloud.google.com/go/firestore"
)

type FirestoreRepository struct {
	client *firestore.Client
}

func NewClientFirestore(ctx context.Context, projectID string) (*FirestoreRepository, error) {
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	return &FirestoreRepository{client}, nil
}

func (f *FirestoreRepository) InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyExpensesModel) error {
	ny := f.client.Doc("monthlyExpensesModel")

	_, err := ny.Create(ctx, monthlyExpenses)
	fmt.Println("error--->", err)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreRepository) GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyExpensesModel, error) {
	var monthlyExpensesModel models.MonthlyExpensesModel

	doc := f.client.Doc("monthlyExpensesModel/" + id)

	docsnap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err = docsnap.DataTo(&monthlyExpensesModel); err != nil {
		return nil, err
	}

	monthlyExpensesModel.ID = docsnap.Ref.ID

	return &monthlyExpensesModel, nil
}
