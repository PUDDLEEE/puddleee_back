// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"

	ent "github.com/PUDDLEEE/puddleee_back/ent"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IAuthRepository is an autogenerated mock type for the IAuthRepository type
type IAuthRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1, _a2
func (_m *IAuthRepository) Create(_a0 context.Context, _a1 *ent.Client, _a2 authdto.CreateVerificationDTO) (*ent.Verification, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *ent.Verification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, authdto.CreateVerificationDTO) (*ent.Verification, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, authdto.CreateVerificationDTO) *ent.Verification); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Verification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Client, authdto.CreateVerificationDTO) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1, _a2
func (_m *IAuthRepository) Delete(_a0 context.Context, _a1 *ent.Client, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOneByUUID provides a mock function with given fields: _a0, _a1, _a2
func (_m *IAuthRepository) FindOneByUUID(_a0 context.Context, _a1 *ent.Client, _a2 string) (*ent.Verification, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for FindOneByUUID")
	}

	var r0 *ent.Verification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, string) (*ent.Verification, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, string) *ent.Verification); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Verification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Client, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateCode provides a mock function with given fields: _a0
func (_m *IAuthRepository) GenerateCode(_a0 int) string {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GenerateCode")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateUUIDFromString provides a mock function with given fields: _a0
func (_m *IAuthRepository) GenerateUUIDFromString(_a0 string) *uuid.UUID {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GenerateUUIDFromString")
	}

	var r0 *uuid.UUID
	if rf, ok := ret.Get(0).(func(string) *uuid.UUID); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uuid.UUID)
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *IAuthRepository) Update(_a0 context.Context, _a1 *ent.Client, _a2 string, _a3 authdto.UpdateVerificationDTO) (*ent.Verification, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *ent.Verification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, string, authdto.UpdateVerificationDTO) (*ent.Verification, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, string, authdto.UpdateVerificationDTO) *ent.Verification); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Verification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Client, string, authdto.UpdateVerificationDTO) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIAuthRepository creates a new instance of IAuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthRepository {
	mock := &IAuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
