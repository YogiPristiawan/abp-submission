package repo

import (
	"context"
	"database/sql"
	"todo/internal/todo/model"
)

func (r *Repo) GetById(ctx context.Context, id int64) (out model.GetByIdOut, err error) {
	if id == 0 {
		return model.GetByIdOut{}, sql.ErrNoRows
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return
	}

	var stmt = `
	SELECT
		id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at
	FROM
		todos
	WHERE
		id = ?
		AND
		deleted_at IS NULL`

	err = tx.QueryRowContext(ctx, stmt, id).Scan(
		&out.ID, &out.ActivityGroupID, &out.Title, &out.IsActive, &out.Priority,
		&out.CreatedAt, &out.UpdatedAt, &out.DeletedAt,
	)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.GetByIdOut{}, err2
		}

		return
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.GetByIdOut{}, err2
		}

		return
	}

	return
}
