package scanner

import (
	"fmt"
	"todo-app/todo"

	"github.com/k0kubun/pp"
)

func printResult(result string) {
	fmt.Println("Резльтат выполнения команды:", result)
	fmt.Println("")
}

func printPromt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы... До скорого!")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена.")
	fmt.Println("")
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
	fmt.Println("")
}

func printDone(title string) {
	fmt.Println("Задача '" + title + "' помечена как выполненная.")
	fmt.Println("")
}

func printDel(title string) {
	fmt.Println("Задача '" + title + "' успешно удалена.")
	fmt.Println("")
}

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("")
	fmt.Println("1. help")
	fmt.Println("—— узнать доступные команды")
	fmt.Println("")
	fmt.Println("2. add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов}")
	fmt.Println("—— добавить новую задачу")
	fmt.Println("")
	fmt.Println("3. list")
	fmt.Println("—— посмотреть полный список задач")
	fmt.Println("")
	fmt.Println("4. del {заголовок существующей задачи}")
	fmt.Println("—— удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("5. done {заголовок существующей задачи}")
	fmt.Println("—— отметить задачу как выполненную")
	fmt.Println("")
	fmt.Println("6. events")
	fmt.Println("—— получить список всех событий")
	fmt.Println("")
	fmt.Println("7. exit")
	fmt.Println("—— завершить выполнение программы")
	fmt.Println("")
}

func printEvents(events []Event) {
	pp.Println("События:", events)
}
