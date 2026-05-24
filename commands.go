package main

import (
	"bufio"
	"fmt"
	"flag"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
	Interactive bool
}

func NewCmdFlags() *cmdFlags {
	cf := &cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and new text (format: index:new text)")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the completion status of a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.Interactive, "i", false, "Run in interactive mode")

	flag.Usage = func() {
		fmt.Printf("Usage of KAMI:\n")
		flag.PrintDefaults()
	}
	
	flag.Parse()
	return cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	// Check if any flag was provided (besides interactive)
	hasFlag := cf.Add != "" || cf.Del >= 0 || cf.Edit != "" || cf.Toggle >= 0 || cf.List

	if !hasFlag || cf.Interactive {
		cf.RunREPL(todos)
		return
	}

	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Del >= 0:
		if err := todos.Delete(cf.Del); err != nil {
			fmt.Println("Error deleting todo:", err)
		}
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit format. Use: -edit index:new text")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index for edit.")
			return
		}
		if err := todos.Edit(index, parts[1]); err != nil {
			fmt.Println("Error editing todo:", err)
		}
	case cf.Toggle >= 0:
		if err := todos.Toggle(cf.Toggle); err != nil {
			fmt.Println("Error toggling todo:", err)
		}
	}
}

func (cf *cmdFlags) RunREPL(todos *Todos) {
	scanner := bufio.NewScanner(os.Stdin)
	// Using the lighter color from the right side of the logo (255;177;153)
	fmt.Println("\x1b[38;2;255;177;153mWelcome to KAMI! Type /help for commands, /exit to quit.\x1b[0m")

	for {
		// Using an orange-red from the logo's middle section (255;69;99) for the prompt
		fmt.Print("\x1b[38;2;255;69;99mKAMI > \x1b[0m")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "/help":
			fmt.Println("Available commands: /list, /add <task>, /toggle <index>, /edit <index> <text>, /del <index>, /exit")
		case "/exit", "/quit":
			fmt.Println("Goodbye!")
			return
		case "/list":
			todos.Print()
		case "/add":
			if len(args) == 0 {
				fmt.Println("Usage: /add <task>")
				continue
			}
			todos.Add(strings.Join(args, " "))
		case "/toggle":
			if len(args) == 0 {
				fmt.Println("Usage: /toggle <index>")
				continue
			}
			idx, _ := strconv.Atoi(args[0])
			todos.Toggle(idx)
		case "/del":
			if len(args) == 0 {
				fmt.Println("Usage: /del <index>")
				continue
			}
			idx, _ := strconv.Atoi(args[0])
			todos.Delete(idx)
		case "/edit":
			if len(args) < 2 {
				fmt.Println("Usage: /edit <index> <new text>")
				continue
			}
			idx, _ := strconv.Atoi(args[0])
			todos.Edit(idx, strings.Join(args[1:], " "))
		default:
			fmt.Printf("Unknown command: %s (Try /help)\n", command)
		}
	}
}
