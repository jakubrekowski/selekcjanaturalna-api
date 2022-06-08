package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/jakubrekowski/selekcjanaturalna-api/utilities"
	"google.golang.org/api/iterator"
)

type Contact struct {
	Discord  string
	Email    string
	Phone    string
	Facebook string
}

type User struct {
	Number               int8
	Name                 string
	Contact              *Contact
	Paid                 int16
	Needed               int16 // (Money)
	Type                 int8
	Permissions          int16
	TemporaryPermissions []*utilities.Permission
	AddedBy              string
	AddedDate            int64
}

type UserDoc struct {
	ID   string
	User *User
}

type UserDatabase interface {
	ListUsers(context.Context) ([]*UserDoc, error)
	GetUser(ctx context.Context, id string) (user *UserDoc, err error)
	AddUser(ctx context.Context, u *UserDoc) (id string, err error)
	DeleteUser(ctx context.Context, u *UserDoc) error
}

func (db *firestoreDB) ListUsers(ctx context.Context) (user []*UserDoc, err error) {
	users := make([]*UserDoc, 0)
	iter := db.client.Collection(db.collection).Query.OrderBy("Number", firestore.Asc).Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("firebaseDB: User: could not list users: %v", err)

		}
		u := &User{}
		doc.DataTo(u)
		ud := &UserDoc{
			ID:   doc.Ref.ID,
			User: u,
		}
		users = append(users, ud)
	}
	return users, nil
}

func (db *firestoreDB) GetUser(ctx context.Context, id string) (user *UserDoc, err error) {
	ds, err := db.client.Collection(db.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("firebaseDB: User: Get: %v", err)
	}
	u := &User{}
	ds.DataTo(u)
	ud := &UserDoc{
		ID:   id,
		User: u,
	}
	return ud, nil
}

func (db *firestoreDB) AddUser(ctx context.Context, u *UserDoc) (id string, err error) {
	ref := db.client.Collection(db.collection).NewDoc()

	if _, err := ref.Create(ctx, u.User); err != nil {
		return "", fmt.Errorf("user: create: %v", err)
	}
	return ref.ID, nil
}

func (db *firestoreDB) DeleteUser(ctx context.Context, u *UserDoc) error {
	if _, err := db.client.Collection(db.collection).Doc(u.ID).Delete(ctx); err != nil {
		return fmt.Errorf("user: delete: %v", err)
	}
	return nil
}
