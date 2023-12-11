package room

import (
	"context"
	"github.com/PUDDLEEE/puddleee_back/ent"
)

func InitializeController(ctx context.Context, client *ent.Client) *RoomController {
	roomRepository := NewRoomRepository()
	roomService := NewRoomService(roomRepository, ctx, client)
	RoomController := NewRoomController(roomService)

	return RoomController
}
