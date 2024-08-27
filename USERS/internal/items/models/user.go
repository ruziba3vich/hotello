package models

import (
	"github.com/ruziba3vich/hotello-users/genprotos/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		Id       primitive.ObjectID `bson:"_id"`
		FullName string             `bson:"full_name"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
		Email    string             `bson:"email"`
		Deleted  string             `bson:"deleted"`
	}

	UpdateUserRequest struct {
		FieldName string
		Value     string
	}
)

func (u *User) ToProto() *users.User {
	return &users.User{
		Id:       u.Id.Hex(),
		FullName: u.FullName,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Deleted:  u.Deleted,
	}
}

func (u *User) FromProto(obj *users.User) error {
	id, err := primitive.ObjectIDFromHex(obj.Id)
	if err != nil {
		return err
	}
	u.Id = id
	u.FullName = obj.FullName
	u.Username = obj.Username
	u.Password = obj.Password
	u.Email = obj.Email
	u.Deleted = obj.Deleted
	return nil
}

func (u *User) FromCreateUserRequest(obj *users.CreateUserRequest) {
	u.Id = primitive.NewObjectID()
	u.FullName = obj.FullName
	u.Username = obj.Username
	u.Password = obj.Password
	u.Email = obj.Email
}
