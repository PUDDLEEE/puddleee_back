// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	jwtAuth "github.com/PUDDLEEE/puddleee_back/internal/jwtAuth"
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"
)

// IJwtAuthService is an autogenerated mock type for the IJwtAuthService type
type IJwtAuthService struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: _a0
func (_m *IJwtAuthService) CreateAccessToken(_a0 jwtAuth.AuthTokenClaims) (*string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccessToken")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(jwtAuth.AuthTokenClaims) (*string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(jwtAuth.AuthTokenClaims) *string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(jwtAuth.AuthTokenClaims) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: _a0
func (_m *IJwtAuthService) CreateRefreshToken(_a0 jwt.RegisteredClaims) (*string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(jwt.RegisteredClaims) (*string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(jwt.RegisteredClaims) *string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(jwt.RegisteredClaims) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseAccessToken provides a mock function with given fields: _a0
func (_m *IJwtAuthService) ParseAccessToken(_a0 string) (*jwt.Token, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ParseAccessToken")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRefreshToken provides a mock function with given fields: _a0
func (_m *IJwtAuthService) ParseRefreshToken(_a0 string) (*jwt.Token, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ParseRefreshToken")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReassignAccessToken provides a mock function with given fields: _a0, _a1
func (_m *IJwtAuthService) ReassignAccessToken(_a0 string, _a1 string) (*string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ReassignAccessToken")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) *string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyAccessToken provides a mock function with given fields: _a0
func (_m *IJwtAuthService) VerifyAccessToken(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for VerifyAccessToken")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyRefreshToken provides a mock function with given fields: _a0
func (_m *IJwtAuthService) VerifyRefreshToken(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for VerifyRefreshToken")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIJwtAuthService creates a new instance of IJwtAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIJwtAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IJwtAuthService {
	mock := &IJwtAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
