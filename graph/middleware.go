package graph

import (
	"context"
	"net/http"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	//	"go.mongodb.org/mongo-driver/bson"
	//"context"
	//"github.com/Az3z3l/GQL-SERVER/graph/model"
	//jwt "github.com/dgrijalva/jwt-go"
)

type tokener struct {
	operationName string
	query         string
	variables     string
}

var userCtxKey = &contextKey{"username"}

type contextKey struct {
	username string
}

// User A stand-in for our database backed user object
type User struct {
	username string
	email    string
}

// Admin looks
type Admin struct {
	username string
	isAdmin  bool
}

// GetUserIDByUsername uses user object
func GetUserIDByUsername(u User) (string, error) {
	var newlogin *model.User
	client, err := MongoConnect()
	collection := client.Database("ctf").Collection("players")
	err = collection.FindOne(context.TODO(), bson.M{"Email": u.email}).Decode(&newlogin)
	if err != nil {
		MongoDisconnect(client)
		return "", err
	}
	if newlogin.Username != "" && newlogin.Email != "" {
		MongoDisconnect(client)
		return string(newlogin.ID), nil
	}
	MongoDisconnect(client)
	return "", nil
}

// Middleware Checks cookie of user sending the request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			var key string

			cookie, err := r.Cookie("auth")
			if err != nil {
				http.Error(w, "data: Not Logged In", http.StatusBadRequest)
				return
			}

			var cookievalue = cookie.Value
			token := cookievalue
			username, email, err := ParseToken(token)
			if err != nil {
				http.Error(w, "data:User Not Logged In", http.StatusBadRequest)
				return
			}
			if username == "" || email == "" {
				http.Error(w, "data:User Not Logged In", http.StatusBadRequest)
				return
			}

			s := User{username: username, email: email}

			key, err = GetUserIDByUsername(s)

			// fmt.Println(key)
			// fmt.Println(ForContext(ctx))
			ctx := context.WithValue(r.Context(), userCtxKey, key)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// AdminMiddleware gql playground only to admins
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("auth")
		if err != nil {
			http.Error(w, "Do you have any idea how stupid I am? Don’t underestimate me.” – Kondou Isao", http.StatusBadRequest)
			return
		}

		var cookievalue = cookie.Value
		token := cookievalue
		username, email, err := ParseToken(token)
		if err != nil {
			http.Error(w, "Do you have any idea how stupid I am? Don’t underestimate me.” – Kondou Isao", http.StatusBadRequest)
			return
		}
		if username == "" || email == "" {
			http.Error(w, "Do you have any idea how stupid I am? Don’t underestimate me.” – Kondou Isao", http.StatusBadRequest)
			return
		}

		if email == "admin@ctf.com" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Do you have any idea how stupid I am? Don’t underestimate me.” – Kondou Isao", http.StatusBadRequest)
			return
		}
	})
}

// ForContext Get the data from context
func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
