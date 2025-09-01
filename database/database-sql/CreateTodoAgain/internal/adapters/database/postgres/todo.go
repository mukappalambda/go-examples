package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mukappalambda/go-examples/examples/database/database-sql/CreateTodoAgain/internal/core/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

var _ domain.TodoRepository = (*PostgresRepository)(nil)

func New(dataSourceName string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	repo := &PostgresRepository{db: db}
	return repo, nil
}

func (r *PostgresRepository) CreateTodo(todo *domain.Todo) error {
	return nil
}

func (r *PostgresRepository) GetTodoByName(name string) (*domain.Todo, error) {
	return nil, nil
}

func (r *PostgresRepository) DeleteTodoByName(name string) error {
	return nil
}
