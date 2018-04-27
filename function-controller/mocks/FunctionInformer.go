// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import cache "k8s.io/client-go/tools/cache"
import mock "github.com/stretchr/testify/mock"
import projectriffv1alpha1 "github.com/projectriff/riff/kubernetes-crds/pkg/client/listers/projectriff/v1alpha1"

// FunctionInformer is an autogenerated mock type for the FunctionInformer type
type FunctionInformer struct {
	mock.Mock
}

// Informer provides a mock function with given fields:
func (_m *FunctionInformer) Informer() cache.SharedIndexInformer {
	ret := _m.Called()

	var r0 cache.SharedIndexInformer
	if rf, ok := ret.Get(0).(func() cache.SharedIndexInformer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.SharedIndexInformer)
		}
	}

	return r0
}

// Lister provides a mock function with given fields:
func (_m *FunctionInformer) Lister() projectriffv1alpha1.FunctionLister {
	ret := _m.Called()

	var r0 projectriffv1alpha1.FunctionLister
	if rf, ok := ret.Get(0).(func() projectriffv1alpha1.FunctionLister); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(projectriffv1alpha1.FunctionLister)
		}
	}

	return r0
}
