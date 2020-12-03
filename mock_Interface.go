// Code generated by mockery v1.0.0. DO NOT EDIT.
package vslack

import mock "github.com/stretchr/testify/mock"

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

// validate provides a mock function with given fields:
func (_m *MockInterface) validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields:
func (_m *MockInterface) Send() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendAsync provides a mock function with given fields: e
func (_m *MockInterface) SendAsync(e chan error) {
	_m.Called(e)
}

// SetAttachments provides a mock function with given fields: a
func (_m *MockInterface) SetAttachments(a ...Attachment) *VSlack {
	_va := make([]interface{}, len(a))
	for _i := range a {
		_va[_i] = a[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *VSlack
	if rf, ok := ret.Get(0).(func(...Attachment) *VSlack); ok {
		r0 = rf(a...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VSlack)
		}
	}

	return r0
}

// SetChannel provides a mock function with given fields: c
func (_m *MockInterface) SetChannel(c string) *VSlack {
	ret := _m.Called(c)

	var r0 *VSlack
	if rf, ok := ret.Get(0).(func(string) *VSlack); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VSlack)
		}
	}

	return r0
}

// SetIconEmoji provides a mock function with given fields: i
func (_m *MockInterface) SetIconEmoji(i string) *VSlack {
	ret := _m.Called(i)

	var r0 *VSlack
	if rf, ok := ret.Get(0).(func(string) *VSlack); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VSlack)
		}
	}

	return r0
}

// SetIncomingWebhookURI provides a mock function with given fields: h
func (_m *MockInterface) SetIncomingWebhookURI(h string) *VSlack {
	ret := _m.Called(h)

	var r0 *VSlack
	if rf, ok := ret.Get(0).(func(string) *VSlack); ok {
		r0 = rf(h)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VSlack)
		}
	}

	return r0
}

// SetMessage provides a mock function with given fields: m
func (_m *MockInterface) SetMessage(m string) *VSlack {
	ret := _m.Called(m)

	var r0 *VSlack
	if rf, ok := ret.Get(0).(func(string) *VSlack); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VSlack)
		}
	}

	return r0
}

// SetUsername provides a mock function with given fields: u
func (_m *MockInterface) SetUsername(u string) *VSlack {
	ret := _m.Called(u)

	var r0 *VSlack
	if rf, ok := ret.Get(0).(func(string) *VSlack); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VSlack)
		}
	}

	return r0
}