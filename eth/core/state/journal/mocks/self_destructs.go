// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	journal "github.com/berachain/polaris/eth/core/state/journal"
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// SelfDestructs is an autogenerated mock type for the SelfDestructs type
type SelfDestructs struct {
	mock.Mock
}

type SelfDestructs_Expecter struct {
	mock *mock.Mock
}

func (_m *SelfDestructs) EXPECT() *SelfDestructs_Expecter {
	return &SelfDestructs_Expecter{mock: &_m.Mock}
}

// Clone provides a mock function with given fields:
func (_m *SelfDestructs) Clone() journal.SelfDestructs {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clone")
	}

	var r0 journal.SelfDestructs
	if rf, ok := ret.Get(0).(func() journal.SelfDestructs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(journal.SelfDestructs)
		}
	}

	return r0
}

// SelfDestructs_Clone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clone'
type SelfDestructs_Clone_Call struct {
	*mock.Call
}

// Clone is a helper method to define mock.On call
func (_e *SelfDestructs_Expecter) Clone() *SelfDestructs_Clone_Call {
	return &SelfDestructs_Clone_Call{Call: _e.mock.On("Clone")}
}

func (_c *SelfDestructs_Clone_Call) Run(run func()) *SelfDestructs_Clone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SelfDestructs_Clone_Call) Return(_a0 journal.SelfDestructs) *SelfDestructs_Clone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructs_Clone_Call) RunAndReturn(run func() journal.SelfDestructs) *SelfDestructs_Clone_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields:
func (_m *SelfDestructs) Finalize() {
	_m.Called()
}

// SelfDestructs_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type SelfDestructs_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
func (_e *SelfDestructs_Expecter) Finalize() *SelfDestructs_Finalize_Call {
	return &SelfDestructs_Finalize_Call{Call: _e.mock.On("Finalize")}
}

func (_c *SelfDestructs_Finalize_Call) Run(run func()) *SelfDestructs_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SelfDestructs_Finalize_Call) Return() *SelfDestructs_Finalize_Call {
	_c.Call.Return()
	return _c
}

func (_c *SelfDestructs_Finalize_Call) RunAndReturn(run func()) *SelfDestructs_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// GetSelfDestructs provides a mock function with given fields:
func (_m *SelfDestructs) GetSelfDestructs() []common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSelfDestructs")
	}

	var r0 []common.Address
	if rf, ok := ret.Get(0).(func() []common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	return r0
}

// SelfDestructs_GetSelfDestructs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSelfDestructs'
type SelfDestructs_GetSelfDestructs_Call struct {
	*mock.Call
}

// GetSelfDestructs is a helper method to define mock.On call
func (_e *SelfDestructs_Expecter) GetSelfDestructs() *SelfDestructs_GetSelfDestructs_Call {
	return &SelfDestructs_GetSelfDestructs_Call{Call: _e.mock.On("GetSelfDestructs")}
}

func (_c *SelfDestructs_GetSelfDestructs_Call) Run(run func()) *SelfDestructs_GetSelfDestructs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SelfDestructs_GetSelfDestructs_Call) Return(_a0 []common.Address) *SelfDestructs_GetSelfDestructs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructs_GetSelfDestructs_Call) RunAndReturn(run func() []common.Address) *SelfDestructs_GetSelfDestructs_Call {
	_c.Call.Return(run)
	return _c
}

// HasSelfDestructed provides a mock function with given fields: _a0
func (_m *SelfDestructs) HasSelfDestructed(_a0 common.Address) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for HasSelfDestructed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SelfDestructs_HasSelfDestructed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasSelfDestructed'
type SelfDestructs_HasSelfDestructed_Call struct {
	*mock.Call
}

// HasSelfDestructed is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *SelfDestructs_Expecter) HasSelfDestructed(_a0 interface{}) *SelfDestructs_HasSelfDestructed_Call {
	return &SelfDestructs_HasSelfDestructed_Call{Call: _e.mock.On("HasSelfDestructed", _a0)}
}

func (_c *SelfDestructs_HasSelfDestructed_Call) Run(run func(_a0 common.Address)) *SelfDestructs_HasSelfDestructed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *SelfDestructs_HasSelfDestructed_Call) Return(_a0 bool) *SelfDestructs_HasSelfDestructed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructs_HasSelfDestructed_Call) RunAndReturn(run func(common.Address) bool) *SelfDestructs_HasSelfDestructed_Call {
	_c.Call.Return(run)
	return _c
}

// RegistryKey provides a mock function with given fields:
func (_m *SelfDestructs) RegistryKey() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RegistryKey")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SelfDestructs_RegistryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegistryKey'
type SelfDestructs_RegistryKey_Call struct {
	*mock.Call
}

// RegistryKey is a helper method to define mock.On call
func (_e *SelfDestructs_Expecter) RegistryKey() *SelfDestructs_RegistryKey_Call {
	return &SelfDestructs_RegistryKey_Call{Call: _e.mock.On("RegistryKey")}
}

func (_c *SelfDestructs_RegistryKey_Call) Run(run func()) *SelfDestructs_RegistryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SelfDestructs_RegistryKey_Call) Return(_a0 string) *SelfDestructs_RegistryKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructs_RegistryKey_Call) RunAndReturn(run func() string) *SelfDestructs_RegistryKey_Call {
	_c.Call.Return(run)
	return _c
}

// RevertToSnapshot provides a mock function with given fields: _a0
func (_m *SelfDestructs) RevertToSnapshot(_a0 int) {
	_m.Called(_a0)
}

// SelfDestructs_RevertToSnapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertToSnapshot'
type SelfDestructs_RevertToSnapshot_Call struct {
	*mock.Call
}

// RevertToSnapshot is a helper method to define mock.On call
//   - _a0 int
func (_e *SelfDestructs_Expecter) RevertToSnapshot(_a0 interface{}) *SelfDestructs_RevertToSnapshot_Call {
	return &SelfDestructs_RevertToSnapshot_Call{Call: _e.mock.On("RevertToSnapshot", _a0)}
}

func (_c *SelfDestructs_RevertToSnapshot_Call) Run(run func(_a0 int)) *SelfDestructs_RevertToSnapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *SelfDestructs_RevertToSnapshot_Call) Return() *SelfDestructs_RevertToSnapshot_Call {
	_c.Call.Return()
	return _c
}

func (_c *SelfDestructs_RevertToSnapshot_Call) RunAndReturn(run func(int)) *SelfDestructs_RevertToSnapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SelfDestruct provides a mock function with given fields: _a0
func (_m *SelfDestructs) SelfDestruct(_a0 common.Address) {
	_m.Called(_a0)
}

// SelfDestructs_SelfDestruct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelfDestruct'
type SelfDestructs_SelfDestruct_Call struct {
	*mock.Call
}

// SelfDestruct is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *SelfDestructs_Expecter) SelfDestruct(_a0 interface{}) *SelfDestructs_SelfDestruct_Call {
	return &SelfDestructs_SelfDestruct_Call{Call: _e.mock.On("SelfDestruct", _a0)}
}

func (_c *SelfDestructs_SelfDestruct_Call) Run(run func(_a0 common.Address)) *SelfDestructs_SelfDestruct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *SelfDestructs_SelfDestruct_Call) Return() *SelfDestructs_SelfDestruct_Call {
	_c.Call.Return()
	return _c
}

func (_c *SelfDestructs_SelfDestruct_Call) RunAndReturn(run func(common.Address)) *SelfDestructs_SelfDestruct_Call {
	_c.Call.Return(run)
	return _c
}

// Selfdestruct6780 provides a mock function with given fields: _a0
func (_m *SelfDestructs) Selfdestruct6780(_a0 common.Address) {
	_m.Called(_a0)
}

// SelfDestructs_Selfdestruct6780_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Selfdestruct6780'
type SelfDestructs_Selfdestruct6780_Call struct {
	*mock.Call
}

// Selfdestruct6780 is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *SelfDestructs_Expecter) Selfdestruct6780(_a0 interface{}) *SelfDestructs_Selfdestruct6780_Call {
	return &SelfDestructs_Selfdestruct6780_Call{Call: _e.mock.On("Selfdestruct6780", _a0)}
}

func (_c *SelfDestructs_Selfdestruct6780_Call) Run(run func(_a0 common.Address)) *SelfDestructs_Selfdestruct6780_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *SelfDestructs_Selfdestruct6780_Call) Return() *SelfDestructs_Selfdestruct6780_Call {
	_c.Call.Return()
	return _c
}

func (_c *SelfDestructs_Selfdestruct6780_Call) RunAndReturn(run func(common.Address)) *SelfDestructs_Selfdestruct6780_Call {
	_c.Call.Return(run)
	return _c
}

// Snapshot provides a mock function with given fields:
func (_m *SelfDestructs) Snapshot() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Snapshot")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SelfDestructs_Snapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Snapshot'
type SelfDestructs_Snapshot_Call struct {
	*mock.Call
}

// Snapshot is a helper method to define mock.On call
func (_e *SelfDestructs_Expecter) Snapshot() *SelfDestructs_Snapshot_Call {
	return &SelfDestructs_Snapshot_Call{Call: _e.mock.On("Snapshot")}
}

func (_c *SelfDestructs_Snapshot_Call) Run(run func()) *SelfDestructs_Snapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SelfDestructs_Snapshot_Call) Return(_a0 int) *SelfDestructs_Snapshot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructs_Snapshot_Call) RunAndReturn(run func() int) *SelfDestructs_Snapshot_Call {
	_c.Call.Return(run)
	return _c
}

// NewSelfDestructs creates a new instance of SelfDestructs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSelfDestructs(t interface {
	mock.TestingT
	Cleanup(func())
}) *SelfDestructs {
	mock := &SelfDestructs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
