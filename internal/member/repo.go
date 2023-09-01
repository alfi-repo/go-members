package member

import (
	"github.com/google/uuid"
)

type Repository interface {
	RepositoryRead
	RepositoryWrite
}

type RepositoryRead interface {
	FindByID(id uuid.UUID) (Member, error)
	List() ([]Member, error)
}

type RepositoryWrite interface {
	Delete(id uuid.UUID) error
	Save(member *Member) error
}
