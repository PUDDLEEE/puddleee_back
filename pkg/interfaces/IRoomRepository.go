package interfaces

import (
	"context"
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/room/dto"
)

type IRoomRepository interface {
	Create(context.Context, *ent.Client, roomdto.CreateRoomDTO) (*ent.Room, error)
	Find(context.Context, *ent.Client) ([]*ent.Room, error)
	FindOneById(context.Context, *ent.Client, int) (*ent.Room, error)
	Update(context.Context, *ent.Client, int, roomdto.UpdateRoomDTO) (*ent.Room, error)
	Delete(context.Context, *ent.Client, int) error
}
