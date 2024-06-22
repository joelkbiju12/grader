package school

import (
    "fmt"
    "encoding/json"
)

type School struct {
    ID        int
    Name      string
    Students  []Student
    Courses   []Course
    Classes   []Class
}

func NewSchool(id int, name) *School {
    return &School{
        ID:        id,
        Name:      name,
    }
}
func (s *School) AddStudent(student Student) {
    s.Students = append(s.Students, student)
}