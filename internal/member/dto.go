package member

import (
	"time"

	"github.com/google/uuid"
)

type newRequest struct {
	Name     string `json:"name"`
	Verified bool   `json:"verified"`
}

type editRequest struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name"`
	Verified bool      `json:"verified"`
}

type findRequest struct {
	ID uuid.UUID `json:"-"`
}

type response struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Verified   bool       `json:"verified"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func toResponse(member Member) response {
	return response{
		ID:         member.ID,
		Name:       member.Name,
		Verified:   member.IsVerified(),
		VerifiedAt: member.VerifiedAt,
		CreatedAt:  member.CreatedAt,
		UpdatedAt:  member.UpdatedAt,
	}
}

func toListResponse(members []Member) []response {
	responses := make([]response, len(members))
	for i, member := range members {
		responses[i] = toResponse(member)
	}
	return responses
}
