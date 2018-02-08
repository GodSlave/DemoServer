package sayHello

import (
	"github.com/GodSlave/MyGoServer/module/base"
	"github.com/GodSlave/MyGoServer/module"
	"github.com/GodSlave/MyGoServer/conf"
	"github.com/GodSlave/MyGoServer/base"
	"fmt"
)

type HelloModule struct {
	basemodule.BaseModule
}

var Module = func() module.Module {
	newHelloModule := new(HelloModule)
	return newHelloModule
}

func (m *HelloModule) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
	m.GetServer().RegisterGO("SayHello", 1, m.SayHello)
	m.GetServer().RegisterGO("UserSayHello", 2, m.UserSayHello)
	m.GetServer().RegisterGO("LoginUserSayHello", 3, m.LoginUserSayHello)
	m.GetServer().RegisterGO("SayHelloWithParam", 4, m.SayHelloWithParam)
	m.GetServer().RegisterGO("LogInSayHelloWithParam", 5, m.LogInSayHelloWithParam)
}

func (m *HelloModule) Version() string {
	return "1.0.1"
}

func (m *HelloModule) GetType() string {
	return "Hello"
}

func (m *HelloModule) Run(closeSig chan bool) {

	<-closeSig
}

/**
 * 公开接口，不需要用户登陆
 */
func (m *HelloModule) SayHello(sessionId string) (*Hello_SayHello_Response, *base.ErrorCode) {
	return &Hello_SayHello_Response{"Hello"}, base.ErrNil
}

/**
 * 手动验证用户信息
 */
func (m *HelloModule) UserSayHello(sessionId string) (*Hello_SayHello_Response, *base.ErrorCode) {
	return &Hello_SayHello_Response{fmt.Sprintf("Hello %s", sessionId)}, base.ErrNil
}

/**
 * 只有登陆过的用户才可以调用
 */
func (m *HelloModule) LoginUserSayHello(user *base.BaseUser) (*Hello_SayHello_Response, *base.ErrorCode) {
	return &Hello_SayHello_Response{fmt.Sprintf("Hello %s %s", user.Name, user.UserID)}, base.ErrNil
}

/**
 * 带参数的公开接口
 */
func (m *HelloModule) SayHelloWithParam(form *Hello_SayHello_Request) (*Hello_SayHello_Response, *base.ErrorCode) {
	return &Hello_SayHello_Response{fmt.Sprintf("Hello %s", form.Param)}, base.ErrNil
}

/**
 * 公开接口,传递参数
 */
func (m *HelloModule) LogInSayHelloWithParam(user *base.BaseUser, form *Hello_SayHello_Request) (*Hello_SayHello_Response, *base.ErrorCode) {
	return &Hello_SayHello_Response{fmt.Sprintf("Hello %s %s", user.UserID, form.Param)}, base.ErrNil
}
