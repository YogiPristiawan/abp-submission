package repo

import (
	"context"
	"database/sql"
	"todo/internal/activity/model"
)

func (r *Repo) FindAll(ctx context.Context) (out []model.FindAllOut, err error) {
	var stmt = `
	SELECT
		id, title, email, created_at, updated_at, deleted_at
	FROM
		activities
	WHERE
		deleted_at IS NULL`

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return
	}

	rows, err := tx.QueryContext(ctx, stmt)
	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return nil, err2
		}

		return
	}
	defer func() {
		err = rows.Close()
		return
	}()

	for rows.Next() {
		var row model.FindAllOut
		err = rows.Scan(&row.ID, &row.Title, &row.Email, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt)
		if err != nil {
			return
		}

		out = append(out, row)
	}

	return
}
