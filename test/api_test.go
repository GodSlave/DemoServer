package test

import (
	"testing"
	"github.com/GodSlave/MyGoServer/testbase"
	"github.com/GodSlave/MyGoServer/module/userModule"
	"github.com/GodSlave/MyGoServer/module/gate"
	"ServerImpl/src/module/sayHello"
)

func TestRegister(t *testing.T) {
	client := testbase.InitClient()
	register := &userModule.User_Register_Request{
		Username:   "TestUser",
		Password:   "UserPassword",
		VerifyCode: "aabbcc",
	}
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	message := testbase.BuildIPublishMessage(client, register, "User", "Register")
	testbase.SendMessageWithCheck(client, checkChan, message, t)
}

func TestLogin(t *testing.T) {

	client := testbase.InitClient()
	logIn := &userModule.User_Login_Request{
		Username: "TestUser",
		Password: "UserPassword",
	}
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	message := testbase.BuildIPublishMessage(client, logIn, "User", "Login")
	testbase.SendMessageWithCheck(client, checkChan, message, t)
}

func TestSayHello(t *testing.T) {
	client := testbase.InitClient()
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	message := testbase.BuildIPublishMessage(client, nil, "Hello", "SayHello")
	testbase.SendMessageWithCheck(client, checkChan, message, t)

}

func TestSayHelloWithSession(t *testing.T) {
	client := testbase.InitClient()
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	message := testbase.BuildIPublishMessage(client, nil, "Hello", "UserSayHello")
	testbase.SendMessageWithCheck(client, checkChan, message, t)
}

func TestLoginSayHello(t *testing.T) {
	client := testbase.InitClient()
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	logIn := &userModule.User_Login_Request{
		Username: "TestUser",
		Password: "UserPassword",
	}
	loginMessage := testbase.BuildIPublishMessage(client, logIn, "User", "Login")
	testbase.SendMessageWithCheck(client, checkChan, loginMessage, t)
	message := testbase.BuildIPublishMessage(client, nil, "Hello", "LoginUserSayHello")
	testbase.SendMessageWithCheck(client, checkChan, message, t)
}

func TestSayHelloWithParam(t *testing.T) {
	client := testbase.InitClient()
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	param := &sayHello.Hello_SayHello_Request{
		Param: "Hi,Server",
	}

	message := testbase.BuildIPublishMessage(client, param, "Hello", "SayHelloWithParam")
	testbase.SendMessageWithCheck(client, checkChan, message, t)
}
func TestLoginSayHelloWithParam(t *testing.T) {
	client := testbase.InitClient()
	checkChan := make(chan *gate.AllResponse)
	testbase.SubI(client, checkChan)
	logIn := &userModule.User_Login_Request{
		Username: "TestUser",
		Password: "UserPassword",
	}
	loginMessage := testbase.BuildIPublishMessage(client, logIn, "User", "Login")
	testbase.SendMessageWithCheck(client, checkChan, loginMessage, t)
	param := &sayHello.Hello_SayHello_Request{
		Param: "Hi,Server",
	}
	message := testbase.BuildIPublishMessage(client, param, "Hello", "LogInSayHelloWithParam")
	testbase.SendMessageWithCheck(client, checkChan, message, t)
}
