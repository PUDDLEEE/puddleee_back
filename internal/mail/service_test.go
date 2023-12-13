package mail

import (
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"testing"
	"time"

	smtpmock "github.com/mocktools/go-smtp-mock/v2"
	"github.com/stretchr/testify/require"
)

func TestMailService_Send(t *testing.T) {
	tests := []struct {
		name   string
		expect func(*testing.T, *MailService)
	}{
		{
			name: "Verify valid access token",
			expect: func(t *testing.T, service *MailService) {
				err := service.Send("123@naver.com", "hello")
				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := smtpmock.New(smtpmock.ConfigurationAttr{
				LogToStdout:       true,
				LogServerActivity: true,
			})
			if err := server.Start(); err != nil {
				t.Fatal(err)
			}
			port := strconv.Itoa(server.PortNumber())
			params := MailServiceParams{}
			params.server = "127.0.0.1"
			params.port = port
			params.from = "123@naver.com"
			params.password = "123"
			service := NewMailService(params)

			address := fmt.Sprintf("%s:%s", params.server, port)
			timeout := time.Duration(2) * time.Second

			connection, _ := net.DialTimeout("tcp", address, timeout)
			_, _ = smtp.NewClient(connection, params.server)

			server.Messages()
			tt.expect(t, service)
			if err := server.Stop(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
