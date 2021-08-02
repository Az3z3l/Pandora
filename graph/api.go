package graph

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddChallengeFile api to upload challenge files
func AddChallengeFile(response http.ResponseWriter, request *http.Request) {

	challid := request.FormValue("id")
	docid, err := primitive.ObjectIDFromHex(challid)
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Challenge ID error" }`))
		return
	}

	file, handler, err := request.FormFile("file")
	fileName := request.FormValue("file_name")
	fileName = handler.Filename
	if err != nil {
		log.Panic(err)
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Could not get the file" }`))
		return
	}

	defer file.Close()
	// challenges will have the files grouped by the challenge id
	folder := "challenges/"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0700)
		f, err := os.OpenFile("challenges/index.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Panic(err)
			f.Close()
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Status": "Error on creating folder" }`))
			return
		}
		f.Close()

	}
	folder = "challenges/" + challid + "/"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0700)
	}

	f, err := os.OpenFile(folder+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Unable to save the file" }`))
		log.Panic(err)
		return

	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	client, err := MongoConnect()
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Could'nt connect to db" }`))
		log.Panic(err)
		return
	}
	collection := client.Database("ctf").Collection("challenges")
	filter := bson.D{{"_id", docid}}
	var chall *model.Challenge
	err = collection.FindOne(context.Background(), filter).Decode(&chall)
	if err != nil {
		MongoDisconnect(client)
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Unable to find the Challenge" }`))
		log.Panic(err)
		return
	}

	if chall.File == nil {
		update := bson.M{"$set": bson.M{"file": []string{fileName}}}
		_, erro := collection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Status": "Unable to add chall to db" }`))
			log.Panic(err)
			return
		}
	} else {
		for _, i := range chall.File {
			if i == fileName {
				MongoDisconnect(client)
			}
		}
	}

	if chall.File != nil {
		update := bson.M{"$push": bson.M{"file": fileName}}
		_, erro := collection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			Logger(erro)
			MongoDisconnect(client)
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Status": "Unable to add chall to db" }`))
			log.Panic(err)
			return
		}
	}
	response.Header().Set("content-type", "application/json")
	response.Write([]byte(`{ "Status": "OK" }`))
	return
}

func UserPhoto(response http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(1 << 20)
	ctx := request.Context()
	userid := ForContext(ctx)

	_, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Invalid User ID" }`))
		return
	}

	clientFile, _, _ := request.FormFile("file")
	defer clientFile.Close()

	//
	b, err := ioutil.ReadAll(clientFile)
	if err != nil {
		panic(err)
	}

	reader1 := bytes.NewReader(b)
	var buffer bytes.Buffer
	fileSize, err := buffer.ReadFrom(reader1)

	if fileSize/1024 > 500 {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Size greater than 400kb" }`))
		return
	}

	reader2 := bytes.NewReader(b)
	buff := make([]byte, 512)
	if _, err := reader2.Read(buff); err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Error determining File" }`))
		return
	}
	//
	contentType := http.DetectContentType(buff)

	extAllowed := []string{"image/jpeg", "image/gif", "image/png"}
	valid := false
	extension := "xoxo"
	for _, s := range extAllowed {
		valid = strings.EqualFold(contentType, s)
		if valid {
			valid = true
			extension = userid + "." + strings.Split(s, "/")[1]
			break
		}
	}
	if !valid {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Unsupported File" }`))
		return
	}

	fmt.Println(extension)

	folder := "userphoto/"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0700)
		f, err := os.OpenFile("userphoto/index.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Panic(err)
			f.Close()
			response.Header().Set("content-type", "application/json")
			response.Write([]byte(`{ "Status": "Error on creating folder" }`))
			return
		}
		f.Close()

	}

	files, err := filepath.Glob(folder + "/" + userid + "*")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(folder+"/"+extension, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		response.Header().Set("content-type", "application/json")
		response.Write([]byte(`{ "Status": "Unable to save the file" }`))
		log.Panic(err)
		return

	}
	defer f.Close()
	reader3 := bytes.NewReader(b)
	_, _ = io.Copy(f, reader3)

	response.Header().Set("content-type", "application/json")
	response.Write([]byte(`{ "Status": "OK" }`))
	return
}
