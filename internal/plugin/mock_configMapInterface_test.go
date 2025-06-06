// Code generated by mockery v2.53.3. DO NOT EDIT.

package plugin

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/core/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// mockConfigMapInterface is an autogenerated mock type for the configMapInterface type
type mockConfigMapInterface struct {
	mock.Mock
}

type mockConfigMapInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *mockConfigMapInterface) EXPECT() *mockConfigMapInterface_Expecter {
	return &mockConfigMapInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, configMap, opts
func (_m *mockConfigMapInterface) Apply(ctx context.Context, configMap *v1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ConfigMap, error) {
	ret := _m.Called(ctx, configMap, opts)

	if len(ret) == 0 {
		panic("no return value specified for Apply")
	}

	var r0 *corev1.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ConfigMapApplyConfiguration, metav1.ApplyOptions) (*corev1.ConfigMap, error)); ok {
		return rf(ctx, configMap, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ConfigMapApplyConfiguration, metav1.ApplyOptions) *corev1.ConfigMap); ok {
		r0 = rf(ctx, configMap, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ConfigMapApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, configMap, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type mockConfigMapInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - configMap *v1.ConfigMapApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *mockConfigMapInterface_Expecter) Apply(ctx interface{}, configMap interface{}, opts interface{}) *mockConfigMapInterface_Apply_Call {
	return &mockConfigMapInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, configMap, opts)}
}

func (_c *mockConfigMapInterface_Apply_Call) Run(run func(ctx context.Context, configMap *v1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions)) *mockConfigMapInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ConfigMapApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_Apply_Call) Return(result *corev1.ConfigMap, err error) *mockConfigMapInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *mockConfigMapInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.ConfigMapApplyConfiguration, metav1.ApplyOptions) (*corev1.ConfigMap, error)) *mockConfigMapInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, configMap, opts
func (_m *mockConfigMapInterface) Create(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.CreateOptions) (*corev1.ConfigMap, error) {
	ret := _m.Called(ctx, configMap, opts)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *corev1.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.ConfigMap, metav1.CreateOptions) (*corev1.ConfigMap, error)); ok {
		return rf(ctx, configMap, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.ConfigMap, metav1.CreateOptions) *corev1.ConfigMap); ok {
		r0 = rf(ctx, configMap, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.ConfigMap, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, configMap, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type mockConfigMapInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - configMap *corev1.ConfigMap
//   - opts metav1.CreateOptions
func (_e *mockConfigMapInterface_Expecter) Create(ctx interface{}, configMap interface{}, opts interface{}) *mockConfigMapInterface_Create_Call {
	return &mockConfigMapInterface_Create_Call{Call: _e.mock.On("Create", ctx, configMap, opts)}
}

func (_c *mockConfigMapInterface_Create_Call) Run(run func(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.CreateOptions)) *mockConfigMapInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.ConfigMap), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_Create_Call) Return(_a0 *corev1.ConfigMap, _a1 error) *mockConfigMapInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockConfigMapInterface_Create_Call) RunAndReturn(run func(context.Context, *corev1.ConfigMap, metav1.CreateOptions) (*corev1.ConfigMap, error)) *mockConfigMapInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *mockConfigMapInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockConfigMapInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type mockConfigMapInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *mockConfigMapInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *mockConfigMapInterface_Delete_Call {
	return &mockConfigMapInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *mockConfigMapInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *mockConfigMapInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_Delete_Call) Return(_a0 error) *mockConfigMapInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockConfigMapInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *mockConfigMapInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *mockConfigMapInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockConfigMapInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type mockConfigMapInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *mockConfigMapInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *mockConfigMapInterface_DeleteCollection_Call {
	return &mockConfigMapInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *mockConfigMapInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *mockConfigMapInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_DeleteCollection_Call) Return(_a0 error) *mockConfigMapInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockConfigMapInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *mockConfigMapInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *mockConfigMapInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
	ret := _m.Called(ctx, name, opts)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *corev1.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*corev1.ConfigMap, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.ConfigMap); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type mockConfigMapInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *mockConfigMapInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *mockConfigMapInterface_Get_Call {
	return &mockConfigMapInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *mockConfigMapInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *mockConfigMapInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_Get_Call) Return(_a0 *corev1.ConfigMap, _a1 error) *mockConfigMapInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockConfigMapInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*corev1.ConfigMap, error)) *mockConfigMapInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *mockConfigMapInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *corev1.ConfigMapList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*corev1.ConfigMapList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.ConfigMapList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ConfigMapList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type mockConfigMapInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *mockConfigMapInterface_Expecter) List(ctx interface{}, opts interface{}) *mockConfigMapInterface_List_Call {
	return &mockConfigMapInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *mockConfigMapInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *mockConfigMapInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_List_Call) Return(_a0 *corev1.ConfigMapList, _a1 error) *mockConfigMapInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockConfigMapInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*corev1.ConfigMapList, error)) *mockConfigMapInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *mockConfigMapInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.ConfigMap, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Patch")
	}

	var r0 *corev1.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.ConfigMap, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.ConfigMap); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type mockConfigMapInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *mockConfigMapInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *mockConfigMapInterface_Patch_Call {
	return &mockConfigMapInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *mockConfigMapInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *mockConfigMapInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(metav1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *mockConfigMapInterface_Patch_Call) Return(result *corev1.ConfigMap, err error) *mockConfigMapInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *mockConfigMapInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.ConfigMap, error)) *mockConfigMapInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, configMap, opts
func (_m *mockConfigMapInterface) Update(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.UpdateOptions) (*corev1.ConfigMap, error) {
	ret := _m.Called(ctx, configMap, opts)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *corev1.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.ConfigMap, metav1.UpdateOptions) (*corev1.ConfigMap, error)); ok {
		return rf(ctx, configMap, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.ConfigMap, metav1.UpdateOptions) *corev1.ConfigMap); ok {
		r0 = rf(ctx, configMap, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.ConfigMap, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, configMap, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type mockConfigMapInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - configMap *corev1.ConfigMap
//   - opts metav1.UpdateOptions
func (_e *mockConfigMapInterface_Expecter) Update(ctx interface{}, configMap interface{}, opts interface{}) *mockConfigMapInterface_Update_Call {
	return &mockConfigMapInterface_Update_Call{Call: _e.mock.On("Update", ctx, configMap, opts)}
}

func (_c *mockConfigMapInterface_Update_Call) Run(run func(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.UpdateOptions)) *mockConfigMapInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.ConfigMap), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_Update_Call) Return(_a0 *corev1.ConfigMap, _a1 error) *mockConfigMapInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockConfigMapInterface_Update_Call) RunAndReturn(run func(context.Context, *corev1.ConfigMap, metav1.UpdateOptions) (*corev1.ConfigMap, error)) *mockConfigMapInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *mockConfigMapInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for Watch")
	}

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigMapInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type mockConfigMapInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *mockConfigMapInterface_Expecter) Watch(ctx interface{}, opts interface{}) *mockConfigMapInterface_Watch_Call {
	return &mockConfigMapInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *mockConfigMapInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *mockConfigMapInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *mockConfigMapInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *mockConfigMapInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockConfigMapInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *mockConfigMapInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// newMockConfigMapInterface creates a new instance of mockConfigMapInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockConfigMapInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockConfigMapInterface {
	mock := &mockConfigMapInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
