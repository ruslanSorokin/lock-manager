// Code generated by mockery. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// ServiceMetric is an autogenerated mock type for the ServiceMetricI type
type ServiceMetric struct {
	mock.Mock
}

type ServiceMetric_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceMetric) EXPECT() *ServiceMetric_Expecter {
	return &ServiceMetric_Expecter{mock: &_m.Mock}
}

// IncLockedTotal provides a mock function with given fields:
func (_m *ServiceMetric) IncLockedTotal() {
	_m.Called()
}

// ServiceMetric_IncLockedTotal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncLockedTotal'
type ServiceMetric_IncLockedTotal_Call struct {
	*mock.Call
}

// IncLockedTotal is a helper method to define mock.On call
func (_e *ServiceMetric_Expecter) IncLockedTotal() *ServiceMetric_IncLockedTotal_Call {
	return &ServiceMetric_IncLockedTotal_Call{Call: _e.mock.On("IncLockedTotal")}
}

func (_c *ServiceMetric_IncLockedTotal_Call) Run(run func()) *ServiceMetric_IncLockedTotal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceMetric_IncLockedTotal_Call) Return() *ServiceMetric_IncLockedTotal_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServiceMetric_IncLockedTotal_Call) RunAndReturn(run func()) *ServiceMetric_IncLockedTotal_Call {
	_c.Call.Return(run)
	return _c
}

// IncUnlockedTotal provides a mock function with given fields:
func (_m *ServiceMetric) IncUnlockedTotal() {
	_m.Called()
}

// ServiceMetric_IncUnlockedTotal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncUnlockedTotal'
type ServiceMetric_IncUnlockedTotal_Call struct {
	*mock.Call
}

// IncUnlockedTotal is a helper method to define mock.On call
func (_e *ServiceMetric_Expecter) IncUnlockedTotal() *ServiceMetric_IncUnlockedTotal_Call {
	return &ServiceMetric_IncUnlockedTotal_Call{Call: _e.mock.On("IncUnlockedTotal")}
}

func (_c *ServiceMetric_IncUnlockedTotal_Call) Run(run func()) *ServiceMetric_IncUnlockedTotal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceMetric_IncUnlockedTotal_Call) Return() *ServiceMetric_IncUnlockedTotal_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServiceMetric_IncUnlockedTotal_Call) RunAndReturn(run func()) *ServiceMetric_IncUnlockedTotal_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceMetric creates a new instance of ServiceMetric. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceMetric(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceMetric {
	mock := &ServiceMetric{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}