package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) string {
	_, ok := l.tasks[task.Title]
	if ok {
		return taskExists
	}

	l.tasks[task.Title] = task
	return ""
}

func (l *List) ListTasks() map[string]Task {
	return l.tasks
}

func (l *List) DoneTask(title string) string {
	task, ok := l.tasks[title]
	if !ok {
		return taskNotFound
	}

	task.Done()
	l.tasks[title] = task

	return ""
}

func (l *List) DeleteTask(title string) string {
	_, ok := l.tasks[title]
	if !ok {
		return taskNotFound
	}

	delete(l.tasks, title)

	return ""
}
