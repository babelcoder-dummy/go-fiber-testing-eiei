// Code generated by mockery v2.36.1. DO NOT EDIT.

package mock_service

import (
	dto "github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	mock "github.com/stretchr/testify/mock"

	model "github.com/babelcoder-enterprise-courses/go-fiber-testing/model"

	multipart "mime/multipart"
)

// MockProducter is an autogenerated mock type for the Producter type
type MockProducter struct {
	mock.Mock
}

type MockProducter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProducter) EXPECT() *MockProducter_Expecter {
	return &MockProducter_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockProducter) Create(_a0 *dto.CreateProductForm, _a1 *multipart.FileHeader) (*model.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.CreateProductForm, *multipart.FileHeader) (*model.Product, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(*dto.CreateProductForm, *multipart.FileHeader) *model.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.CreateProductForm, *multipart.FileHeader) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProducter_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockProducter_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 *dto.CreateProductForm
//   - _a1 *multipart.FileHeader
func (_e *MockProducter_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockProducter_Create_Call {
	return &MockProducter_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockProducter_Create_Call) Run(run func(_a0 *dto.CreateProductForm, _a1 *multipart.FileHeader)) *MockProducter_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dto.CreateProductForm), args[1].(*multipart.FileHeader))
	})
	return _c
}

func (_c *MockProducter_Create_Call) Return(_a0 *model.Product, _a1 error) *MockProducter_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProducter_Create_Call) RunAndReturn(run func(*dto.CreateProductForm, *multipart.FileHeader) (*model.Product, error)) *MockProducter_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *MockProducter) Delete(_a0 uint) {
	_m.Called(_a0)
}

// MockProducter_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockProducter_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 uint
func (_e *MockProducter_Expecter) Delete(_a0 interface{}) *MockProducter_Delete_Call {
	return &MockProducter_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *MockProducter_Delete_Call) Run(run func(_a0 uint)) *MockProducter_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockProducter_Delete_Call) Return() *MockProducter_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockProducter_Delete_Call) RunAndReturn(run func(uint)) *MockProducter_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: _a0
func (_m *MockProducter) FindAll(_a0 string) []model.Product {
	ret := _m.Called(_a0)

	var r0 []model.Product
	if rf, ok := ret.Get(0).(func(string) []model.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Product)
		}
	}

	return r0
}

// MockProducter_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockProducter_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - _a0 string
func (_e *MockProducter_Expecter) FindAll(_a0 interface{}) *MockProducter_FindAll_Call {
	return &MockProducter_FindAll_Call{Call: _e.mock.On("FindAll", _a0)}
}

func (_c *MockProducter_FindAll_Call) Run(run func(_a0 string)) *MockProducter_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProducter_FindAll_Call) Return(_a0 []model.Product) *MockProducter_FindAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProducter_FindAll_Call) RunAndReturn(run func(string) []model.Product) *MockProducter_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: _a0
func (_m *MockProducter) FindOne(_a0 uint) (*model.Product, error) {
	ret := _m.Called(_a0)

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*model.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) *model.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProducter_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockProducter_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - _a0 uint
func (_e *MockProducter_Expecter) FindOne(_a0 interface{}) *MockProducter_FindOne_Call {
	return &MockProducter_FindOne_Call{Call: _e.mock.On("FindOne", _a0)}
}

func (_c *MockProducter_FindOne_Call) Run(run func(_a0 uint)) *MockProducter_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockProducter_FindOne_Call) Return(_a0 *model.Product, _a1 error) *MockProducter_FindOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProducter_FindOne_Call) RunAndReturn(run func(uint) (*model.Product, error)) *MockProducter_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockProducter) Update(_a0 uint, _a1 *multipart.FileHeader, _a2 *dto.UpdateProductForm) (*model.Product, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, *multipart.FileHeader, *dto.UpdateProductForm) (*model.Product, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(uint, *multipart.FileHeader, *dto.UpdateProductForm) *model.Product); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, *multipart.FileHeader, *dto.UpdateProductForm) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProducter_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockProducter_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 uint
//   - _a1 *multipart.FileHeader
//   - _a2 *dto.UpdateProductForm
func (_e *MockProducter_Expecter) Update(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockProducter_Update_Call {
	return &MockProducter_Update_Call{Call: _e.mock.On("Update", _a0, _a1, _a2)}
}

func (_c *MockProducter_Update_Call) Run(run func(_a0 uint, _a1 *multipart.FileHeader, _a2 *dto.UpdateProductForm)) *MockProducter_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(*multipart.FileHeader), args[2].(*dto.UpdateProductForm))
	})
	return _c
}

func (_c *MockProducter_Update_Call) Return(_a0 *model.Product, _a1 error) *MockProducter_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProducter_Update_Call) RunAndReturn(run func(uint, *multipart.FileHeader, *dto.UpdateProductForm) (*model.Product, error)) *MockProducter_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProducter creates a new instance of MockProducter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProducter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProducter {
	mock := &MockProducter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
