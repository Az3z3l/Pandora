package graph

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Reg structure thats used
type Reg struct {
	Email    string `bson:"Email" json:"Email,omitempty"`
	Username string `bson:"Username" json:"Username,omitempty"`
	Password string `bson:"Password" json:"Password,omitempty"`
	Age      string `bson:"Age" json:"Age,omitempty"`
	Gender   string `bson:"Gender" json:"Gender,omitempty"`
	// Institution string `bson:"Institution" json:"Institution,omitempty"`
	// Contact     string `bson:"Contact" json:"Contact,omitempty"`
	// Place       string `bson:"Place" json:"Place,omitempty"`
	// District    string `bson:"District" json:"District,omitempty"`
	// State       string `bson:"State" json:"State,omitempty"`
}

// LoginHandler login
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		client, err := MongoConnect()
		if err != nil {
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "Couldnt connect to database" }`))
			return
		}
		type logindata struct {
			User string
			Pass string
		}
		decoder := json.NewDecoder(request.Body)
		var t logindata
		err = decoder.Decode(&t)
		if err != nil {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "why you tampering  ;-;" }`))
			return
		}

		if t.User == "" || t.Pass == "" {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "empty data sent :/" }`))
			return
		}
		username := t.User
		password := t.Pass

		collection := client.Database("ctf").Collection("players")
		cursor, err := collection.Find(context.TODO(), bson.M{"Email": username})
		if err != nil {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "wrong username/password" }`))
			return
		}
		defer cursor.Close(context.TODO())
		var result *model.User
		for cursor.Next(context.TODO()) {
			var v *model.User
			err := cursor.Decode(&v)
			if err != nil {
				MongoDisconnect(client)
				response.Header().Set("content-type", "application/json")
				response.Write([]byte(`{ "Error": "wrong username/password" }`))
				return
			}
			if CheckPasswordHash(password, v.Password) {
				result = v
			}
		}
		// var result *model.User
		filter := bson.M{"Email": username}
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "wrong username/password" }`))
			return
		}
		if !CheckPasswordHash(password, result.Password) {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "wrong username/password" }`))
			return
		}
		if result.Username != "" && result.Email != "" {
			tokenString := CreateTokenEndpoint(result)
			// TODO return the API key

			cookie := http.Cookie{
				Name:  "auth",
				Value: tokenString,
				// Expires: expire,
				SameSite: 2,
				// SameSite: 4,
				Path: "/",
				// Secure:   true,	// only https
				// Domain: ".inctf.in",
			}

			http.SetCookie(response, &cookie)

			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error":  "ok"}`))
			return
		} else {

			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "Invalid" }`))
			return
		}

	} else {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Error": "Really?? ರ╭╮ರ" }`))
	}
}

// RegisterHandler register
func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		client, err := MongoConnect()
		if err != nil {
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "Couldnt connect to database" }`))
			return
		}
		type registerdata struct {
			Email    string
			Fullname string
			Pword1   string
			Pword2   string
			Uname    string
			Age      string
			Gender   string
		}
		decoder := json.NewDecoder(request.Body)
		var t registerdata
		err = decoder.Decode(&t)
		if err != nil {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "why you tampering  ;-;" }`))
			return
		}
		ok := true
		if t.Uname == "" {
			t.Uname = RandStringBytes()
		}
		/*if t.Age == "" || t.Gender == "" {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "empty fields spotted  o.O " }`))
			return
		}*/

		filternames := []string{"bi0s", "teambi0s", "team bi0s", "bios", "team bios", "admin", "administrator"}

		for _, s := range filternames {
			res1 := strings.EqualFold(t.Uname, s)
			res2 := strings.EqualFold(t.Fullname, s)
			if res1 || res2 {
				MongoDisconnect(client)
				response.Header().Set("content-type", "application/json")
				response.Write([]byte(`{ "Error": "I don't think you're from bi0s ಠ益ಠ" }`))
				return
			}
		}

		/*	listogender := map[string]bool{
				"Male":   true,
				"Female": true,
				"Other":  true,
			}

			if !listogender[t.Gender] {
				MongoDisconnect(client)
				response.Header().Set("content-type", "application/json")
				response.Write([]byte(`{ "Error": "Please select a given option for Gender" }`))
				return
			}*/

		if len(t.Pword1) < 8 || len(t.Pword2) < 8 {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "Backend Devs: Securing and testing all corner cases.Users: Using passwords less than length 8 (╯°□°）╯︵ ┻━┻" }`))
			return
		}

		// validate data
		if !isEmail(t.Email) {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "invalid email  ¯\_ಠಠ_/¯" }`))
			return
		}
		/*
			if _, err := strconv.ParseInt(t.Age, 10, 64); err != nil {
				MongoDisconnect(client)
				response.Header().Set("content-type", "application/json")
				response.Write([]byte(`{ "Error": "Are you really that old??? ಠ_ಠ" }`))
				return
			}

			age, _ := strconv.ParseInt(t.Age, 10, 64)
			if age < 1 || age > 100 {
				MongoDisconnect(client)
				response.Header().Set("content-type", "application/json")
				response.Write([]byte(`{ "Error": "Are you really that old??? ಠ_ಠ" }`))
				return
			}
		*/

		if t.Pword1 != t.Pword2 {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "Passwords do not match (╯︵╰,)" }`))
			return
		}
		pass, _ := HashPassword(t.Pword1)
		if ok {
			collection := client.Database("ctf").Collection("players")
			var result Reg
			err := collection.FindOne(context.TODO(), bson.M{"Email": t.Email}).Decode(&result)
			err1 := collection.FindOne(context.TODO(), bson.M{"Username": t.Uname}).Decode(&result)
			if err != nil {
				if err1 != nil {
					ins := map[string]string{
						"Email":    t.Email,
						"Fullname": t.Fullname,
						"Username": t.Uname,
						"Password": pass,
						"Age":      "",
						"Gender":   "",
					}
					_, err := collection.InsertOne(context.TODO(), ins)
					if err != nil {
						MongoDisconnect(client)
						response.Header().Set("content-type", "application/json")
						response.Write([]byte(`{ "Error": "error...rror...ror...or...r" }`))
						return
					}

					MongoDisconnect(client)
					response.Header().Set("content-type", "application/json")
					response.Write([]byte(`{ "Error": "Successful" }`))
				} else {
					MongoDisconnect(client)
					response.Header().Set("content-type", "application/json")
					response.Write([]byte(`{ "Error": "An account with that username already exists (─.─||）" }`))
				}
			} else {
				MongoDisconnect(client)
				response.Header().Set("content-type", "application/json")
				response.Write([]byte(`{ "Error": "An account with that email already exists (─.─||）" }`))
			}
		} else {
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Error": "data not defined  O_o" }`))
		}
	} else {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Error": "Really?? ರ╭╮ರ" }`))
	}
}

// CookieChecker - setting isAdmin in frontend
func CookieChecker(response http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("auth")
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "no" }`))
		return
	}

	var cookievalue = cookie.Value
	token := cookievalue
	username, email, err := ParseToken(token)
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "no" }`))
		return
	}
	if username == "" || email == "" {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "no" }`))
		return
	}
	if email == "admin@ctf.com" {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "ok" }`))
		return
	} else {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "logged in" }`))
		return
	}
}
