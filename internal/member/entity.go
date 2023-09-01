package member

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrMemberNotFound        = errors.New("member not found")
	ErrMemberValidation      = errors.New("member validation error")
	ErrMemberIDFailure       = errors.New("member failed to generate unique id")
	ErrMemberAlreadyVerified = errors.New("member has been verified")
)

type Member struct {
	ID         uuid.UUID
	Name       string
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (m *Member) BeforeSave() error {
	// On new.
	if m.ID == uuid.Nil {
		uuid4, err := uuid.NewRandom()
		if err != nil {
			return ErrMemberIDFailure
		}
		m.ID = uuid4
		m.CreatedAt = time.Now().UTC()
		m.UpdatedAt = time.Now().UTC()
		return nil
	}

	// On Update.
	m.UpdatedAt = time.Now().UTC()

	return m.Validate()
}

func (m *Member) IsVerified() bool {
	return m.VerifiedAt != nil
}

func (m *Member) MakeVerified() error {
	if m.IsVerified() {
		return ErrMemberAlreadyVerified
	}

	now := time.Now().UTC()
	m.VerifiedAt = &now
	return nil
}

func (m *Member) MakeUnverified() {
	m.VerifiedAt = nil
}

func (m *Member) Validate() error {
	err := ValidateName(m.Name)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrMemberValidation, err)
	}
	return nil
}

func NewMember(name string, verified bool) (Member, error) {
	member := Member{
		Name:       name,
		VerifiedAt: nil,
	}

	if verified {
		_ = member.MakeVerified()
	}
	return member, nil
}
