// Code generated by MockGen. DO NOT EDIT.
// Source: customer.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "github.com/tusmasoma/microservice-k8s-demo/customer/entity"
	usecase "github.com/tusmasoma/microservice-k8s-demo/customer/usecase"
)

// MockCustomerUsecase is a mock of CustomerUsecase interface.
type MockCustomerUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerUsecaseMockRecorder
}

// MockCustomerUsecaseMockRecorder is the mock recorder for MockCustomerUsecase.
type MockCustomerUsecaseMockRecorder struct {
	mock *MockCustomerUsecase
}

// NewMockCustomerUsecase creates a new mock instance.
func NewMockCustomerUsecase(ctrl *gomock.Controller) *MockCustomerUsecase {
	mock := &MockCustomerUsecase{ctrl: ctrl}
	mock.recorder = &MockCustomerUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerUsecase) EXPECT() *MockCustomerUsecaseMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method.
func (m *MockCustomerUsecase) CreateCustomer(ctx context.Context, params *usecase.CreateCustomerParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockCustomerUsecaseMockRecorder) CreateCustomer(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockCustomerUsecase)(nil).CreateCustomer), ctx, params)
}

// DeleteCustomer mocks base method.
func (m *MockCustomerUsecase) DeleteCustomer(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomer indicates an expected call of DeleteCustomer.
func (mr *MockCustomerUsecaseMockRecorder) DeleteCustomer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomer", reflect.TypeOf((*MockCustomerUsecase)(nil).DeleteCustomer), ctx, id)
}

// GetCustomer mocks base method.
func (m *MockCustomerUsecase) GetCustomer(ctx context.Context, id string) (*entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", ctx, id)
	ret0, _ := ret[0].(*entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockCustomerUsecaseMockRecorder) GetCustomer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockCustomerUsecase)(nil).GetCustomer), ctx, id)
}

// ListCustomers mocks base method.
func (m *MockCustomerUsecase) ListCustomers(ctx context.Context) ([]entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCustomers", ctx)
	ret0, _ := ret[0].([]entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCustomers indicates an expected call of ListCustomers.
func (mr *MockCustomerUsecaseMockRecorder) ListCustomers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCustomers", reflect.TypeOf((*MockCustomerUsecase)(nil).ListCustomers), ctx)
}

// UpdateCustomer mocks base method.
func (m *MockCustomerUsecase) UpdateCustomer(ctx context.Context, params *usecase.UpdateCustomerParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomer", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomer indicates an expected call of UpdateCustomer.
func (mr *MockCustomerUsecaseMockRecorder) UpdateCustomer(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomer", reflect.TypeOf((*MockCustomerUsecase)(nil).UpdateCustomer), ctx, params)
}