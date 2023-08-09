package database

import (
	"context"
	"fmt"
	"time"

	"baptiste.com/models"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (f *FirestoreRepository) InsertTrackingMonthlyFixedExpense(ctx context.Context, object *models.TrackingMonthlyFixedExpensesInsert) error {
	trackingMonthlyFixedExpenseCollection := f.client.Collection("trackingMonthlyFixedExpense")

	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	object.Status = true

	wr, err := trackingMonthlyFixedExpenseCollection.NewDoc().Create(ctx, object)
	if err != nil {
		fmt.Println("error al intentar crear el documento(trackingMonthlyFixedExpense) en firestore:", err)
		return err
	}

	fmt.Println("El documento(trackingMonthlyFixedExpense) se creo con exito ☻", wr)

	return nil
}

func (f *FirestoreRepository) GetTrackingMonthlyFixedExpense(ctx context.Context, id string) (*models.TrackingMonthlyFixedExpenses, error) {
	var trackingMonthlyFixedExpense models.TrackingMonthlyFixedExpenses

	doc := f.client.Doc("trackingMonthlyFixedExpense/" + id)

	docsnap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err = docsnap.DataTo(&trackingMonthlyFixedExpense); err != nil {
		return nil, err
	}

	trackingMonthlyFixedExpense.ID = docsnap.Ref.ID

	return &trackingMonthlyFixedExpense, nil
}

func (f *FirestoreRepository) GetAllTrackingMonthlyFixedExpenses(ctx context.Context) (*[]models.TrackingMonthlyFixedExpenses, error) {
	collection := f.client.Collection("trackingMonthlyFixedExpense")

	var trackingMonthlyFixedExpense models.TrackingMonthlyFixedExpenses
	var trackingMonthlyFixedExpenses []models.TrackingMonthlyFixedExpenses

	iter := collection.Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error al obtener el siguiente documento: %v", err)
			return nil, err
		}

		if err = doc.DataTo(&trackingMonthlyFixedExpense); err != nil {
			return nil, err
		}
		trackingMonthlyFixedExpense.ID = doc.Ref.ID

		trackingMonthlyFixedExpenses = append(trackingMonthlyFixedExpenses, trackingMonthlyFixedExpense)
		fmt.Println(doc.Data())
	}

	return &trackingMonthlyFixedExpenses, nil
}

func (f *FirestoreRepository) UpdateTrackingMonthlyFixedExpense(ctx context.Context, object *models.TrackingMonthlyFixedExpenses) error {
	doc := f.client.Doc("trackingMonthlyFixedExpense/" + object.ID)

	object.UpdatedAt = time.Now()

	_, err := doc.Update(ctx, []firestore.Update{
		{Path: "Monthly", Value: object.Monthly},
		{Path: "MonthlyCost", Value: object.MonthlyCost},
		{Path: "MonthlyCostForEachUser", Value: object.MonthlyCostForEachUser},
		{Path: "PaymentStatus", Value: object.PaymentStatus},
		{Path: "UpdatedAt", Value: object.UpdatedAt},
		{Path: "Status", Value: object.Status},
	})

	if err != nil {
		return err
	}

	return nil
}
