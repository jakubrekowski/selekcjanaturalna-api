package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/jakubrekowski/selekcjanaturalna-api/utilities"
	"io"
	"log"
	"net/http"
	"reflect"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bearerToken := r.Header.Get("Authorization")
	claims, err := utilities.VerifyTokenFromHeader(bearerToken)
	if err != nil {
		status := &Status{
			Status:  401,
			Message: "Authorization token not provided.",
		}

		w.WriteHeader(http.StatusUnauthorized)
		statusJson, _ := json.Marshal(status)
		io.WriteString(w, string(statusJson))
		return
	}

	var permissions = reflect.ValueOf(claims["permissions"])

	//permissions.Float()
	//missingPermissions := utilities.MissingPermissions(2, 0)
	missingPermissions := utilities.MissingPermissions(2, int64(permissions.Float()))

	if missingPermissions != 0 {
		status := &Status{
			Status:  401,
			Message: "You don't have all the required permissions: { READ_USERS }",
		}

		w.WriteHeader(http.StatusUnauthorized)
		statusJson, _ := json.Marshal(status)
		io.WriteString(w, string(statusJson))
		return
	}

	//ctx := r.Context()

	//users, err := database.UserDatabase.ListUsers(ctx)
	//if err != nil {
	//	return
	//}

	//usersJson, _ := json.Marshal(users)

	//io.WriteString(w, string(usersJson))
	io.WriteString(w, "{}")

	log.Println("GET /users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bearerToken := r.Header.Get("Authorization")

	claims, err := utilities.VerifyTokenFromHeader(bearerToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(claims)
}
