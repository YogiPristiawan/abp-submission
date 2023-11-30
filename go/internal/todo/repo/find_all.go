package repo

import (
	"context"
	"database/sql"
	"errors"
	"todo/internal/todo/model"
)

func (r *Repo) FindAll(ctx context.Context, activityGroupId ...int64) (out []model.FindAllOut, err error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return
	}

	var stmt = `
	SELECT
		id, activity_group_id, title, is_active, priority, created_at, updated_at,
		deleted_at
	FROM
		todos
	WHERE
		deleted_at IS NULL`

	var binding = make([]any, len(activityGroupId))
	if len(activityGroupId) > 0 {
		stmt += `AND activity_group_id = ?`
		binding = append(binding, activityGroupId[0])
	}

	rows, err := tx.QueryContext(ctx, stmt, binding...)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			err2 := tx.Rollback()
			if err2 != nil {
				return []model.FindAllOut{}, err2
			}
		}

		return []model.FindAllOut{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.FindAllOut
		err = rows.Scan(
			&todo.ID, &todo.ActivityGroupID, &todo.Title, &todo.IsActive,
			&todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt,
		)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				return []model.FindAllOut{}, err2
			}

			return
		}

		out = append(out, todo)
	}

	return
}
