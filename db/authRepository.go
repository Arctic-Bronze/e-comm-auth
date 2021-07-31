package db

import (
	"e-comm-auth/fb"
	"e-comm-auth/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserExists(email string) bool {
	repo := GetCollection("user")
	ctx, _ := GetCtx()
	number, err := repo.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false
	}
	if number == 0 {
		return false
	}
	return true
}

func InsertUser(user *model.User, userModel *model.UserSignupModel) error {
	ctx, _ := GetCtx()

	//session, err := client.StartSession();
	//if err != nil {
	//	return &fiber.Error{Code: 500, Message: "User could not signed-up!"}
	//}
	//if err = session.StartTransaction(); err != nil {
	//	return &fiber.Error{Code: 500, Message: "User could not signed-up!"}
	//}

	err := client.UseSession(ctx, func(sc mongo.SessionContext) error {
		err := sc.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = GetCollection("user").InsertOne(sc, user)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = fb.CreateUser(userModel)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if err = sc.CommitTransaction(sc); err != nil {
			fmt.Println(err)
			return err
		}
		sc.EndSession(ctx)
		return nil
	})

	return err

}
