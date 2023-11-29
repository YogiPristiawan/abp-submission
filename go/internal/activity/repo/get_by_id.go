package repo

import (
	"context"
	"database/sql"
	"todo/internal/activity/model"
)

func (r *Repo) GetById(ctx context.Context, activityId int64) (out model.GetByIdOut, err error) {
	if activityId == 0 {
		return model.GetByIdOut{}, sql.ErrNoRows
	}

	var stmt = `SELECT id, title, email, created_at, updated_at, deleted_at FROM activities WHERE id = ? AND deleted_at IS NULL`

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return
	}

	err = tx.QueryRowContext(ctx, stmt, activityId).Scan(
		&out.ID, &out.Title, &out.Email, &out.CreatedAt, &out.UpdatedAt, &out.DeletedAt,
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
