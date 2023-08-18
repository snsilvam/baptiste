package database

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreRepository struct {
	client *firestore.Client
}

func NewClientFirestore(ctx context.Context, projectID string) (*FirestoreRepository, error) {
	credentialsPath := "./baptiste-389101-3f5834d15587.json"
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, err
	}

	return &FirestoreRepository{client}, nil
}
