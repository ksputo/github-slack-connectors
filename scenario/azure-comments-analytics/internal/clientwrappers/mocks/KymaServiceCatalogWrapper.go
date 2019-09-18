// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import k8scomponents "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
import mock "github.com/stretchr/testify/mock"

// KymaServiceCatalogWrapper is an autogenerated mock type for the KymaServiceCatalogWrapper type
type KymaServiceCatalogWrapper struct {
	mock.Mock
}

// BindingUsage provides a mock function with given fields: namespace
func (_m *KymaServiceCatalogWrapper) BindingUsage(namespace string) k8scomponents.BindingUsage {
	ret := _m.Called(namespace)

	var r0 k8scomponents.BindingUsage
	if rf, ok := ret.Get(0).(func(string) k8scomponents.BindingUsage); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(k8scomponents.BindingUsage)
		}
	}

	return r0
}
