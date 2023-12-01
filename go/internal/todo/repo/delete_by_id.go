package repo

import (
	"context"
	"database/sql"
	"time"
)

func (r *Repo) DeleteById(ctx context.Context, id int64) (affected int64, err error) {
	if id == 0 {
		return 0, sql.ErrNoRows
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return
	}

	var stmt = `UPDATE todos SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL`

	commandTag, err := tx.ExecContext(ctx, stmt, time.Now().Unix(), id)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return 0, err2
		}

		return
	}

	affected, err = commandTag.RowsAffected()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return 0, err2
		}

		return
	}
	if affected == 0 {
		err2 := tx.Rollback()
		if err2 != nil {
			return 0, err2
		}

		return
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return 0, err2
		}

		return
	}

	return
}
