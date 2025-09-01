package domain

type TodoRepository interface {
	CreateTodo(todo *Todo) error
	GetTodoByName(name string) (*Todo, error)
	DeleteTodoByName(name string) error
}
