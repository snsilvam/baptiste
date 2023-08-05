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

func (f *FirestoreRepository) InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyFixedExpensesModelInsert) error {
	collectionMonthlyFixedExpenses := f.client.Collection("monthlyFixedExpensesModel")

	wr, err := collectionMonthlyFixedExpenses.NewDoc().Create(ctx, monthlyExpenses)
	if err != nil {
		fmt.Println("error al intentar crear el documento en firestore:", err)
		return err
	}

	fmt.Println("El documento se creo con exito ☻", wr)

	return nil
}

func (f *FirestoreRepository) GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyFixedExpensesModel, error) {
	var monthlyFixedExpensesModel models.MonthlyFixedExpensesModel

	doc := f.client.Doc("monthlyFixedExpensesModel/" + id)

	docsnap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err = docsnap.DataTo(&monthlyFixedExpensesModel); err != nil {
		return nil, err
	}

	monthlyFixedExpensesModel.ID = docsnap.Ref.ID

	return &monthlyFixedExpensesModel, nil
}

func (f *FirestoreRepository) UpdateMonthlyExpense(ctx context.Context, monthlyExpense *models.MonthlyFixedExpensesModelUpdate) error {
	doc := f.client.Doc("monthlyFixedExpensesModel/" + monthlyExpense.ID)

	_, err := doc.Update(ctx, []firestore.Update{{Path: "NameFixedExpense", Value: monthlyExpense.NameFixedExpense}, {Path: "DueDate", Value: monthlyExpense.DueDate}, {Path: "Status", Value: monthlyExpense.Status}})
	if err != nil {
		return err
	}
	return nil
}

/* func (f *FirestoreRepository) GetAllMonthlyFixedExpenses(ctx context.Context, userId string) ([]*models.MonthlyExpensesModel, error) {

	return
} */
