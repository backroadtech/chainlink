// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"
)

// Fetcher is an autogenerated mock type for the Fetcher type
type Fetcher struct {
	mock.Mock
}

// Fetch provides a mock function with given fields:
func (_m *Fetcher) Fetch() (decimal.Decimal, error) {
	ret := _m.Called()

	var r0 decimal.Decimal
	if rf, ok := ret.Get(0).(func() decimal.Decimal); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
