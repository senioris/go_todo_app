package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/senioris/go_todo_app/entity"
)

func (r *Repository) RegisterUesr(ctx context.Context, db Execer, u *entity.User) error {
	u.Created = r.Clocker.Now()
	u.Modified = u.Created

	sql := `INSERT INTO user(
		name, password, role, created, modified
	) VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, u.Name, u.Password, u.Role, u.Created, u.Modified)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLDuplicateEntry {
			return fmt.Errorf("cannot craete same name user: %w", ErrAlreadyEntry)
		}
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = entity.UserID(id)
	return nil
}
