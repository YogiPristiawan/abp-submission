package repo

import (
	"context"
	"database/sql"
	"time"
	"todo/internal/activity/model"
)

func (r *Repo) UpdateById(ctx context.Context, id int64, in model.UpdateByIdIn) (out model.UpdateByIdOut, err error) {
	// reduce I/O
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

	// var countStmt = `SELECT COUNT(id) FROM activities WHERE id = ? AND deleted_at IS NULL`
	// var totalRow int64
	// err = tx.QueryRowContext(ctx, countStmt, id).Scan(&totalRow)
	// if err != nil {
	// 	err2 := tx.Rollback()
	// 	if err2 != nil {
	// 		return 0, err2
	// 	}

	// 	return
	// }
	// if totalRow == 0 {
	// 	return 0, sql.ErrNoRows
	// }

	var updateStmt = `
	UPDATE
		activities
	SET
		title = ?, updated_at = ?
	WHERE
		id = ?
		AND
		deleted_at IS NULL`

	now := time.Now().Unix()
	commandTag, err := tx.ExecContext(ctx, updateStmt, in.Title, now, id)
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

	var findUpdatedStmt = `SELECT id, title, email, created_at, updated_at, deleted_at FROM activities WHERE id = ?`
	err = tx.QueryRowContext(ctx, findUpdatedStmt, id).Scan(
		&out.ID, &out.Title, &out.Email, &out.CreatedAt, &out.UpdatedAt, &out.DeletedAt,
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
