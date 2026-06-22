package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang full course"
var rewardDessert = "Reward myself with a donut"

var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {

	fmt.Println("####### Welcome to our Todolist application ######")
	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/show-task", showTask)

	http.ListenAndServe(":8080", nil)

}

func helloUser(w http.ResponseWriter, r *http.Request) {
	var greeting = "Welcome to our Todolist application"
	fmt.Fprintln(w, greeting)
}

func showTask(w http.ResponseWriter, r *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(w, task)
	}
}

func printTask(taskItems []string) {

	fmt.Println("List of my Todos")
	for index, item := range taskItems {
		fmt.Printf("%d: %s\n", index+1, item)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
