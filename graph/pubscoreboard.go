package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserScoreboard struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}

func Pubscore(response http.ResponseWriter, request *http.Request) {

	client, err := MongoConnect()
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(""))
		return
	}

	var v model.Manager
	collectionsettings := client.Database("ctf").Collection("settings")
	err = collectionsettings.FindOne(context.Background(), bson.D{}).Decode(&v)
	stat := *v.ScoreboardStats
	if stat == 0 { // check if user is logged in
		cookie, err := request.Cookie("auth")
		if err != nil {
			http.Error(response, "data: Not Logged In", http.StatusBadRequest)
			return
		}

		var cookievalue = cookie.Value
		token := cookievalue
		username, email, err := ParseToken(token)
		if err != nil {
			http.Error(response, "data:User Not Logged In", http.StatusBadRequest)
			return
		}
		if username == "" || email == "" {
			http.Error(response, "data:User Not Logged In", http.StatusBadRequest)
			return
		}

	}

	collection := client.Database("ctf").Collection("players")
	sortscore := bson.D{{"$sort", bson.D{{"Score", -1}, {"lastsolve", 1}}}}
	// sorttime := bson.D{{"$sort", bson.D{{"lastsolve", 1}}}}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{sortscore})
	if err != nil {
		Logger(err)
		MongoDisconnect(client)
		val := ""
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(val))
		return
	}
	defer cursor.Close(context.TODO())
	result := []UserScoreboard{}
	for cursor.Next(context.TODO()) {
		var v UserScoreboard
		err := cursor.Decode(&v)
		if err != nil {
			MongoDisconnect(client)
			val := ""
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(val))
			return
		}
		result = append(result, v)

	}
	MongoDisconnect(client)
	b, err := json.Marshal(result)
	fmt.Println(string(b))
	val := "{\"scoreboard\" :" + string(b) + "}"
	response.Header().Set("content-type", "application/json")
	response.Write([]byte(val))
	return
}
