package main

import (
	"fmt"
)

func main(){
	PrintLogo()
	storage := NewStorage[Todo]("todos.json")
	data, err := storage.Load()
	if err != nil {
		panic(err)
	}
	todos := Todos(data)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	if err := storage.Save(todos); err != nil {
		fmt.Println("Error saving todos:", err)
	}
}
