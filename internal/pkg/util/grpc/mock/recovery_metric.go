// Code generated by mockery. DO NOT EDIT.

package grpcutilmock

import mock "github.com/stretchr/testify/mock"

// RecoveryMetric is an autogenerated mock type for the RecoveryMetricI type
type RecoveryMetric struct {
	mock.Mock
}

type RecoveryMetric_Expecter struct {
	mock *mock.Mock
}

func (_m *RecoveryMetric) EXPECT() *RecoveryMetric_Expecter {
	return &RecoveryMetric_Expecter{mock: &_m.Mock}
}

// Inc provides a mock function with given fields:
func (_m *RecoveryMetric) Inc() {
	_m.Called()
}

// RecoveryMetric_Inc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Inc'
type RecoveryMetric_Inc_Call struct {
	*mock.Call
}

// Inc is a helper method to define mock.On call
func (_e *RecoveryMetric_Expecter) Inc() *RecoveryMetric_Inc_Call {
	return &RecoveryMetric_Inc_Call{Call: _e.mock.On("Inc")}
}

func (_c *RecoveryMetric_Inc_Call) Run(run func()) *RecoveryMetric_Inc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RecoveryMetric_Inc_Call) Return() *RecoveryMetric_Inc_Call {
	_c.Call.Return()
	return _c
}

func (_c *RecoveryMetric_Inc_Call) RunAndReturn(run func()) *RecoveryMetric_Inc_Call {
	_c.Call.Return(run)
	return _c
}

// NewRecoveryMetric creates a new instance of RecoveryMetric. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRecoveryMetric(t interface {
	mock.TestingT
	Cleanup(func())
}) *RecoveryMetric {
	mock := &RecoveryMetric{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
