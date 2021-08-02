package model

import "time"

type Challenge struct {
	ID          string       `bson:"_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Category    []*string    `json:"category"`
	Tags        []*string    `json:"tags"`
	Value       int          `json:"value"`
	Flags       string       `json:"flags"`
	Solves      int          `json:"solves"`
	Teamssolved []*Challwhen `json:"teamssolved"`
	Visibility  bool         `json:"visibility"`
	File        []string     `json:"file"`
}

type User struct {
	ID        string      `bson:"_id"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	Email     string      `json:"email"`
	Solved    []*Userwhen `json:"solved"`
	Score     int         `json:"score"`
	Lastsolve *time.Time  `json:"lastsolve"`
}

type NoobUser struct {
	ID        string      `bson:"_id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Solved    []*Userwhen `json:"solved"`
	Score     int         `json:"score"`
	Lastsolve *time.Time  `json:"lastsolve"`
}

type Challwhen struct {
	User      string    `bson:"UserID"`
	TimeStamp time.Time `json:"TimeStamp"`
}

type Userwhen struct {
	Challenge string    `bson:"ChallID"`
	Timestamp time.Time `json:"Timestamp"`
}

type Fulluser struct {
	Username    string      `json:"username"`
	Fullname    string      `json:"fullname"`
	Email       string      `json:"email"`
	Solved      []*Userwhen `json:"solved"`
	Score       int         `json:"score"`
	Lastsolve   *time.Time  `json:"lastsolve"`
	Age         string      `json:"age"`
	Gender      string      `json:"gender"`
	Institution string      `json:"institution"`
	Contact     string      `json:"contact"`
	Place       string      `json:"place"`
	District    string      `json:"district"`
	State       string      `json:"state"`
}

type Passwdreset struct {
	ID       string `bson:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Notification struct {
	ID          string    `bson:"_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}
