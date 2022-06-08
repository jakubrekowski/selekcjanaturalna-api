package main

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/jakubrekowski/selekcjanaturalna-api/database"
	"github.com/jakubrekowski/selekcjanaturalna-api/middleware"
	"github.com/jakubrekowski/selekcjanaturalna-api/utilities"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
)

func main() {
	credentials := os.Getenv("GOOGLE_CLOUD_CREDENTIALS")
	if credentials == "" {
		log.Fatal("GOOGLE_CLOUD_CREDENTIALS must be set")
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal("GOOGLE_CLOUD_PROJECT must be set")
	}

	ctx := context.Background()
	utilities.ProvideContext(ctx)

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentials))
	if err != nil {
		log.Fatalf("database.NewClient: %v", err)
	}

	receiptsDB, err := database.NewFirestoreDB(client, "Receipts")
	if err != nil {
		log.Fatalf("NewFirestoreDB (receipts): %v", err)
	}

	usersDB, err := database.NewFirestoreDB(client, "Users")
	if err != nil {
		log.Fatalf("NewFirestoreDB (users): %v", err)
	}

	utilities.PermsInit()

	fmt.Println(receiptsDB, usersDB)

	//token, err := utilities.SignToken(jwt.MapClaims{
	//	"sub":         "1234567890",
	//	"name":        "John Doe",
	//	"iat":         1516239022,
	//	"permissions": 7,
	//})
	//if err != nil {
	//	return
	//}
	//fmt.Println(token)

	middleware.Router()

	//for _, v := range Permissions {
	//	fmt.Println(v)
	//}
	//
	//pass := utilities.MissingPermissions(97, 73)
	//sth := utilities.GetPermissions(Permissions, pass)

	//fmt.Printf("%v: %v", pass, sth)
}
