/*
 * @Descripttion: Do not edit
 * @version: v0.1.0
 * @Author: TSZ201
 * @Date: 2021-02-27 23:17:19
 * @LastEditors: TSZ201
 * @LastEditTime: 2021-02-27 23:17:20
 */
package control

import (
	"irisdemo/commons"
	"irisdemo/models"
	"irisdemo/service"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/lithammer/shortuuid/v3"
	"github.com/spf13/cast"
)

type StudentController struct {
	Ctx     iris.Context
	Service service.StudentService
}

func NewStudentController() *StudentController {
	return &StudentController{Service: service.NewStudentService()}
}

func (t StudentController) Get() (result mvc.Result) {
	if auth := commons.Auth(t.Ctx); !auth {
		t.Ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	if !t.Ctx.URLParamExists("limit") || !t.Ctx.URLParamExists("offset") {
		result = mvc.Response{Code: 400, Object: commons.ErrorsMap["NeedParams"]}
		return
	}

	limit := t.Ctx.URLParamEscape("limit")
	offset := t.Ctx.URLParamEscape("offset")
	res := t.Service.List(cast.ToUint(limit), cast.ToUint(offset))
	result = mvc.Response{Code: 200, Object: res.Data}
	return
}

func (t StudentController) Post() (result mvc.Result) {
	if auth := commons.Auth(t.Ctx); !auth {
		t.Ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	var student models.Students
	if err := t.Ctx.ReadJSON(&student); err != nil {
		log.Printf("Json Parse Error:%v", err)
		result = mvc.Response{Code: 400, Object: commons.ErrorsMap["JsonParseError"]}
		return
	}
	student.Id = shortuuid.New()
	res := t.Service.Create(student)
	if res.Flag {
		result = mvc.Response{Code: 200, Object: res.Data}
	} else {
		result = mvc.Response{Code: 400, Object: res.Data}
	}
	return
}

func (t StudentController) GetBy(id string) (result mvc.Result) {
	if auth := commons.Auth(t.Ctx); !auth {
		t.Ctx.StatusCode(iris.StatusUnauthorized)
		return
	}
	res := t.Service.Retrieve(id)
	if res.Flag {
		result = mvc.Response{Code: 200, Object: res.Data}
	} else {
		result = mvc.Response{Code: 400, Object: res.Data}
	}
	return
}

func (t StudentController) PutBy(id string) (result mvc.Result) {
	if auth := commons.Auth(t.Ctx); !auth {
		t.Ctx.StatusCode(iris.StatusUnauthorized)
		return
	}
	var student models.Students
	if err := t.Ctx.ReadJSON(&student); err != nil {
		log.Printf("Json Parse Error:%v", err)
		result = mvc.Response{Code: 400, Object: commons.ErrorsMap["JsonParseError"]}
		return
	}
	res := t.Service.Update(id, student)
	if res.Flag {
		result = mvc.Response{Code: 200, Object: res.Data}
	} else {
		result = mvc.Response{Code: 400, Object: res.Data}
	}
	return
}

func (t StudentController) DeleteBy(id string) (result mvc.Result) {
	if auth := commons.Auth(t.Ctx); !auth {
		t.Ctx.StatusCode(iris.StatusUnauthorized)
		return
	}
	res := t.Service.Destory(id)
	if res.Flag {
		result = mvc.Response{Code: 200, Object: res.Data}
	} else {
		result = mvc.Response{Code: 400, Object: res.Data}
	}
	return
}
