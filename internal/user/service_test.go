package user

import (
	"context"
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.MockIUserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserService, *ent.Client, context.Context)
	}{
		{
			name: "Create/username_already_exists",
			before: func(t *testing.T, repo *mocks.MockIUserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, service *UserService, client *ent.Client, ctx context.Context) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			//mockCtrl := gomock.NewController(t)
			//mockUserRepo := mocks.NewMockIUserRepository(mockCtrl)
			//
			//userService := NewService(mockUserRepo, context.Background(), client)
			//tt.before(t, mockUserRepo, client, context.Background())
			//tt.expect(t, &userService, client, context.Background())
		})
	}
}
