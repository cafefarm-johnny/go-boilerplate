package user

import (
	"context"
	"go-boilerplate/domain/model/user"
	"go-boilerplate/server/configure/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Dao struct {
}

func NewDao() *Dao {
	dao := new(Dao)

	return dao
}

func (dao *Dao) Find() ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	col := mongoDB.DB().Collection("user")
	filter := bson.M{}
	cursor, err := col.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cursor.Close(ctx)
	}()

	var users []user.User
	err = cursor.All(ctx, &users)

	return users, err
}

func (dao *Dao) FindOne(name string) (user.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	col := mongoDB.DB().Collection("user")
	filter := bson.M{"name": name}

	var u user.User
	err := col.FindOne(ctx, filter).Decode(&u)

	return u, err
}
