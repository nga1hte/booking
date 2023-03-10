// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nga1hte/booking/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	db "github.com/nga1hte/booking/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// BookingTx mocks base method.
func (m *MockStore) BookingTx(arg0 context.Context, arg1 db.BookingTxParams) (db.BookingTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookingTx", arg0, arg1)
	ret0, _ := ret[0].(db.BookingTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookingTx indicates an expected call of BookingTx.
func (mr *MockStoreMockRecorder) BookingTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookingTx", reflect.TypeOf((*MockStore)(nil).BookingTx), arg0, arg1)
}

// CreateBooking mocks base method.
func (m *MockStore) CreateBooking(arg0 context.Context, arg1 db.CreateBookingParams) (db.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBooking", arg0, arg1)
	ret0, _ := ret[0].(db.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBooking indicates an expected call of CreateBooking.
func (mr *MockStoreMockRecorder) CreateBooking(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBooking", reflect.TypeOf((*MockStore)(nil).CreateBooking), arg0, arg1)
}

// CreatePayment mocks base method.
func (m *MockStore) CreatePayment(arg0 context.Context, arg1 int64) (db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", arg0, arg1)
	ret0, _ := ret[0].(db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockStoreMockRecorder) CreatePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockStore)(nil).CreatePayment), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteBooking mocks base method.
func (m *MockStore) DeleteBooking(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBooking", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBooking indicates an expected call of DeleteBooking.
func (mr *MockStoreMockRecorder) DeleteBooking(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBooking", reflect.TypeOf((*MockStore)(nil).DeleteBooking), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetBooking mocks base method.
func (m *MockStore) GetBooking(arg0 context.Context, arg1 int64) (db.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooking", arg0, arg1)
	ret0, _ := ret[0].(db.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooking indicates an expected call of GetBooking.
func (mr *MockStoreMockRecorder) GetBooking(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooking", reflect.TypeOf((*MockStore)(nil).GetBooking), arg0, arg1)
}

// GetBookingPayment mocks base method.
func (m *MockStore) GetBookingPayment(arg0 context.Context, arg1 int64) (db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookingPayment", arg0, arg1)
	ret0, _ := ret[0].(db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookingPayment indicates an expected call of GetBookingPayment.
func (mr *MockStoreMockRecorder) GetBookingPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookingPayment", reflect.TypeOf((*MockStore)(nil).GetBookingPayment), arg0, arg1)
}

// GetBookings mocks base method.
func (m *MockStore) GetBookings(arg0 context.Context, arg1 db.GetBookingsParams) ([]db.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookings", arg0, arg1)
	ret0, _ := ret[0].([]db.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookings indicates an expected call of GetBookings.
func (mr *MockStoreMockRecorder) GetBookings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockStore)(nil).GetBookings), arg0, arg1)
}

// GetBookingsFromToday mocks base method.
func (m *MockStore) GetBookingsFromToday(arg0 context.Context, arg1 time.Time) ([]db.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookingsFromToday", arg0, arg1)
	ret0, _ := ret[0].([]db.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookingsFromToday indicates an expected call of GetBookingsFromToday.
func (mr *MockStoreMockRecorder) GetBookingsFromToday(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookingsFromToday", reflect.TypeOf((*MockStore)(nil).GetBookingsFromToday), arg0, arg1)
}

// GetPayments mocks base method.
func (m *MockStore) GetPayments(arg0 context.Context, arg1 db.GetPaymentsParams) ([]db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayments", arg0, arg1)
	ret0, _ := ret[0].([]db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayments indicates an expected call of GetPayments.
func (mr *MockStoreMockRecorder) GetPayments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayments", reflect.TypeOf((*MockStore)(nil).GetPayments), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserBookings mocks base method.
func (m *MockStore) GetUserBookings(arg0 context.Context, arg1 db.GetUserBookingsParams) ([]db.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBookings", arg0, arg1)
	ret0, _ := ret[0].([]db.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBookings indicates an expected call of GetUserBookings.
func (mr *MockStoreMockRecorder) GetUserBookings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBookings", reflect.TypeOf((*MockStore)(nil).GetUserBookings), arg0, arg1)
}

// GetUserLogin mocks base method.
func (m *MockStore) GetUserLogin(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLogin", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLogin indicates an expected call of GetUserLogin.
func (mr *MockStoreMockRecorder) GetUserLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLogin", reflect.TypeOf((*MockStore)(nil).GetUserLogin), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockStore) GetUsers(arg0 context.Context, arg1 db.GetUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStore)(nil).GetUsers), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
