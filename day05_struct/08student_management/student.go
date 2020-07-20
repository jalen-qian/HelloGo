package main

type student struct {
	id    int
	name  string
	age   int
	score int
}

/**
构造函数
*/
func newStudent(id int, name string, age int, score int) *student {
	return &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}
