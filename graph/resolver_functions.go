package graph

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func addchallenge(ctx context.Context, challenge *model.AddChallengeData) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
		// return "Who da heck are you??", nil
	}
	client, err := MongoConnect()
	if err != nil {
		return "Database Error", nil
	}

	collection := client.Database("ctf").Collection("settings")
	exist := false
	listo, err := client.Database("ctf").ListCollectionNames(nil, bson.D{})
	for _, names := range listo {
		if names == "settings" {
			exist = true
		}
	}
	if !exist {
		var condition *model.Manager
		*condition.Status = 1
		*condition.Details = "CTF has not started"
		_, _ = Updaterules(ctx, condition)
	}
	collection = client.Database("ctf").Collection("challenges")
	result, err := collection.InsertOne(
		ctx,
		&challenge)
	if err != nil {
		MongoDisconnect(client)
		Logger(err)
		return "Failed", nil
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	filter := bson.D{{"_id", oid}}
	update := bson.M{"$set": bson.M{"visibility": false}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		return "Failed", nil
	}

	return "OK", nil

}

func editchallenge(ctx context.Context, input *model.EditChallengeData) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	client, err := MongoConnect()
	if err != nil {
		return "Mongo Connection Error ", nil
	}
	collection := client.Database("ctf").Collection("challenges")
	id := input.ID
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		MongoDisconnect(client)
		return "Cannot find object ID", nil
	}
	filter := bson.D{{"_id", docid}}
	update := bson.M{"$set": bson.M{"name": input.Name, "description": input.Description, "category": input.Category, "tags": input.Tags, "value": input.Value, "flags": input.Flags}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	MongoDisconnect(client)
	if erro != nil {
		Logger(erro)
		return "Error updating", nil
	}

	return "OK", nil
}

func removefilechall(ctx context.Context, input *model.Delfile) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	filey := "challenges/" + input.ID + "/" + input.Name
	if _, err := os.Stat(filey); os.IsNotExist(err) {
		return "file does not exist", nil
	}
	client, err := MongoConnect()
	if err != nil {
		return "Mongo Connection Error ", nil
	}
	collection := client.Database("ctf").Collection("challenges")
	id := input.ID
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		MongoDisconnect(client)
		return "Cannot find object ID", nil
	}
	filter := bson.D{{"_id", docid}}
	update := bson.M{"$pull": bson.M{"file": input.Name}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		return "error removing name from db", nil
	}
	err = os.Remove(filey)
	if err != nil {
		return "Removed file in db. Unable to delete it from server", nil
	}

	return "OK", nil
}

func changevisibility(ctx context.Context, input *model.Public) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}

	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	client, err := MongoConnect()
	if err != nil {
		return "Mongo Connection Error ", nil
	}
	collection := client.Database("ctf").Collection("challenges")
	id := input.ID
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		MongoDisconnect(client)
		return "Cannot find object ID", nil
	}
	filter := bson.D{{"_id", docid}}
	update := bson.M{"$set": bson.M{"visibility": input.Visibility}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	MongoDisconnect(client)
	if erro != nil {
		Logger(err)
		return "Error updating", nil
	}

	return "OK", nil
}

func DelChall(ctx context.Context, id string) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}

	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	client, err := MongoConnect()
	if err != nil {
		return "Mongo Connection Error ", nil
	}
	collection := client.Database("ctf").Collection("challenges")
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		MongoDisconnect(client)
		return "Cannot find object ID", nil
	}
	filter := bson.D{{"_id", docid}}
	_, erro := collection.DeleteOne(
		context.Background(),
		filter,
	)
	MongoDisconnect(client)
	if erro != nil {
		Logger(err)
		return "Error deleting challenge", nil
	}
	folder := "challenges/" + id + "/"
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		err := os.RemoveAll(folder)
		if err != nil {
			return "Challenge removed. Unable to delete associated files", nil
		}
	}

	return "OK", nil
}

func user(ctx context.Context) ([]*model.User, error) {

	// fmt.Println(ForContext(ctx))

	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		Logger(err)
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	var result []*model.User

	for cursor.Next(context.TODO()) {
		var v *model.User
		err := cursor.Decode(&v)
		if err != nil {
			MongoDisconnect(client)
			return nil, nil
		}
		// fmt.Println(v.ID)
		// fmt.Println(v.Solved)
		result = append(result, v)

	}
	MongoDisconnect(client)
	return result, nil

}

func auser(ctx context.Context, username string) (*model.User, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	filter := bson.D{{"Username", username}}
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

func UserbyID(ctx context.Context) (*model.User, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	id := ForContext(ctx)
	docid, err := primitive.ObjectIDFromHex(id)
	//fmt.Println(username)
	filter := bson.D{{"_id", docid}}
	if err != nil {
		MongoDisconnect(client)
		return nil, err
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

func challenge(ctx context.Context) ([]*model.Challenge, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("challenges")
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}

	var v model.Manager
	collectionsettings := client.Database("ctf").Collection("settings")
	err = collectionsettings.FindOne(ctx, bson.D{}).Decode(&v)
	stat := *v.Status
	if stat == 1 && !admin {
		MongoDisconnect(client)
		return nil, nil
	}

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	var result []*model.Challenge

	// remove flag  if not admin
	if !admin {
		for cursor.Next(context.TODO()) {
			var v *model.Challenge
			err := cursor.Decode(&v)
			if err != nil {
				MongoDisconnect(client)
				return nil, nil
			}
			if v.Visibility == true {
				v.Flags = "flag{No,I_am_your_father}"
				result = append(result, v)
			}
		}
	} else {
		for cursor.Next(context.TODO()) {
			var v *model.Challenge
			err := cursor.Decode(&v)
			if err != nil {
				Logger(err)
				MongoDisconnect(client)
				return nil, nil
			}
			result = append(result, v)
		}
	}
	MongoDisconnect(client)
	return result, nil
}

func achallenge(ctx context.Context, id string) (*model.Challenge, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	var v model.Manager
	collectionsettings := client.Database("ctf").Collection("settings")
	err = collectionsettings.FindOne(ctx, bson.D{}).Decode(&v)
	stat := *v.Status
	if stat == 1 && !admin {
		MongoDisconnect(client)
		return nil, nil
	}

	collection := client.Database("ctf").Collection("challenges")
	docid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", docid}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	var result *model.Challenge
	for cursor.Next(context.TODO()) {
		var v *model.Challenge
		err := cursor.Decode(&v)
		if err != nil {
			Logger(err)
			MongoDisconnect(client)
			return nil, nil
		}
		result = v

	}
	if !admin {
		result.Flags = "flag{No,I_am_your_father}"
	}
	MongoDisconnect(client)
	return result, nil
}

func useredit(ctx context.Context, input *model.Useredit) (string, error) {

	x := *input
	updata := map[string]interface{}{}
	filternames := []string{"bi0s", "teambi0s", "team bi0s", "bios", "team bios", "admin", "administrator"}

	if x.Username != nil {

		for _, s := range filternames {
			res2 := strings.EqualFold(*x.Username, s)
			if res2 {
				return "Invalid_Username", nil
			}
		}

		updata["Username"] = *x.Username
	}
	if x.Fullname != nil {
		for _, s := range filternames {
			res2 := strings.EqualFold(*x.Fullname, s)
			if res2 {
				return "Invalid_Fullname", nil
			}
		}

		updata["Fullname"] = *x.Fullname
	}
	if x.Age != nil {

		if _, err := strconv.ParseInt(*input.Age, 10, 64); err != nil {
			return "Age_Error_not_int", nil
		}
		age, _ := strconv.ParseInt(*input.Age, 10, 64)
		if age < 1 || age > 100 {
			return "Age_Error_invalid", nil

		}

		updata["Age"] = *x.Age
	}
	if x.Gender != nil {

		listogender := map[string]bool{
			"Male":   true,
			"Female": true,
			"Other":  true,
		}

		if !listogender[*input.Gender] {
			return "Invalid_Gender", nil
		}

		updata["Gender"] = *x.Gender
	}
	if x.Institution != nil {
		updata["Institution"] = *x.Institution
	}

	if x.Contact != nil {
		if _, err := strconv.ParseInt(*input.Contact, 10, 64); err != nil {
			return "Invalid_Contact_not_int", nil
		}
		age, _ := strconv.ParseInt(*input.Contact, 10, 64)
		if age < 1000000000 || age > 9999999999 {
			return "Invalid_Contact_length", nil
		}

		updata["Contact"] = *x.Contact
	}
	if x.Place != nil {
		updata["Place"] = *x.Place
	}
	if x.District != nil {
		updata["District"] = *x.District
	}
	if x.State != nil {
		updata["State"] = *x.State
	}
	client, err := MongoConnect()
	collection := client.Database("ctf").Collection("players")
	id := ForContext(ctx)
	// fmt.Println(id)
	docid, err := primitive.ObjectIDFromHex(id)
	// fmt.Println(docid)
	filter := bson.D{{"_id", docid}}
	if err != nil {
		MongoDisconnect(client)
		return "Invalid cookie", nil
	}
	var u *model.Fulluser
	// email exists
	err1 := collection.FindOne(context.TODO(), filter).Decode(&u)
	if err1 != nil {
		MongoDisconnect(client)
		return "Invalid Cookie", nil
	}

	update := bson.M{"$set": updata}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		Logger(erro)
		MongoDisconnect(client)
		return "Error updating", nil
	}

	// 	MongoDisconnect(client)
	// fmt.Println(*input)

	return "OK", nil
}

func ressetpwd(ctx context.Context, input *model.Resetpwd) (string, error) {
	var u *model.Passwdreset
	client, err := MongoConnect()
	collection := client.Database("ctf").Collection("players")
	id := ForContext(ctx)
	docid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		MongoDisconnect(client)
		return "Cannot find object ID", nil
	}
	filter := bson.D{{"_id", docid}}
	err1 := collection.FindOne(context.TODO(), filter).Decode(&u)
	if err1 != nil {
		MongoDisconnect(client)
		return "Invalid Cookie", nil
	}
	if !CheckPasswordHash(input.Oldpwd, u.Password) {
		MongoDisconnect(client)
		return "Old Password Mismatch", nil
	}
	if input.Newpwda != input.Newpwdb {
		MongoDisconnect(client)
		return "New Password Mismatch", nil
	}
	if len(input.Newpwda) < 8 {
		MongoDisconnect(client)
		return "Weak Password", nil
	}
	pass, _ := HashPassword(input.Newpwda)
	update := bson.M{"$set": bson.M{"password": pass}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		MongoDisconnect(client)
		Logger(erro)
		return "Error updating", nil
	}

	MongoDisconnect(client)
	return "OK", nil
}

func submitFlag(ctx context.Context, input string) (string, error) {
	client, err := MongoConnect()
	Chacollection := client.Database("ctf").Collection("challenges")
	Placollection := client.Database("ctf").Collection("players")
	Setcollection := client.Database("ctf").Collection("settings")

	var v model.Manager
	err = Setcollection.FindOne(ctx, bson.D{}).Decode(&v)
	stat := *v.Status
	if stat == 3 {
		MongoDisconnect(client)
		return "Flag submission stopped", nil
	}

	id := ForContext(ctx)
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		MongoDisconnect(client)
		return "Cannot find object ID", nil
	}
	var chall *model.Challenge
	var playa *model.User
	flag := input
	filter := bson.M{"flags": bson.M{"$eq": flag}}
	err = Chacollection.FindOne(ctx, filter).Decode(&chall)
	if err != nil {
		MongoDisconnect(client)
		return "Check your flag again", nil
	}
	if chall.Visibility == false {
		MongoDisconnect(client)
		return "Check your flag again", nil
	}

	if stat == 2 {
		MongoDisconnect(client)
		return "Correct flag!! ", nil
	}

	// chall HAS NOW THE CHALLENGE THAT HAS THE FLAG
	challid, err := primitive.ObjectIDFromHex(chall.ID)
	filter = bson.M{"_id": bson.M{"$eq": docid}}
	err = Placollection.FindOne(ctx, filter).Decode(&playa)
	if err != nil {
		MongoDisconnect(client)
		return "Something went wrong...", nil
	}
	tstamp := time.Now().UnixNano() / 1e6
	fsu := false
	// this is the first solve for the team ..... create an array inside the key --> solved
	if playa.Solved == nil {
		update := bson.M{"$set": bson.M{"solved": []interface{}{
			bson.M{
				"ChallID":   challid,
				"Timestamp": tstamp,
			}}}}
		_, erro := Placollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			return "Flag Submission Failed Something went wrong...", nil
		}
		fsu = true
	} else { // already solved atleast one challenge ..... see if this challenge is solved already
		for _, i := range playa.Solved {
			if i.Challenge == chall.ID {
				MongoDisconnect(client)
				return "Already solved this Challenge", nil
			}
		}
	}
	// this challenge not solved yet
	if !fsu {
		update := bson.M{"$push": bson.M{"solved": bson.M{"ChallID": challid, "Timestamp": tstamp}}}
		_, erro := Placollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			return "Flag Submission Failed Something went wrong...", nil
		}
	}
	// Added data to USERS side =>  Make the changes in challenges side
	filter = bson.M{"_id": bson.M{"$eq": challid}}
	// First blood needs spl attention ~ create an array before pushing
	if chall.Teamssolved == nil {
		update := bson.M{"$set": bson.M{"Teamssolved": []interface{}{
			bson.M{
				"UserID":    docid,
				"Timestamp": tstamp,
			}}}}
		_, erro := Chacollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			return "Flag Submission Failed Something went wrong...", nil
		}
	} else {
		update := bson.M{"$push": bson.M{"Teamssolved": bson.M{"UserID": docid, "Timestamp": tstamp}}}
		_, erro := Chacollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			MongoDisconnect(client)
			return "Flag Submission Failed Something went wrong...", nil
		}
	}
	// atomic

	filter = bson.M{"_id": bson.M{"$eq": docid}}
	score := chall.Value
	update := bson.M{"$inc": bson.M{"Score": score}}
	_, erro := Placollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		MongoDisconnect(client)
		return "Error updating", nil
	}
	filter = bson.M{"_id": bson.M{"$eq": challid}}
	solves := 1
	update = bson.M{"$inc": bson.M{"solves": solves}}
	_, erro = Chacollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		MongoDisconnect(client)
		return "Error updating", nil
	}

	// add last solved time to rank players with same score
	filter = bson.M{"_id": bson.M{"$eq": docid}}
	update = bson.M{"$set": bson.M{"lastsolve": tstamp}}
	_, erro = Placollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		MongoDisconnect(client)
		Logger(erro)
		return "Error updating", nil
	}

	MongoDisconnect(client)
	return "Congratulations on the solve", nil

}

func rankin(ctx context.Context) ([]*model.User, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	sortscore := bson.D{{"$sort", bson.D{{"Score", -1}, {"lastsolve", 1}}}}
	// sorttime := bson.D{{"$sort", bson.D{{"lastsolve", 1}}}}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{sortscore})
	if err != nil {
		Logger(err)
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	var result []*model.User
	for cursor.Next(context.TODO()) {
		var v *model.User
		err := cursor.Decode(&v)
		if err != nil {
			MongoDisconnect(client)
			return nil, nil
		}
		result = append(result, v)

	}
	MongoDisconnect(client)
	return result, nil
}

func fullyuser(ctx context.Context, ida string) (*model.Fulluser, error) {
	// if it is admin return data of the username passed
	// else return data from ctx
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	var result *model.Fulluser
	var admin bool
	admin = false
	done := false
	if adminCheck(ctx) {
		admin = true
	}
	var id string
	if admin && ida != "me" {
		id = ida
		filter := bson.D{{"Username", id}}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			MongoDisconnect(client)
			return nil, nil
		}
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var v *model.Fulluser
			err := cursor.Decode(&v)
			if err != nil {
				Logger(err)
				MongoDisconnect(client)
				return nil, nil
			}
			result = v
		}
		done = true
	}
	if !done {
		id = ForContext(ctx)
		docid, err := primitive.ObjectIDFromHex(id)
		//fmt.Println(username)
		filter := bson.D{{"_id", docid}}
		if err != nil {
			MongoDisconnect(client)
			return nil, err
		}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			MongoDisconnect(client)
			return nil, nil
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var v *model.Fulluser
			err := cursor.Decode(&v)
			if err != nil {
				Logger(err)
				MongoDisconnect(client)
				return nil, nil
			}
			result = v
		}
	}
	return result, nil
}

func meeseeks(ctx context.Context) (*model.Fulluser, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("players")
	var result *model.Fulluser

	id := ForContext(ctx)
	docid, err := primitive.ObjectIDFromHex(id)
	//fmt.Println(username)
	filter := bson.D{{"_id", docid}}
	if err != nil {
		MongoDisconnect(client)
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var v *model.Fulluser
		err := cursor.Decode(&v)
		if err != nil {
			Logger(err)
			MongoDisconnect(client)
			return nil, nil
		}
		result = v
	}
	return result, nil
}

func adminpassreset(ctx context.Context, input *model.AdminPreset) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	id, err := GetUserIdfromname(input.Name)
	if err != nil {
		return "error trying to get user", nil
	}
	if id == "" {
		return "no user found", nil
	}
	docid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", docid}}
	if err != nil {
		return "objectid error", nil
	}
	client, err := MongoConnect()
	if err != nil {
		MongoDisconnect(client)
		return "Error with mongodb", nil
	}
	collection := client.Database("ctf").Collection("players")
	pass, _ := HashPassword(input.Password)
	update := bson.M{"$set": bson.M{"password": pass}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		Logger(err)
		MongoDisconnect(client)
		return "Error updating", nil
	}

	MongoDisconnect(client)
	return "ok", nil
}

func add_notification(ctx context.Context, input model.Notificationinp) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	client, err := MongoConnect()
	if err != nil {
		return "error connecting to MongoDB", nil
	}
	// var x model.Notification
	// x.Name = input.Name
	// x.Description = input.Description
	tstamp := time.Now().UnixNano() / 1e6
	collection := client.Database("ctf").Collection("notifications")
	update := bson.M{"name": input.Name, "description": input.Description, "timestamp": tstamp}
	_, err = collection.InsertOne(
		ctx,
		&update,
	)

	if err != nil {
		Logger(err)
		MongoDisconnect(client)
		return "Failed", nil
	}
	MongoDisconnect(client)
	return "OK", nil
}

func editNoti(ctx context.Context, input model.Notifiedit) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	client, err := MongoConnect()
	if err != nil {
		MongoDisconnect(client)
		return "Cannot connect to Database", nil
	}
	docid, err := primitive.ObjectIDFromHex(input.ID)
	collection := client.Database("ctf").Collection("notifications")
	filter := bson.D{{"_id", docid}}
	update := bson.M{"$set": bson.M{"name": input.Name, "description": input.Description}}
	_, erro := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		Logger(erro)
		MongoDisconnect(client)
		return "Error updating", nil
	}
	return "OK", nil
}

func delNoti(ctx context.Context, id string) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "id error", nil
	}
	client, err := MongoConnect()
	if err != nil {
		// err1 := errors.New("connect error")
		MongoDisconnect(client)
		return "Error with mongoDB", nil
	}
	collection := client.Database("ctf").Collection("notifications")
	filter := bson.D{{"_id", docid}}
	_, erro := collection.DeleteOne(
		context.Background(),
		filter,
	)
	MongoDisconnect(client)
	if erro != nil {
		Logger(err)
		return "Error deleting challenge", nil
	}
	return "OK", nil
}

func OneNoti(ctx context.Context, id string) (*model.Notification, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, nil
	}
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil
	}
	collection := client.Database("ctf").Collection("notifications")
	filter := bson.D{{"_id", docid}}
	var v *model.Notification
	err = collection.FindOne(ctx, filter).Decode(&v)
	MongoDisconnect(client)
	return v, nil
}

func get_notifications(ctx context.Context) ([]*model.Notification, error) {
	client, err := MongoConnect()
	if err != nil {
		// err1 := errors.New("connect error")
		MongoDisconnect(client)
		return nil, nil
	}
	collection := client.Database("ctf").Collection("notifications")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		// err1 := errors.New("")
		MongoDisconnect(client)
		return nil, nil
	}
	defer cursor.Close(context.TODO())
	var result []*model.Notification

	for cursor.Next(context.TODO()) {
		var v *model.Notification
		err := cursor.Decode(&v)
		if err != nil {
			Logger(err)
			MongoDisconnect(client)
			return nil, nil
		}

		result = append(result, v)

	}
	MongoDisconnect(client)
	return result, nil
}

func Updaterules(ctx context.Context, input *model.Manager) (string, error) {
	var admin bool
	admin = false
	if adminCheck(ctx) {
		admin = true
	}
	if !admin {
		err1 := errors.New("Who da heck are you")
		return "Who da heck are you??", err1
	}

	x := *input
	updata := map[string]interface{}{}

	if x.ScoreboardStats != nil {
		// 0 - only logged in
		// 1 - all
		// 2 - only admin

		ss := *x.ScoreboardStats
		if ss > 2 || ss < 0 {
			return "invalid response provided", nil
		}
		updata["ScoreboardStats"] = ss
	}

	if x.Status != nil {

		// 0 - Start/Resume CTF - allow flag submission, scoreboard changes
		// 1 - Stop CTF - hide all challenges, stop scoreboard, no flag check
		// 2 - Stop CTF - show existing challenges, stop scoreboard, allow flag check
		// 3 - Stop CTF - show existing challenges, stop scoreboard, no flag check

		stat := *x.Status
		if stat >= 4 || stat < 0 {
			return "invalid status provided", nil
		}

		updata["Status"] = stat
	}

	if x.Details != nil {
		updata["Details"] = *x.Details
	} else {
		updata["Details"] = "CTF stopped / paused / not started"
	}
	if updata == nil {
		return "no data provided", nil
	}
	client, err := MongoConnect()
	if err != nil {
		return "cannot connect to mongodb", nil
	}
	collection := client.Database("ctf").Collection("settings")
	exist := false
	listo, err := client.Database("ctf").ListCollectionNames(nil, bson.D{})
	for _, names := range listo {
		if names == "settings" {
			exist = true
		}
	}
	if !exist {
		_, erro := collection.InsertOne(
			context.Background(),
			&updata,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			return "Error updating", nil
		}
	} else {
		update := bson.M{"$set": updata}
		_, erro := collection.UpdateOne(
			context.Background(),
			bson.M{},
			update,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			return "Error updating", nil
		}
	}

	return "OK", nil
}

func viewRules(ctx context.Context) (*model.Managerial, error) {
	client, err := MongoConnect()
	if err != nil {
		return nil, errors.New("mongo error")
	}
	var v model.Managerial
	collectionsettings := client.Database("ctf").Collection("settings")
	err = collectionsettings.FindOne(ctx, bson.D{}).Decode(&v)
	return &v, nil
}
