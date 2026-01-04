package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Toggle int
	Edit   string
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := &CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by its index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the completion status of a todo by its index")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit the title of a todo by its index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":",2)
		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Use index:new title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err!= nil{
			fmt.Println("Error:Invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println(("Invalid command"))
	}
}
