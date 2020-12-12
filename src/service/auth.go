package service

 import (
	 "news/utils/exception"
	 "news/utils/logging"
	 "news/utils"
	 "news/serializer"
 )

 // AuthService 媒体列表服务
type AuthService struct {
	Name string `form:"name" json:"name"`
}
 
 // Login 用户登录函数
 func (service *AuthService) Auth() serializer.Response {
	code := exception.SUCCESS
	token, err := utils.GenerateToken(service.Name)
	if err != nil {
		logging.Info(err)
		code = exception.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
		}
	}
	return serializer.Response{
		Data:   serializer.TokenData{Token: token},
		Status: code,
		Msg:    exception.GetMsg(code),
	}
 }