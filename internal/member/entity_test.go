package member

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMember_IsVerified(t *testing.T) {
	id, _ := uuid.NewRandom()
	now := time.Now().UTC()

	type fields struct {
		ID         uuid.UUID
		Name       string
		VerifiedAt *time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "+Member is verified",
			fields: fields{
				ID:         id,
				Name:       "John Doe",
				VerifiedAt: &now,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			want: true,
		},
		{
			name: "+Member is not verified",
			fields: fields{
				ID:         id,
				Name:       "John Doe",
				VerifiedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Member{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				VerifiedAt: tt.fields.VerifiedAt,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if got := m.IsVerified(); got != tt.want {
				t.Errorf("IsVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMember_MakeVerified(t *testing.T) {
	id, _ := uuid.NewRandom()
	now := time.Now().UTC()

	type fields struct {
		ID         uuid.UUID
		Name       string
		VerifiedAt *time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "+Member is unverified",
			fields: fields{
				ID:         id,
				Name:       "John Doe",
				VerifiedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: false,
		},
		{
			name: "-Member already verified",
			fields: fields{
				ID:         id,
				Name:       "John Doe",
				VerifiedAt: &now,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Member{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				VerifiedAt: tt.fields.VerifiedAt,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := m.MakeVerified(); (err != nil) != tt.wantErr {
				t.Errorf("MakeVerified() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMember_MakeUnverified(t *testing.T) {
	now := time.Now().UTC()
	type fields struct {
		ID         uuid.UUID
		Name       string
		VerifiedAt *time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name         string
		fields       fields
		wantVerified bool
	}{
		{
			name: "+Member is verified to unverified",
			fields: fields{
				Name:       "John Doe",
				VerifiedAt: &now,
			},
			wantVerified: false,
		},
		{
			name: "+Member is unverified to unverified",
			fields: fields{
				Name:       "John Doe",
				VerifiedAt: nil,
			},
			wantVerified: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Member{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				VerifiedAt: tt.fields.VerifiedAt,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			m.MakeUnverified()
			if tt.wantVerified != m.IsVerified() {
				t.Errorf("MakeUnverified() wantVerified %v", tt.wantVerified)
			}
		})
	}
}

func TestMember_Validate(t *testing.T) {
	id, _ := uuid.NewRandom()
	now := time.Now().UTC()

	type fields struct {
		ID         uuid.UUID
		Name       string
		VerifiedAt *time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "+Name is valid",
			fields: fields{
				ID:         id,
				Name:       "John Doe",
				VerifiedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: false,
		},
		{
			name: "-Name is empty",
			fields: fields{
				ID:         id,
				Name:       "",
				VerifiedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: true,
		},
		{
			name: "-Name too short",
			fields: fields{
				ID:         id,
				Name:       "my",
				VerifiedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: true,
		},
		{
			name: "-Name too long",
			fields: fields{
				ID:         id,
				Name:       "111111111111111111111111111111111111111111111111111",
				VerifiedAt: nil,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Member{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				VerifiedAt: tt.fields.VerifiedAt,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMember(t *testing.T) {
	now := time.Now().UTC()

	type args struct {
		name     string
		verified bool
	}
	tests := []struct {
		name    string
		args    args
		want    Member
		wantErr bool
	}{
		{
			name: "+Member is valid",
			args: args{
				name:     "John Doe",
				verified: false,
			},
			want: Member{
				Name:       "John Doe",
				VerifiedAt: nil,
			},
			wantErr: false,
		},
		{
			name: "+Member is valid (verified)",
			args: args{
				name:     "John Doe",
				verified: true,
			},
			want: Member{
				Name:       "John Doe",
				VerifiedAt: &now,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMember(tt.args.name, tt.args.verified)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMember() error = %v,\n wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.verified && tt.want.VerifiedAt == nil {
				tt.want.VerifiedAt = got.VerifiedAt
			}
			tt.want.ID = got.ID
			tt.want.CreatedAt = got.CreatedAt
			tt.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMember() got = %v,\n want %v", got, tt.want)
			}
		})
	}
}
