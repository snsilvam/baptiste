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
	var trackingMonthlyFixedExpense models.TrackingMonthlyFixedExpenses
	var counter int64
	var idsMonthlyFixedExpenses []string

	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	object.Status = true
	object.Month = object.CreatedAt.Month()
	object.MonthlyCostForEachUser = object.MonthlyCost

	trackingMonthlyFixedExpenseCollection := f.client.Collection("trackingMonthlyFixedExpense")

	/*1. Se realiza una consulta con el objetivo de filtrar los seguimientos del mes actual.*/
	response := trackingMonthlyFixedExpenseCollection.Where("Status", "==", true).Where("Monthly", "==", object.Month)

	/*2. Se comprueba si en los seguimientos encontrados existen relacion con el gasto fijo mensual al
	que se le va realizar el seguimiento por medio de su IDFixedExpense*/
	iter := response.Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if err = doc.DataTo(&trackingMonthlyFixedExpense); err != nil {
			return err
		}

		if trackingMonthlyFixedExpense.IDFixedExpense == object.IDFixedExpense {
			idsMonthlyFixedExpenses = append(idsMonthlyFixedExpenses, doc.Ref.ID)
			counter++
		}

		fmt.Println(doc.Data())
	}

	/*3. Si existen registros relacionados con el gasto fijo mensual, se divide el costo de este mes en
	la cantidad de registros encontrados, mas el seguimiento actual.*/
	if counter > 0 {
		fmt.Println("IDS:------------------>", idsMonthlyFixedExpenses)
		counter++
		object.MonthlyCostForEachUser = object.MonthlyCostForEachUser / counter

		wr, err := trackingMonthlyFixedExpenseCollection.NewDoc().Create(ctx, object)
		if err != nil {
			fmt.Println("error al intentar crear el documento(trackingMonthlyFixedExpense) en firestore:", err)
			return err
		}

		fmt.Println("El documento(trackingMonthlyFixedExpense) se creo con exito ☻", wr)
		/*4.  Se actualiza cada  seguimiento del gasto fijo mensual, con el nuevo costo de este mes
		para cada usuario.*/
		for _, id := range idsMonthlyFixedExpenses {
			doc := f.client.Doc("trackingMonthlyFixedExpense/" + id)

			_, err := doc.Update(ctx, []firestore.Update{
				{Path: "MonthlyCostForEachUser", Value: object.MonthlyCostForEachUser},
				{Path: "UpdatedAt", Value: object.UpdatedAt},
			})
			if err != nil {
				return err
			}
		}

	} else {
		wr, err := trackingMonthlyFixedExpenseCollection.NewDoc().Create(ctx, object)
		if err != nil {
			fmt.Println("error al intentar crear el documento(trackingMonthlyFixedExpense) en firestore:", err)
			return err
		}

		fmt.Println("El documento(trackingMonthlyFixedExpense) se creo con exito ☻", wr)
	}

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
		{Path: "Monthly", Value: object.Month},
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
