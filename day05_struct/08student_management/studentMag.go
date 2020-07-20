package main

import "fmt"

type studentMag struct {
	allStudents []*student
}

/**
构造函数
*/
func newStudentMag() *studentMag {
	return &studentMag{
		allStudents: make([]*student, 0, 100),
	}
}

/**
添加学生信息
*/
func (s *studentMag) addStudent(student *student) {
	s.allStudents = append(s.allStudents, student)
	fmt.Println("添加学生成功！")
}

/**
编辑学生信息
*/
func (s *studentMag) editStudent(student *student) {
	for i, stu := range s.allStudents {
		if stu.id == student.id {
			s.allStudents[i] = student
			return
		}
	}
	fmt.Printf("编辑失败，学号为%d的学生不存在\n", student.id)
}

/**
根据学号删除学生
*/
func (s *studentMag) deleteStudent(id int) {
	//根据学号找到索引
	for i, stu := range s.allStudents {
		if stu.id == id {
			//找到了学生，并且知道索引是i，执行删除
			s.allStudents = append(s.allStudents[:i], s.allStudents[(i + 1):]...)
			return
		}
	}
	fmt.Printf("删除失败，没有找到学号为【%d】的学生\n", id)
}

/**
展示所有学生信息
*/
func (s *studentMag) showAllStudents() {
	fmt.Println("######################################")
	for _, stu := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d \n", stu.id, stu.name, stu.age, stu.score)
	}
	fmt.Println("######################################")
}
