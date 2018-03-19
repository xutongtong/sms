package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"hbsms/forms"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"hbsms/utils"
)

type SmsCodeController struct {
	beego.Controller
}

// 发送短信验证码
func (c *SmsCodeController) Code() {
	sms := forms.SMS{}
	fmt.Println(sms)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &sms)
	if err != nil {
		log.Println(err)
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&sms)
	if err != nil {
		fmt.Println("ok" + err.Error())
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	length := 6
	utils.GenerateRandom(length)


}