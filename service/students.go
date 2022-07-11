/*
 * @Descripttion: Do not edit
 * @version: v0.1.0
 * @Author: TSZ201
 * @Date: 2021-02-27 23:17:19
 * @LastEditors: TSZ201
 * @LastEditTime: 2021-02-27 23:17:20
 */
package service

import (
	"irisdemo/commons"
	"irisdemo/models"

	"github.com/prometheus/common/log"
)

type StudentService interface {
	List(limit uint, offset uint) (result models.ResBody)
	Create(student models.Students) (result models.ResBody)
	Update(id string, student models.Students) (result models.ResBody)
	Retrieve(id string) (result models.ResBody)
	Destory(id string) (result models.ResBody)
}

type studentService struct{}

func NewStudentService() *studentService {
	return &studentService{}
}

var studentEntity = models.NewStudentEntity()

func (t *studentService) List(limit uint, offset uint) (result models.ResBody) {
	total, students := studentEntity.StudentList(limit, offset)
	maps := make(map[string]interface{}, 10)
	maps["Total"] = total
	maps["List"] = students
	result.Data = maps
	result.Flag = true
	return
}

func (t *studentService) Create(student models.Students) (result models.ResBody) {
	var class models.Classes
	var teacher models.Teachers
	var err error
	class, err = classEntity.GetClass(student.Class.Id)
	if err != nil {
		result.Flag = false
		result.Data = commons.ErrorsMap["MissEntity"]
		return
	}
	teacher, err = teacherEntity.GetTeacher(student.Teacher.Id)
	if err != nil {
		result.Flag = false
		result.Data = commons.ErrorsMap["MissEntity"]
		return
	}

	student.Class = class
	student.Teacher = teacher

	res := studentEntity.NameDuplicated("", student.Name)
	if res {
		result.Flag = false
		result.Data = commons.ErrorsMap["NameDuplicated"]
		return
	}

	err = studentEntity.AddStudent(student)
	if err != nil {
		result.Flag = false
		result.Data = commons.ErrorsMap["InsertError"]
	} else {
		result.Flag = true
		result.Data = student
	}
	return
}

func (t *studentService) Update(id string, student models.Students) (result models.ResBody) {
	res := studentEntity.NameDuplicated(id, student.Name)
	if res {
		result.Flag = false
		result.Data = commons.ErrorsMap["NameDuplicated"]
		return
	}
	err := studentEntity.UpdateStudent(id, student)
	if err != nil {
		result.Flag = false
		result.Data = commons.ErrorsMap["UpdateError"]
		log.Errorf("Update student Error:%v", err)
	} else {
		result.Flag = true
		student.Id = id
		result.Data = student
	}
	return
}

func (n *studentService) Retrieve(id string) (result models.ResBody) {
	var student models.Students
	var err error
	student, err = studentEntity.GetStudent(id)
	if err != nil {
		result.Flag = false
		result.Data = commons.ErrorsMap["RetrieveError"]
	} else {
		result.Flag = true
		result.Data = student
	}
	return
}

func (t *studentService) Destory(id string) (result models.ResBody) {
	err := studentEntity.DelStudent(id)
	if err != nil {
		result.Flag = false
		result.Data = commons.ErrorsMap["DeleteError"]
	} else {
		result.Flag = true
		result.Data = models.Students{}
	}
	return
}
