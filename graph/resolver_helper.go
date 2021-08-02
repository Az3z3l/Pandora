package graph

import (
	"context"
	"crypto/rand"
	"log"
	"os"

	// "log"
	"regexp"

	//b64 "encoding/base64"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func isEmail(str string) bool {
	var email string
	email = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	rx := regexp.MustCompile(email)
	return rx.MatchString(str)
}

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ.abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// MongoConnect returns mongo client
func MongoConnect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Logger(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Logger(err)
		return nil, err
	}
	return client, err
}

// GetUserIdfromname uses String
func GetUserIdfromname(u string) (string, error) {
	var newlogin *model.User
	client, err := MongoConnect()
	collection := client.Database("ctf").Collection("players")
	err = collection.FindOne(context.TODO(), bson.M{"Username": u}).Decode(&newlogin)
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

// adminCheck - isAdmin?
func adminCheck(ctx context.Context) bool {
	client, err := MongoConnect()
	collection1 := client.Database("ctf").Collection("players")
	id := ForContext(ctx)
	docid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", docid}}
	if err != nil {
		MongoDisconnect(client)
		return false
	}
	var v1 *model.User
	err1 := collection1.FindOne(context.TODO(), filter).Decode(&v1)

	if err1 != nil {
		MongoDisconnect(client)
		return false
	}
	if v1.Email == "admin@ctf.com" {
		MongoDisconnect(client)
		return true
	}
	MongoDisconnect(client)
	return false
}

// MongoDisconnect stops mongo client - need an alternative
func MongoDisconnect(client *mongo.Client) {
	_ = client.Disconnect(context.TODO())
}

// Have a seperate file for logging errors
func Logger(err error) {
	file, er := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		log.Panic(er)
	}

	log.SetOutput(file)
	log.Println("=============================================================================================================")
	log.Println(err)
	log.Println("=============================================================================================================")
}
