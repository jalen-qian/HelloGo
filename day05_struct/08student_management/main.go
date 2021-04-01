package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.修改学员信息")
	fmt.Println("3.显示所有学员信息")
	fmt.Println("4.删除学员")
	fmt.Println("5.退出系统")
}

func getInput() *student {
	var (
		id, age, score int
		name           string
	)

tId:
	fmt.Print("请输入学号：")
	_, err := fmt.Scanf("%d\n", &id)
	if err != nil {
		fmt.Println("输入有误，学号必须是数字", err)
		goto tId
	}

tName:
	fmt.Print("请输入姓名：")
	_, errName := fmt.Scanf("%s\n", &name)
	if errName != nil {
		fmt.Println("输入有误，姓名必须是字符串", errName)
		goto tName
	}

tAge:
	fmt.Print("请输入年龄：")
	_, errAge := fmt.Scanf("%d\n", &age)
	if errAge != nil {
		fmt.Println("输入有误，年龄必须是数字", errAge)
		goto tAge
	}

tScore:
	fmt.Print("请输入分数：")
	_, errScore := fmt.Scanf("%d\n", &score)
	if errScore != nil {
		fmt.Println("输入有误，分数必须是数字", errScore)
		goto tScore
	}
	return newStudent(id, name, age, score)
}

func main() {
	stuMag := newStudentMag()
	for {
		showMenu()
		var input int
		fmt.Print("请输入菜单编号：")
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			fmt.Println("\n输入错误，请输入数字...")
			continue
		}
		fmt.Printf("您选择了%d\n", input)
		switch input {
		case 1:
			// 添加学员
			stuMag.addStudent(getInput())
		case 2:
			// 修改学员
			stuMag.editStudent(getInput())
		case 3:
			// 显示所有学员
			stuMag.showAllStudents()
		case 4:
			// 删除学员
			var id int
		tId:
			fmt.Print("请输入学号：")
			_, err := fmt.Scanf("%d\n", &id)
			if err != nil {
				fmt.Println("输入有误，学号必须是数字", err)
				goto tId
			}
			stuMag.deleteStudent(id)
		case 5:
			// 退出系统
			fmt.Println("欢迎您再次使用")
			os.Exit(0)

		}
	}
}
