// Code generated by mockery v2.38.0. DO NOT EDIT.

package oidcmocks

import (
	oidc "github.com/kyma-project/test-infra/pkg/oidc"
	mock "github.com/stretchr/testify/mock"
)

// MockVerifierConfigOption is an autogenerated mock type for the VerifierConfigOption type
type MockVerifierConfigOption struct {
	mock.Mock
}

type MockVerifierConfigOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVerifierConfigOption) EXPECT() *MockVerifierConfigOption_Expecter {
	return &MockVerifierConfigOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockVerifierConfigOption) Execute(_a0 *oidc.VerifierConfig) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*oidc.VerifierConfig) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVerifierConfigOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockVerifierConfigOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *oidc.VerifierConfig
func (_e *MockVerifierConfigOption_Expecter) Execute(_a0 interface{}) *MockVerifierConfigOption_Execute_Call {
	return &MockVerifierConfigOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockVerifierConfigOption_Execute_Call) Run(run func(_a0 *oidc.VerifierConfig)) *MockVerifierConfigOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*oidc.VerifierConfig))
	})
	return _c
}

func (_c *MockVerifierConfigOption_Execute_Call) Return(_a0 error) *MockVerifierConfigOption_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVerifierConfigOption_Execute_Call) RunAndReturn(run func(*oidc.VerifierConfig) error) *MockVerifierConfigOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVerifierConfigOption creates a new instance of MockVerifierConfigOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVerifierConfigOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVerifierConfigOption {
	mock := &MockVerifierConfigOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
