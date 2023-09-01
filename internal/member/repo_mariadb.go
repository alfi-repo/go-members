package member

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

func NewRepoMariaDB(db *sql.DB) RepoMariaDB {
	return RepoMariaDB{
		db: db,
	}
}

type RepoMariaDB struct {
	db *sql.DB
}

func (r RepoMariaDB) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	q := `DELETE FROM members WHERE id = ?`
	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	if affectedRows, _ := res.RowsAffected(); affectedRows == 0 {
		return ErrMemberNotFound
	}
	return nil
}

func (r RepoMariaDB) Save(member *Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := member.BeforeSave(); err != nil {
		return err
	}

	// In the context of mysql/mariadb; we can use an upsert because there is no
	// unique column other than ID which is PK.
	q := `INSERT INTO members(id, name, verified_at, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE name=?, verified_at=?, updated_at=?`
	if _, err := r.db.ExecContext(
		ctx,
		q,
		member.ID,
		member.Name,
		member.VerifiedAt,
		member.CreatedAt,
		member.UpdatedAt,
		member.Name,
		member.VerifiedAt,
		member.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}

func (r RepoMariaDB) FindByID(id uuid.UUID) (Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var member Member
	q := `SELECT id, name, verified_at, created_at, updated_at FROM members WHERE id = ?`
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&member.ID,
		&member.Name,
		&member.VerifiedAt,
		&member.CreatedAt,
		&member.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Member{}, ErrMemberNotFound
		default:
			return Member{}, err
		}
	}

	return member, nil
}

func (r RepoMariaDB) List() ([]Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var members []Member
	q := `SELECT id, name, verified_at, created_at, updated_at FROM members`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return members, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return members, err
		}
		var member Member
		if err = rows.Scan(
			&member.ID,
			&member.Name,
			&member.VerifiedAt,
			&member.CreatedAt,
			&member.UpdatedAt,
		); err != nil {
			return members, err
		}
		members = append(members, member)
	}
	return members, nil
}
