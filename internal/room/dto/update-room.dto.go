package roomdto

type UpdateRoomDTO struct {
	Title       *string
	IsCompleted *bool
	Questioner  *int
	Respondent  *[]int
	Category    *int
}
