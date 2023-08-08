package database

import (
	"context"

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
