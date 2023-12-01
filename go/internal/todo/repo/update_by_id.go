package repo

import (
	"context"
	"database/sql"
	"time"
	"todo/internal/todo/model"
)

func (r *Repo) UpdateById(ctx context.Context, id int64, data model.UpdateByIdIn) (out model.UpdateByIdOut, err error) {
	if id == 0 {
		return model.UpdateByIdOut{}, sql.ErrNoRows
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return
	}

	now := time.Now().Unix()
	var stmt = `
	UPDATE todos SET
		is_active = ?, updated_at = ?
	WHERE
		id = ?
		AND
		deleted_at IS NULL`

	commandTag, err := tx.ExecContext(ctx, stmt, data.IsActive, now, id)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.UpdateByIdOut{}, err2
		}

		return
	}
	affected, err := commandTag.RowsAffected()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.UpdateByIdOut{}, err2
		}

		return
	}
	if affected == 0 {
		return model.UpdateByIdOut{}, sql.ErrNoRows
	}

	var selectStmt = `
	SELECT
		id, title, activity_group_id, is_active, priority, created_at,
		updated_at, deleted_at
	FROM
		todos
	WHERE
		id = ?
		AND
		deleted_at IS NULL`

	err = tx.QueryRowContext(ctx, selectStmt, id).Scan(
		&out.ID, &out.Title, &out.ActivityGroupID, &out.IsActive, &out.Priority,
		&out.CreatedAt, &out.UpdatedAt, &out.DeletedAt,
	)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.UpdateByIdOut{}, err2
		}

		return
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.UpdateByIdOut{}, err2
		}

		return
	}

	return
}
