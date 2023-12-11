package roomdto

type CreateRoomDTO struct {
	Title       string
	IsCompleted bool
	Questioner  int
	Respondent  []int
	Category    int
}
