/*
 * @Descripttion: Do not edit
 * @version: v0.1.0
 * @Author: TSZ201
 * @Date: 2021-02-27 23:17:19
 * @LastEditors: TSZ201
 * @LastEditTime: 2021-02-27 23:17:20
 */
package models

import (
	"irisdemo/commons"
	"log"
)

type Students struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Phone     string
	TeacherId string
	ClassId   string
	Teacher   Teachers `gorm:"foreignKey:teachers(id);references:TeacherId"`
	Class     Classes  `gorm:"foreignKey:classes(id);references:ClassId"`
}

type studentEntity struct{}

func NewStudentEntity() *studentEntity {
	return &studentEntity{}
}

type IStudentEntity interface {
	StudentList(limit uint, offset uint) (total int, Students []Students)
	AddStudent(Students Students) (err error)
	UpdateStudent(id, string, Students Students) (err error)
	GetStudent(id string) (Students Students, err error)
	DelStudent(id string) (err error)

	NameDuplicated(id, string, name string) (result bool)
}

func (t *studentEntity) StudentList(limit uint, offset uint) (total int, students []Students) {
	commons.DbClient.Model(Students{}).Count(&total)
	err := commons.DbClient.Limit(limit).Offset(offset).Find(&students).Error
	if err != nil {
		log.Fatalf("List Students error:%v\n", err)
		panic("select Error")
	}
	return
}

func (t *studentEntity) AddStudent(student Students) (err error) {
	return commons.DbClient.Omit("Teacher").Omit("Class").Create(&student).Error
}

func (t *studentEntity) UpdateStudent(id string, Student Students) (err error) {
	return commons.DbClient.Model(Students{}).Where("id=?", id).Update(&Student).Error
}

func (t *studentEntity) GetStudent(id string) (Student Students, err error) {
	err = commons.DbClient.Find(&Student, "id=?", id).Error
	return
}

func (t *studentEntity) NameDuplicated(id string, name string) (result bool) {
	var count int32
	if id == "" {
		commons.DbClient.Model(Students{}).Where("name=?", name).Count(&count)
	} else {
		commons.DbClient.Model(Students{}).Where("id <>? and name=? ", id, name).Count(&count)
	}
	if count > 0 {
		result = true
	} else {
		result = false
	}
	return
}

func (t studentEntity) DelStudent(id string) (err error) {
	var Student Students
	Student.Id = id
	err = commons.DbClient.Delete(&Student).Error
	return
}
