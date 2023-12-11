package room

import (
	"context"
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/room/dto"
)

type RoomRepository struct{}

func (r *RoomRepository) Create(ctx context.Context, client *ent.Client, dto roomdto.CreateRoomDTO) (*ent.Room, error) {
	return nil, nil
}

func (r *RoomRepository) Find(ctx context.Context, client *ent.Client) ([]*ent.Room, error) {
	return nil, nil
}

func (r *RoomRepository) Update(ctx context.Context, client *ent.Client, id int, dto roomdto.UpdateRoomDTO) (*ent.Room, error) {
	return nil, nil
}

func (r *RoomRepository) FindOneById(ctx context.Context, client *ent.Client, id int) (*ent.Room, error) {
	return nil, nil
}
func (r *RoomRepository) Delete(ctx context.Context, client *ent.Client, id int) error {
	return nil
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}
