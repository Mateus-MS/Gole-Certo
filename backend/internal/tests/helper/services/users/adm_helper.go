package test_helper_users

import (
	"testing"

	user_model "alves.com/modules/users/model"
	user_service "alves.com/modules/users/service"
	"alves.com/pkg/security"
	"go.mongodb.org/mongo-driver/mongo"
)

var temp_name string = "temp_name"
var temp_pass string = "temp_pass"

func LoginTempADM(t *testing.T, usersService user_service.IService) string {
	accessToken, _ := usersService.Login(t.Context(), temp_name, temp_pass)

	return accessToken
}

func RegisterTempADM(t *testing.T, usersCollection *mongo.Collection) {
	hashedPassword, _ := security.HashPassword(temp_pass)
	adm_user := user_model.NewUser(temp_name, hashedPassword)
	adm_user.IsAdmin = true

	usersCollection.InsertOne(t.Context(), adm_user)
}
