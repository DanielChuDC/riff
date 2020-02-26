// Code generated by mockery v1.0.0. DO NOT EDIT.

package kail

import (
	context "context"
	io "io"

	corev1alpha1 "github.com/projectriff/riff/system/pkg/apis/core/v1alpha1"

	knativev1alpha1 "github.com/projectriff/riff/system/pkg/apis/knative/v1alpha1"

	mock "github.com/stretchr/testify/mock"

	streamingv1alpha1 "github.com/projectriff/riff/system/pkg/apis/streaming/v1alpha1"

	time "time"

	v1alpha1 "github.com/projectriff/riff/system/pkg/apis/build/v1alpha1"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// ApplicationLogs provides a mock function with given fields: ctx, application, since, out
func (_m *Logger) ApplicationLogs(ctx context.Context, application *v1alpha1.Application, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, application, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Application, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, application, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CoreDeployerLogs provides a mock function with given fields: ctx, deployer, since, out
func (_m *Logger) CoreDeployerLogs(ctx context.Context, deployer *corev1alpha1.Deployer, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, deployer, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1alpha1.Deployer, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, deployer, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FunctionLogs provides a mock function with given fields: ctx, function, since, out
func (_m *Logger) FunctionLogs(ctx context.Context, function *v1alpha1.Function, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, function, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Function, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, function, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InMemoryGatewayLogs provides a mock function with given fields: ctx, gateway, since, out
func (_m *Logger) InMemoryGatewayLogs(ctx context.Context, gateway *streamingv1alpha1.InMemoryGateway, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, gateway, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *streamingv1alpha1.InMemoryGateway, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, gateway, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KafkaGatewayLogs provides a mock function with given fields: ctx, gateway, since, out
func (_m *Logger) KafkaGatewayLogs(ctx context.Context, gateway *streamingv1alpha1.KafkaGateway, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, gateway, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *streamingv1alpha1.KafkaGateway, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, gateway, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KnativeDeployerLogs provides a mock function with given fields: ctx, deployer, since, out
func (_m *Logger) KnativeDeployerLogs(ctx context.Context, deployer *knativev1alpha1.Deployer, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, deployer, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *knativev1alpha1.Deployer, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, deployer, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PulsarGatewayLogs provides a mock function with given fields: ctx, gateway, since, out
func (_m *Logger) PulsarGatewayLogs(ctx context.Context, gateway *streamingv1alpha1.PulsarGateway, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, gateway, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *streamingv1alpha1.PulsarGateway, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, gateway, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StreamingProcessorLogs provides a mock function with given fields: ctx, processor, since, out
func (_m *Logger) StreamingProcessorLogs(ctx context.Context, processor *streamingv1alpha1.Processor, since time.Duration, out io.Writer) error {
	ret := _m.Called(ctx, processor, since, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *streamingv1alpha1.Processor, time.Duration, io.Writer) error); ok {
		r0 = rf(ctx, processor, since, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
