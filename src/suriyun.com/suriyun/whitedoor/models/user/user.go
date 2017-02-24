package user

import (
	"time"

	"github.com/satori/go.uuid"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"suriyun.com/suriyun/whitedoor/databases/mongodb"
)

// User ... User database model
type User struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	GuestDeviceID string        `json:"guestDeviceId,omitempty" bson:"guestDeviceId,omitempty"`
	Username      string        `json:"username,omitempty" bson:"username,omitempty"`
	Password      string        `json:"password,omitempty" bson:"password,omitempty"`
	Email         string        `json:"email,omitempty" bson:"email,omitempty"`
	LoginToken    string        `json:"loginToken,omitempty" bson:"loginToken,omitempty"`
	FacebookID    string        `json:"facebookId,omitempty" bson:"facebookId,omitempty"`
	FacebookToken string        `json:"facebookToken,omitempty" bson:"facebookToken,omitempty"`
	GoogleID      string        `json:"googleId,omitempty" bson:"googleId,omitempty"`
	GoogleToken   string        `json:"googleToken,omitempty" bson:"googleToken,omitempty"`
	RegisterDate  *time.Time    `json:"registerDate,omitempty" bson:"registerDate,omitempty"`
	LastLoginDate *time.Time    `json:"lastLoginDate,omitempty" bson:"lastLoginDate,omitempty"`
}

// New ... Create new User with required database
func New() User {
	user := User{}
	user.ID = bson.NewObjectId()
	registerDate := time.Now()
	user.RegisterDate = &registerDate
	user.UpdateLogin()
	return user
}

// NewEmpty ... Create new User with empty database
func NewEmpty() User {
	user := User{}
	user.RegisterDate = nil
	user.LastLoginDate = nil
	return user
}

// GetCollection ... Get database collection made it easily to access table
func GetCollection() *mgo.Collection {
	return mongodb.GetCollection("users")
}

// GenerateLoginToken ... Generate new login token
func GenerateLoginToken() string {
	return uuid.NewV4().String()
}

// UpdateLogin ... Update login token and last login date
func (user *User) UpdateLogin() {
	user.LoginToken = GenerateLoginToken()
	lastLoginDate := time.Now()
	user.LastLoginDate = &lastLoginDate
}
