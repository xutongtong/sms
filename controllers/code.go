package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/validation"
	//"log"
	//"hbsms/utils"
	//"fmt"
	"hbsms/models"
	//"fmt"
	//"hbsms/models"
	//"hbsms/services"
	//"github.com/astaxie/beego/config"
	//"reflect"
	//"hbsms/utils"
	//"hbsms/services"
	"hbsms/services"
)

type SmsCodeController struct {
	beego.Controller
}

func (c *SmsCodeController) Test() {
	sms := models.SMS{
		CountryCode: "86",
		Mobile:      "18680663925",
		Source:      "hbtown",
		Template:    "register",
		Data:        map[string]interface{}{"code":"123456"},
	}
	st  := models.SMSTemplate{}
	st.Init(sms.Template)

	services.Send(sms, st)
}

// 发送短信验证码
func (c *SmsCodeController) Code() {
	//beego.LoadAppConfig("yaml", "conf/sms.yaml")
	//
	//fmt.Println(beego.AppConfig.String("goodsArrivalNotice::aliTemplateCode"))
	//fmt.Println("test")
	//fmt.Println(beego.AppConfig)
	//
	//cnf, err := config.NewConfig("yaml", "conf/sms.yaml")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//data, _ := cnf.DIY("goodsArrivalNotice")
	//fmt.Println(utils.GetYamlMap(data))
	//fmt.Println(utils.GetYamlVal("params", data))

	////cnf.Par
	//fmt.Println(cnf)
	//fmt.Println(cnf["register"])
	////fmt.Println(reflect.TypeOf(cnf))
	//fmt.Println("test")
	//fmt.Println(cnf.String("store_apply::aliTemplateCode"))
	//fmt.Println(cnf.String("store_apply::aliTemplateCode"))
	//fmt.Println(cnf.String("goodsArrivalNotice"))
	//fmt.Println(cnf.GetSection("goodsArrivalNotice"))
	//fmt.Println(cnf.``)
	//fmt.Println(cnf.String("goodsArrivalNotice"))



	// 表单数据结构
	//form := new(models.SmsCode)
	//err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(form)
	//services.SendSmsCode(*form)

	//
	//fmt.Println(form)
	//if err != nil {
	//	panic("Json format error")
	//}
	////校验参数
	//valid := validation.Validation{}
	//b, err := valid.Valid(&form)
	//if err != nil {
	//	// handle error
	//	panic("valid error")
	//}
	//if !b {
	//	// validation does not pass
	//	for _, err := range valid.Errors {
	//		log.Println(err.Key, err.Message)
	//	}
	//	panic("Form Validation Error")
	//}
	//codeTypes := []string{"register"}
	//if !utils.InArray(form.Type, codeTypes) {
	//	panic("register type is not true")
	//}
}