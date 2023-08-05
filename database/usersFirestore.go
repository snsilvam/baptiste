package database

import (
	"context"
	"fmt"
	"time"

	"baptiste.com/models"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (f *FirestoreRepository) InsertUser(ctx context.Context, user *models.UserInsert) error {
	usersCollection := f.client.Collection("users")

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	wr, err := usersCollection.NewDoc().Create(ctx, user)
	if err != nil {
		fmt.Println("error al intentar crear el documento(User) en firestore:", err)
		return err
	}

	fmt.Println("El documento se creo con exito ☻", wr)

	return nil
}

func (f *FirestoreRepository) GetUser(ctx context.Context, id string) (*models.Users, error) {
	var user models.Users

	doc := f.client.Doc("users/" + id)

	docsnap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err = docsnap.DataTo(&user); err != nil {
		return nil, err
	}

	user.ID = docsnap.Ref.ID

	return &user, nil
}

func (f *FirestoreRepository) GetAllUsers(ctx context.Context) (*[]models.Users, error) {
	collection := f.client.Collection("users")

	var user models.Users
	var users []models.Users

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

		if err = doc.DataTo(&user); err != nil {
			return nil, err
		}
		user.ID = doc.Ref.ID

		users = append(users, user)
		fmt.Println(doc.Data())
	}

	return &users, nil
}

func (f *FirestoreRepository) UpdateUser(ctx context.Context, user *models.Users) error {
	doc := f.client.Doc("users/" + user.ID)

	user.UpdatedAt = time.Now()

	_, err := doc.Update(ctx, []firestore.Update{
		{Path: "Name", Value: user.Name},
		{Path: "Email", Value: user.Email},
		{Path: "Status", Value: user.Status},
		{Path: "UpdatedAt", Value: user.UpdatedAt},
	})

	if err != nil {
		return err
	}

	return nil
}
