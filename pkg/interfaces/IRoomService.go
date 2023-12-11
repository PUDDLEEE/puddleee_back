package interfaces

import (
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/room/dto"
)

type IRoomService interface {
	CreateRoom(roomdto.CreateRoomDTO) (*ent.Room, error)
	FindOneRoom(int) (*ent.Room, error)
	FindRoom() ([]*ent.Room, error)
	UpdateRoom(int, roomdto.UpdateRoomDTO) (*ent.Room, error)
	DeleteRoom(int) error
}
