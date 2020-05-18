// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	context "context"

	eth "github.com/smartcontractkit/chainlink/core/eth"
	mock "github.com/stretchr/testify/mock"
)

// CallerSubscriber is an autogenerated mock type for the CallerSubscriber type
type CallerSubscriber struct {
	mock.Mock
}

// Call provides a mock function with given fields: result, method, args
func (_m *CallerSubscriber) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: _a0, _a1, _a2
func (_m *CallerSubscriber) Subscribe(_a0 context.Context, _a1 interface{}, _a2 ...interface{}) (eth.Subscription, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) eth.Subscription); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
