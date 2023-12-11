package room

import (
	"context"
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/room/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
)

type RoomService struct {
	roomRepository interfaces.IRoomRepository
	ctx            context.Context
	client         *ent.Client
}

func (r *RoomService) CreateRoom(dto roomdto.CreateRoomDTO) (*ent.Room, error) {
	return nil, nil
}

func (r *RoomService) FindOneRoom(id int) (*ent.Room, error) {
	return nil, nil
}

func (r *RoomService) FindRoom() ([]*ent.Room, error) {
	return nil, nil
}

func (r *RoomService) UpdateRoom(id int, dto roomdto.UpdateRoomDTO) (*ent.Room, error) {
	return nil, nil
}

func (r *RoomService) DeleteRoom(id int) error {
	return nil
}

func NewRoomService(repo interfaces.IRoomRepository, ctx context.Context, client *ent.Client) *RoomService {
	if roomRepository, ok := repo.(*RoomRepository); ok {
		return &RoomService{
			roomRepository: roomRepository,
			ctx:            ctx,
			client:         client,
		}
	}

	if roomRepository, ok := repo.(*mocks.IRoomRepository); ok {
		return &RoomService{
			roomRepository: roomRepository,
			ctx:            ctx,
			client:         client,
		}
	}
	return nil
}
