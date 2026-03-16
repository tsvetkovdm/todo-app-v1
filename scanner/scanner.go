package scanner

import (
	"bufio"
	"os"
	"strings"
	"todo-app/todo"
)

type Scanner struct {
	todoList *todo.List
	events   []Event
}

func NewScanner(todoList *todo.List) *Scanner {
	return &Scanner{
		todoList: todoList,
	}
}

func (s *Scanner) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPromt()

		ok := scanner.Scan()
		if !ok {
			return
		}

		inputString := scanner.Text()

		result := s.process(inputString)
		if result != "" {
			if result == needExit {
				printExit()
				return
			}

			printResult(result)
		}

		event := NewEvent(result, inputString)
		s.events = append(s.events, event)
	}
}

func (s *Scanner) process(inputString string) string {
	fields := strings.Fields(inputString)

	if len(fields) == 0 {
		return emptyInput
	}

	cmd := fields[0]

	if cmd == "exit" {
		return needExit
	}

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "done" {
		return s.cmdDone(fields)
	}

	if cmd == "del" {
		return s.cmdDel(fields)
	}

	if cmd == "help" {
		return s.cmdHelp(fields)
	}
	if cmd == "events" {
		return s.cmdEvents(fields)
	}

	return unknownCommand
}

func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}

	title := fields[1]

	taskText := ""

	for i := 2; i < len(fields); i++ {
		taskText += fields[i]

		if i != len(fields)-1 {
			taskText += " "
		}
	}

	task := todo.NewTask(title, taskText)

	addTaskResult := s.todoList.AddTask(task)
	if addTaskResult != "" {
		return addTaskResult
	}

	printAdd(title)

	return ""
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	tasks := s.todoList.ListTasks()

	printTasks(tasks)

	return ""
}

func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]

	doneTaskResult := s.todoList.DoneTask(title)
	if doneTaskResult != "" {
		return doneTaskResult
	}

	printDone(title)

	return ""
}

func (s *Scanner) cmdDel(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]

	delTaskResult := s.todoList.DeleteTask(title)
	if delTaskResult != "" {
		return delTaskResult
	}

	printDel(title)

	return ""
}

func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printHelp()

	return ""
}

func (s *Scanner) cmdEvents(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printEvents(s.events)

	return ""
}
