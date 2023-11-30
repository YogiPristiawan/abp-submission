package repo

import (
	"context"
	"database/sql"
	"time"
	"todo/internal/todo/model"
)

func (r *Repo) Create(ctx context.Context, in model.CreateIn) (out model.CreateOut, err error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return
	}

	var activityExistsStmt = `SELECT EXISTS(SELECT * FROM activities WHERE id = ?)`
	var activityExists bool
	err = tx.QueryRowContext(ctx, activityExistsStmt, in.ActivityGroupID).Scan(&activityExists)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.CreateOut{}, err2
		}

		return
	}
	if !activityExists {
		return model.CreateOut{}, sql.ErrNoRows
	}

	var insertStmt = `
	INSERT INTO todos
	(
		activity_group_id, title, created_at, updated_at
	)
	VALUES
	(
		?, ?, ?, ?
	)`

	now := time.Now()
	nowInUnix := now.Unix()
	_, err = tx.ExecContext(ctx, insertStmt,
		in.ActivityGroupID, in.Title, nowInUnix, nowInUnix,
	)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.CreateOut{}, err2
		}

		return
	}

	// get the iserted row
	var insertedStmt = `
	SELECT
		id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at
	FROM
		todos
	ORDER BY id DESC LIMIT 1`

	err = tx.QueryRowContext(ctx, insertedStmt).Scan(
		&out.ID, &out.ActivityGroupID, &out.Title, &out.IsActive, &out.Priority,
		&out.CreatedAt, &out.UpdatedAt, &out.DeletedAt,
	)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.CreateOut{}, err2
		}

		return
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.CreateOut{}, err2
		}

		return
	}

	return
}
