package repo

import (
	"context"
	"database/sql"
	"time"
	"todo/internal/activity/model"
)

func (r *Repo) Create(ctx context.Context, in model.CreateIn) (out model.CreateOut, err error) {
	var stmt = `INSERT INTO activities
	(
		title, email, created_at, updated_at
	)
	VALUES
	(
		?, ?, ?, ?
	)`

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return
	}

	now := time.Now().UTC()
	nowInFormat := now.Format("2006-01-02 15:04:05")
	_, err = tx.ExecContext(ctx, stmt, in.Title, in.Email, nowInFormat, nowInFormat)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.CreateOut{}, err2
		}

		return
	}

	// get the last inserted id
	var insertedIdStmt = `SELECT MAX(id) AS inserted_id FROM activities`
	err = tx.QueryRowContext(ctx, insertedIdStmt).Scan(&out.ID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return model.CreateOut{}, err2
		}

		return
	}

	out.Title = in.Title
	out.Email = in.Email
	out.CreatedAt = now
	out.UpdatedAt = now

	return
}
