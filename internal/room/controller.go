package room

import (
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	roomdto "github.com/PUDDLEEE/puddleee_back/internal/room/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomService interfaces.IRoomService
}

// ViewRooms godoc
//
//	@Summary	View Rooms
//	@Schemes
//	@Description	View Rooms List
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]ent.Room
//	@Failure		500	{object}	errors.CustomError
//	@Router			/rooms [get]
func (r *RoomController) ViewRooms(ctx *gin.Context) {
	_ = roomdto.CreateRoomDTO{}
	_ = errors.INTERNAL_ERROR
}

// ViewOneRoom godoc
//
//	@Summary	View Specific Room's Detail
//	@Schemes
//	@Description	View Specific Room's Detail
//	@Tags			rooms
//	@Param			id	path	int	true	"Room ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ent.Room
//	@Failure		400	{object}	errors.CustomError
//	@Failure		404	{object}	errors.CustomError
//	@Failure		500	{object}	errors.CustomError
//	@Router			/rooms/:id [get]
func (r *RoomController) ViewOneRoom(ctx *gin.Context) {

}

// CreateRoom godoc
//
//	@Summary	Create Room
//	@Schemes
//	@Description	Create Room
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ent.Room
//	@Failure		400	{object}	errors.CustomError
//	@Failure		404	{object}	errors.CustomError
//	@Failure		500	{object}	errors.CustomError
//	@Router			/rooms [post]
func (r *RoomController) CreateRoom(ctx *gin.Context) {

}

// UpdateRoomInfo godoc
//
//	@Summary	Update Room's Detail
//	@Schemes
//	@Description	Update Room's Detail
//	@Param			id				path	int						true	"Room ID"
//	@Param			UpdateRoomDTO	body	roomdto.UpdateRoomDTO	true	"Update Room Info"
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ent.Room
//	@Failure		400	{object}	errors.CustomError
//	@Failure		404	{object}	errors.CustomError
//	@Failure		500	{object}	errors.CustomError
//	@Router			/rooms [patch]
func (r *RoomController) UpdateRoomInfo(ctx *gin.Context) {}

// DeleteRoom godoc
//
//	@Summary	Delete One Room
//	@Schemes
//	@Description	Update Room's Detail
//	@Tags			rooms
//	@Param			id	path	int	true	"Room ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ent.Room
//	@Failure		400	{object}	errors.CustomError
//	@Failure		404	{object}	errors.CustomError
//	@Failure		500	{object}	errors.CustomError
//	@Router			/rooms [delete]
func (r *RoomController) DeleteRoom(ctx *gin.Context) {
}

func NewRoomController(service interfaces.IRoomService) *RoomController {
	if roomService, ok := service.(*RoomService); ok {
		return &RoomController{roomService: roomService}
	}

	if mockRoomService, ok := service.(*mocks.IRoomService); ok {
		return &RoomController{roomService: mockRoomService}
	}
	return nil
}
