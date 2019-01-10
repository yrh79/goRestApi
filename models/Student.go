package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type Student struct {
	Id int
	Name string
	Birthdate string
	Gender bool
	Score int
}

func GetAllStudents() []*Student {
	o := orm.NewOrm()
	o.Using("default")
	var students []*Student
	q:= o.QueryTable("student")
	q.All(&students)
	return students

}
func GetStudentById(id int) (Student, error){
	u:=Student{Id:id}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		return u, errors.New("Not found")
	} else if err == orm.ErrMissPK {
		return u, errors.New("Missing primary key")
	} else if err != nil {
		panic(err)
	}

	return u,nil
}
func AddStudent(student *Student) int{
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(student)
	return student.Id
}
func UpdateStudent(student *Student) {
	o := orm.NewOrm()
	o.Using("default")
	o.Update(student)
}

func DeleteStudent(id int){
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Student{Id:id})
}

func init() {
	orm.RegisterModel(new(Student))
}
