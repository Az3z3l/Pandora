package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Az3z3l/GQL-SERVER/graph/generated"
	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *challwhenResolver) User(ctx context.Context, obj *model.Challwhen) (*model.User, error) {
	ouser, err := UserID(ctx, obj.User)
	if err != nil {
		Logger(err)
		return nil, err
	}
	return ouser, nil
}

// Challwhen returns generated.ChallwhenResolver implementation.
func (r *Resolver) Challwhen() generated.ChallwhenResolver { return &challwhenResolver{r} }

type challwhenResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func UserID(ctx context.Context, id1 string) (*model.User, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	id := id1
	docid, err := primitive.ObjectIDFromHex(id)
	//fmt.Println(username)
	filter := bson.D{{"_id", docid}}
	if err != nil {
		MongoDisconnect(client)
		return nil, nil
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	var result *model.User
	for cursor.Next(context.TODO()) {
		var v *model.User
		err := cursor.Decode(&v)
		if err != nil {
			Logger(err)
			MongoDisconnect(client)
			return nil, nil
		}
		result = v
	}
	MongoDisconnect(client)
	return result, nil
}
