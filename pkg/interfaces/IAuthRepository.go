package interfaces

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"
	"github.com/google/uuid"
)

type IAuthRepository interface {
	Create(context.Context, *ent.Client, authdto.CreateVerificationDTO) (*ent.Verification, error)
	Update(context.Context, *ent.Client, string, authdto.UpdateVerificationDTO) (*ent.Verification, error)
	FindOneByUUID(context.Context, *ent.Client, string) (*ent.Verification, error)
	Delete(context.Context, *ent.Client, string) error
	GenerateUUIDFromString(string) *uuid.UUID
	GenerateCode(int) string
}
