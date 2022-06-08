package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

type firestoreDB struct {
	client     *firestore.Client
	collection string
}

var _ UserDatabase = &firestoreDB{}
var _ ReceiptDatabase = &firestoreDB{}

func NewFirestoreDB(client *firestore.Client, collection string) (*firestoreDB, error) {
	ctx := context.Background()

	err := client.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("firebasedb: could not connect: %v", err)
	}

	return &firestoreDB{
		client:     client,
		collection: collection,
	}, nil
}

func (db *firestoreDB) Close(context.Context) error {
	return db.client.Close()
}
