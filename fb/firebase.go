package fb

import (
	"context"
	"e-comm-auth/model"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
)

var client *auth.Client

func init() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		fmt.Println("Err:", err)
	}

	client, err = app.Auth(context.Background())
	if err != nil {
		fmt.Println("Err:", err)
	}

	fmt.Println("Firebase Instance Initialized!")
}

func CreateUser(model *model.UserSignupModel) error {
	user := &auth.UserToCreate{}
	user.Email(model.Email).Password(model.Password).EmailVerified(true).DisplayName(model.Name + " " + model.LastName).Disabled(false)

	_, err := client.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
