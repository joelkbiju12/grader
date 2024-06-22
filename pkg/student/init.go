package student

import (
    "fmt"
    "encoding/json"
)

type Student struct {
    Name string
    Class string
    Courses []string
}


func (*Student) NewStudent(name string, class string, courses []string) *Student {
    return &Student{
        Name: name,
        Class: class,
        Courses: courses,
    }
}

func (s *Student) AddStudent(student Student) {
    studentList = append(studentList, student)
}

func (s *Student) GetStudent() {
    for _, student := range studentList {
        fmt.Println(student)
    }
}

func (s *Student) GetStudentByName(name string) {
    for _, student := range studentList {
        if student.Name == name {
            fmt.Println(student)
        }
    }
}



