package course

import (
    "fmt"
    "encoding/json"
)


type Course struct {
    ID          int
    Class      string
    School string
    Students []Student
}

func NewCourse(id int, class string, school string, students []Student) *Course {
    return &Course{
        ID: id,
        Class: class,
        School: school,
        Students: students,
    }
}

func (c *Course) AddStudent(s Student) {
    c.Students = append(c.Students, s)
}