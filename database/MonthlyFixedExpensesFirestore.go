package database

import (
	"context"
	"fmt"
	"time"

	"baptiste.com/models"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (f *FirestoreRepository) InsertMonthlyFixedExpense(ctx context.Context, monthlyExpense *models.MonthlyFixedExpensesInsert) error {
	collectionMonthlyFixedExpenses := f.client.Collection("monthlyFixedExpenses")

	monthlyExpense.CreatedAt = time.Now()
	monthlyExpense.UpdatedAt = monthlyExpense.CreatedAt
	monthlyExpense.Status = true

	wr, err := collectionMonthlyFixedExpenses.NewDoc().Create(ctx, monthlyExpense)
	if err != nil {
		fmt.Println("error al intentar crear el documento(monthlyExpense) en firestore:", err)
		return err
	}

	fmt.Println("El documento(monthlyExpense) se creo con exito ☻", wr)

	return nil
}

func (f *FirestoreRepository) GetMonthlyFixedExpense(ctx context.Context, id string) (*models.MonthlyFixedExpenses, error) {
	var monthlyFixedExpenses models.MonthlyFixedExpenses

	doc := f.client.Doc("monthlyFixedExpenses/" + id)

	docsnap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err = docsnap.DataTo(&monthlyFixedExpenses); err != nil {
		return nil, err
	}

	monthlyFixedExpenses.ID = docsnap.Ref.ID

	return &monthlyFixedExpenses, nil
}

func (f *FirestoreRepository) GetAllMonthlyFixedExpenses(ctx context.Context) (*[]models.MonthlyFixedExpenses, error) {
	collection := f.client.Collection("monthlyFixedExpenses")

	var monthlyExpense models.MonthlyFixedExpenses
	var monthlyExpenses []models.MonthlyFixedExpenses

	iter := collection.Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error al obtener el siguiente documento: %v", err)
			return nil, err
		}

		if err = doc.DataTo(&monthlyExpense); err != nil {
			return nil, err
		}
		monthlyExpense.ID = doc.Ref.ID

		monthlyExpenses = append(monthlyExpenses, monthlyExpense)
		fmt.Println(doc.Data())
	}

	return &monthlyExpenses, nil
}

func (f *FirestoreRepository) UpdateMonthlyFixedExpense(ctx context.Context, monthlyExpense *models.MonthlyFixedExpenses) error {
	doc := f.client.Doc("monthlyFixedExpenses/" + monthlyExpense.ID)

	monthlyExpense.UpdatedAt = time.Now()

	_, err := doc.Update(ctx, []firestore.Update{
		{Path: "Name", Value: monthlyExpense.Name},
		{Path: "DueDate", Value: monthlyExpense.DueDate},
		{Path: "Status", Value: monthlyExpense.Status},
		{Path: "UpdatedAt", Value: monthlyExpense.UpdatedAt},
	})

	if err != nil {
		return err
	}

	return nil
}
