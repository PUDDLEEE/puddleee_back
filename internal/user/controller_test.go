package user

import (
	"context"
	"errors"
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/middlewares"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserController_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserService, *gin.Engine)
		expect func(*testing.T, UserController, *gin.Engine, *httptest.ResponseRecorder, *mocks.IUserService)
	}{
		{
			name: "CreateUser/email_existed",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {
				service.On("CreateUser", dto.CreateUserDTO{}).
					Return(nil, errors.New("email")).
					Once()
				router.Use(func(ctx *gin.Context) {
					ctx.Set("body", dto.CreateUserDTO{})
				})
				router.Use(middlewares.ErrorHandler())
			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.POST("/api/v1/user", controller.CreateUser)
				req, _ := http.NewRequest(http.MethodPost, "/api/v1/user", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusBadRequest)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
		{
			name: "CreateUser/username_existed",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {
				service.On("CreateUser", dto.CreateUserDTO{}).
					Return(nil, errors.New("username")).
					Once()
				router.Use(func(ctx *gin.Context) {
					ctx.Set("body", dto.CreateUserDTO{})
				})
				router.Use(middlewares.ErrorHandler())
			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.POST("/api/v1/user", controller.CreateUser)
				req, _ := http.NewRequest(http.MethodPost, "/api/v1/user", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusBadRequest)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			mockUserService := mocks.NewIUserService(t)
			userController := NewController(mockUserService)

			require.NoError(t, err)

			router := gin.Default()
			w := httptest.NewRecorder()

			tt.before(t, mockUserService, router)
			tt.expect(t, *userController, router, w, mockUserService)
		})
	}
}

func TestUserController_ViewProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserService, *gin.Engine)
		expect func(*testing.T, UserController, *gin.Engine, *httptest.ResponseRecorder, *mocks.IUserService)
	}{
		{
			name: "ViewProfile/invalid_url_param",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {

				router.Use(middlewares.ErrorHandler())
			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.GET("/api/v1/user/:id", controller.ViewProfile)
				req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/abc", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusBadRequest)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
		{
			name: "ViewProfile/user_not_found",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {
				service.On("FindOneUser", 1).
					Return(nil, errors.New("")).
					Once()
				router.Use(middlewares.ErrorHandler())
			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.GET("/api/v1/user/:id", controller.ViewProfile)
				req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/1", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusNotFound)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			mockUserService := mocks.NewIUserService(t)
			userController := NewController(mockUserService)

			require.NoError(t, err)

			router := gin.Default()
			w := httptest.NewRecorder()

			tt.before(t, mockUserService, router)
			tt.expect(t, *userController, router, w, mockUserService)
		})
	}
}

func TestUserController_UpdateProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserService, *gin.Engine)
		expect func(*testing.T, UserController, *gin.Engine, *httptest.ResponseRecorder, *mocks.IUserService)
	}{
		{
			name: "UpdateProfile/invalid_url_param",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {

				router.Use(middlewares.ErrorHandler())

			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.PATCH("/api/v1/user/:id", controller.UpdateProfile)
				req, _ := http.NewRequest(http.MethodPatch, "/api/v1/user/abc", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusBadRequest)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
		{
			name: "UpdateProfile/update_error",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {
				service.On("UpdateUser", 1, dto.UpdateUserDTO{}).
					Return(nil, errors.New("")).
					Once()
				router.Use(func(ctx *gin.Context) {
					ctx.Set("body", dto.UpdateUserDTO{})
				})
				router.Use(middlewares.ErrorHandler())
			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.PATCH("/api/v1/user/:id", controller.UpdateProfile)
				req, _ := http.NewRequest(http.MethodPatch, "/api/v1/user/1", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusInternalServerError)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			mockUserService := mocks.NewIUserService(t)
			userController := NewController(mockUserService)

			require.NoError(t, err)

			router := gin.Default()
			w := httptest.NewRecorder()

			tt.before(t, mockUserService, router)
			tt.expect(t, *userController, router, w, mockUserService)
		})
	}
}

func TestUserController_DeleteProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserService, *gin.Engine)
		expect func(*testing.T, UserController, *gin.Engine, *httptest.ResponseRecorder, *mocks.IUserService)
	}{
		{
			name: "DeleteProfile/invalid_url_param",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {

				router.Use(middlewares.ErrorHandler())

			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.DELETE("/api/v1/user/:id", controller.DeleteProfile)
				req, _ := http.NewRequest(http.MethodDelete, "/api/v1/user/abc", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusBadRequest)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
		{
			name: "DeleteProfile/delete_error",
			before: func(t *testing.T, service *mocks.IUserService, router *gin.Engine) {
				service.On("DeleteUser", 1).
					Return(errors.New("")).
					Once()
				router.Use(middlewares.ErrorHandler())
			},
			expect: func(t *testing.T, controller UserController, router *gin.Engine, w *httptest.ResponseRecorder, service *mocks.IUserService) {
				router.DELETE("/api/v1/user/:id", controller.DeleteProfile)
				req, _ := http.NewRequest(http.MethodDelete, "/api/v1/user/1", nil)
				router.ServeHTTP(w, req)
				require.Equal(t, w.Code, http.StatusInternalServerError)
				require.NotEmpty(t, w.Body)
				service.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			mockUserService := mocks.NewIUserService(t)
			userController := NewController(mockUserService)

			require.NoError(t, err)

			router := gin.Default()
			w := httptest.NewRecorder()

			tt.before(t, mockUserService, router)
			tt.expect(t, *userController, router, w, mockUserService)
		})
	}
}
