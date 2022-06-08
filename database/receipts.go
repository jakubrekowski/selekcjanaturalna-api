package database

import (
	"context"
	"fmt"
)

type Receipt struct {
	Title     string
	Value     float32
	Filename  string
	Date      string
	User      string
	AddedBy   string
	AddedDate int64
}

type ReceiptDatabase interface {
	AddReceipt(ctx context.Context, r *Receipt) (id string, err error)
}

func (db *firestoreDB) AddReceipt(ctx context.Context, r *Receipt) (id string, err error) {
	ref := db.client.Collection(db.collection).NewDoc()

	if _, err := ref.Create(ctx, r); err != nil {
		return "", fmt.Errorf("receipt: create: %v", err)
	}
	return ref.ID, nil
}
