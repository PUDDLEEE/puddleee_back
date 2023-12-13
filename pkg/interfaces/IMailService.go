package interfaces

type IMailService interface {
	Send(string, string) error
}
