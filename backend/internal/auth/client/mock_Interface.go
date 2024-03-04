// Code generated by mockery. DO NOT EDIT.

package auth_client

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// GetAuthorizationEndpoint provides a mock function with given fields: ctx
func (_m *MockInterface) GetAuthorizationEndpoint(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizationEndpoint")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetAuthorizationEndpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizationEndpoint'
type MockInterface_GetAuthorizationEndpoint_Call struct {
	*mock.Call
}

// GetAuthorizationEndpoint is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInterface_Expecter) GetAuthorizationEndpoint(ctx interface{}) *MockInterface_GetAuthorizationEndpoint_Call {
	return &MockInterface_GetAuthorizationEndpoint_Call{Call: _e.mock.On("GetAuthorizationEndpoint", ctx)}
}

func (_c *MockInterface_GetAuthorizationEndpoint_Call) Run(run func(ctx context.Context)) *MockInterface_GetAuthorizationEndpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInterface_GetAuthorizationEndpoint_Call) Return(_a0 string, _a1 error) *MockInterface_GetAuthorizationEndpoint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetAuthorizationEndpoint_Call) RunAndReturn(run func(context.Context) (string, error)) *MockInterface_GetAuthorizationEndpoint_Call {
	_c.Call.Return(run)
	return _c
}

// GetRefreshedAccessToken provides a mock function with given fields: ctx, clientId, refreshToken
func (_m *MockInterface) GetRefreshedAccessToken(ctx context.Context, clientId string, refreshToken string) (*AuthTokenResponse, error) {
	ret := _m.Called(ctx, clientId, refreshToken)

	if len(ret) == 0 {
		panic("no return value specified for GetRefreshedAccessToken")
	}

	var r0 *AuthTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*AuthTokenResponse, error)); ok {
		return rf(ctx, clientId, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *AuthTokenResponse); ok {
		r0 = rf(ctx, clientId, refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AuthTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clientId, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetRefreshedAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRefreshedAccessToken'
type MockInterface_GetRefreshedAccessToken_Call struct {
	*mock.Call
}

// GetRefreshedAccessToken is a helper method to define mock.On call
//   - ctx context.Context
//   - clientId string
//   - refreshToken string
func (_e *MockInterface_Expecter) GetRefreshedAccessToken(ctx interface{}, clientId interface{}, refreshToken interface{}) *MockInterface_GetRefreshedAccessToken_Call {
	return &MockInterface_GetRefreshedAccessToken_Call{Call: _e.mock.On("GetRefreshedAccessToken", ctx, clientId, refreshToken)}
}

func (_c *MockInterface_GetRefreshedAccessToken_Call) Run(run func(ctx context.Context, clientId string, refreshToken string)) *MockInterface_GetRefreshedAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockInterface_GetRefreshedAccessToken_Call) Return(_a0 *AuthTokenResponse, _a1 error) *MockInterface_GetRefreshedAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetRefreshedAccessToken_Call) RunAndReturn(run func(context.Context, string, string) (*AuthTokenResponse, error)) *MockInterface_GetRefreshedAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokenEndpoint provides a mock function with given fields: ctx
func (_m *MockInterface) GetTokenEndpoint(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenEndpoint")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetTokenEndpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenEndpoint'
type MockInterface_GetTokenEndpoint_Call struct {
	*mock.Call
}

// GetTokenEndpoint is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInterface_Expecter) GetTokenEndpoint(ctx interface{}) *MockInterface_GetTokenEndpoint_Call {
	return &MockInterface_GetTokenEndpoint_Call{Call: _e.mock.On("GetTokenEndpoint", ctx)}
}

func (_c *MockInterface_GetTokenEndpoint_Call) Run(run func(ctx context.Context)) *MockInterface_GetTokenEndpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInterface_GetTokenEndpoint_Call) Return(_a0 string, _a1 error) *MockInterface_GetTokenEndpoint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetTokenEndpoint_Call) RunAndReturn(run func(context.Context) (string, error)) *MockInterface_GetTokenEndpoint_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokenResponse provides a mock function with given fields: ctx, clientId, code, redirecturi
func (_m *MockInterface) GetTokenResponse(ctx context.Context, clientId string, code string, redirecturi string) (*AuthTokenResponse, error) {
	ret := _m.Called(ctx, clientId, code, redirecturi)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenResponse")
	}

	var r0 *AuthTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*AuthTokenResponse, error)); ok {
		return rf(ctx, clientId, code, redirecturi)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *AuthTokenResponse); ok {
		r0 = rf(ctx, clientId, code, redirecturi)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AuthTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, clientId, code, redirecturi)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetTokenResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenResponse'
type MockInterface_GetTokenResponse_Call struct {
	*mock.Call
}

// GetTokenResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - clientId string
//   - code string
//   - redirecturi string
func (_e *MockInterface_Expecter) GetTokenResponse(ctx interface{}, clientId interface{}, code interface{}, redirecturi interface{}) *MockInterface_GetTokenResponse_Call {
	return &MockInterface_GetTokenResponse_Call{Call: _e.mock.On("GetTokenResponse", ctx, clientId, code, redirecturi)}
}

func (_c *MockInterface_GetTokenResponse_Call) Run(run func(ctx context.Context, clientId string, code string, redirecturi string)) *MockInterface_GetTokenResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockInterface_GetTokenResponse_Call) Return(_a0 *AuthTokenResponse, _a1 error) *MockInterface_GetTokenResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetTokenResponse_Call) RunAndReturn(run func(context.Context, string, string, string) (*AuthTokenResponse, error)) *MockInterface_GetTokenResponse_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserInfo provides a mock function with given fields: ctx, accessToken
func (_m *MockInterface) GetUserInfo(ctx context.Context, accessToken string) (*UserInfo, error) {
	ret := _m.Called(ctx, accessToken)

	if len(ret) == 0 {
		panic("no return value specified for GetUserInfo")
	}

	var r0 *UserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*UserInfo, error)); ok {
		return rf(ctx, accessToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *UserInfo); ok {
		r0 = rf(ctx, accessToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetUserInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserInfo'
type MockInterface_GetUserInfo_Call struct {
	*mock.Call
}

// GetUserInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - accessToken string
func (_e *MockInterface_Expecter) GetUserInfo(ctx interface{}, accessToken interface{}) *MockInterface_GetUserInfo_Call {
	return &MockInterface_GetUserInfo_Call{Call: _e.mock.On("GetUserInfo", ctx, accessToken)}
}

func (_c *MockInterface_GetUserInfo_Call) Run(run func(ctx context.Context, accessToken string)) *MockInterface_GetUserInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockInterface_GetUserInfo_Call) Return(_a0 *UserInfo, _a1 error) *MockInterface_GetUserInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetUserInfo_Call) RunAndReturn(run func(context.Context, string) (*UserInfo, error)) *MockInterface_GetUserInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}