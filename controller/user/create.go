package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/handler"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/pkg/errno"
	"github.com/kaijian/gin-vue/utils"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Info("user create function called.", lager.Data{"X-Request-Id": utils.GentReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	//验证数据
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	//加密用户密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	//将用户插入数据库
	if err := u.Create(); err != nil {
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)
}
