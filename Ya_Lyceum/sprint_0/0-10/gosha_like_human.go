package main

import (
	"fmt"
	"time"
)

type Note struct {
	title string
	text  string
}

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (todo ToDoList) TasksCount() int {
	return len(todo.tasks)
}

func (todo ToDoList) NotesCount() int {
	return len(todo.notes)
}

func (todo ToDoList) CountTopPrioritiesTasks() int {
	cnt := 0
	for _, el := range todo.tasks {
		if el.IsTopPriority() {
			cnt++
		}
	}
	return cnt
}

func (todo ToDoList) CountOverdueTasks() int {
	cnt := 0
	for _, el := range todo.tasks {
		if el.IsOverdue() {
			cnt++
		}
	}
	return cnt
}

func (t Task) IsOverdue() bool {
	return t.deadline.Before(time.Now())
}

func (t Task) IsTopPriority() bool {
	return t.priority >= 4
}

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	return s.passingScore <= float64(s.solvedProblems)*s.scoreForOneTask
}

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\n", e.name)
	fmt.Printf("Position: %s\n", e.position)
	var totalSalary = e.salary + e.bonus
	fmt.Printf("Total Salary: %.2f\n", totalSalary)
}

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Address: %s\n", p.address)
}

func main() {
	todo := ToDoList{name: "Gosha ToDo list", tasks: []Task{Task{summary: "Make Yandex Lyceum Task 9", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 9", priority: 5}, Task{summary: "Make Yandex Lyceum Task 10", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 3}}, notes: []Note{Note{title: "ToDo list task", text: "ToDo list task in Yandex Lyceum is very interesting"}}}
	fmt.Println(todo.TasksCount())
	fmt.Println(todo.NotesCount())
	fmt.Println(todo.CountTopPrioritiesTasks())
	fmt.Print(todo.CountOverdueTasks())
}
