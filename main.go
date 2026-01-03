package main

import "fmt"

func main() {
	todos:= Todos{}
	todos.add("Learn Go")
	todos.add("Build a Todo App")
	
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v\n", todos)
}
