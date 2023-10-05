package machinehealthcheck

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlfake "sigs.k8s.io/controller-runtime/pkg/client/fake"

	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	mock_dynamichelper "github.com/Azure/ARO-RP/pkg/util/mocks/dynamichelper"
	_ "github.com/Azure/ARO-RP/pkg/util/scheme"
	utilconditions "github.com/Azure/ARO-RP/test/util/conditions"
	utilerror "github.com/Azure/ARO-RP/test/util/error"
	operatorv1 "github.com/openshift/api/operator/v1"
)

// Test reconcile function
func TestMachineHealthCheckReconciler(t *testing.T) {
	transitionTime := metav1.Time{Time: time.Now()}
	defaultAvailable := utilconditions.ControllerDefaultAvailable(MHCControllerName)
	defaultProgressing := utilconditions.ControllerDefaultProgressing(MHCControllerName)
	defaultDegraded := utilconditions.ControllerDefaultDegraded(MHCControllerName)

	defaultConditions := []operatorv1.OperatorCondition{defaultAvailable, defaultProgressing, defaultDegraded}

	type test struct {
		name             string
		instance         *arov1alpha1.Cluster
		mocks            func(mdh *mock_dynamichelper.MockInterface)
		wantConditions   []operatorv1.OperatorCondition
		wantErr          string
		wantRequeueAfter time.Duration
	}

	for _, tt := range []*test{
		{
			name:           "Failure to get instance",
			mocks:          func(mdh *mock_dynamichelper.MockInterface) {},
			wantConditions: defaultConditions,
			wantErr:        `clusters.aro.openshift.io "cluster" not found`,
		},
		{
			name: "Enabled Feature Flag is false",
			instance: &arov1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: arov1alpha1.SingletonClusterName,
				},
				Spec: arov1alpha1.ClusterSpec{
					OperatorFlags: arov1alpha1.OperatorFlags{
						MHCEnabled: strconv.FormatBool(false),
					},
				},
				Status: arov1alpha1.ClusterStatus{
					Conditions: defaultConditions,
				},
			},
			mocks: func(mdh *mock_dynamichelper.MockInterface) {
				mdh.EXPECT().EnsureDeleted(gomock.Any(), "MachineHealthCheck", "openshift-machine-api", "aro-machinehealthcheck").Times(0)
				mdh.EXPECT().Ensure(gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
			wantConditions: defaultConditions,
			wantErr:        "",
		},
		{
			name: "Managed Feature Flag is false: ensure mhc and its alert are deleted",
			instance: &arov1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: arov1alpha1.SingletonClusterName,
				},
				Spec: arov1alpha1.ClusterSpec{
					OperatorFlags: arov1alpha1.OperatorFlags{
						MHCEnabled: strconv.FormatBool(true),
						MHCManaged: strconv.FormatBool(false),
					},
				},
				Status: arov1alpha1.ClusterStatus{
					Conditions: defaultConditions,
				},
			},
			mocks: func(mdh *mock_dynamichelper.MockInterface) {
				mdh.EXPECT().EnsureDeleted(gomock.Any(), "MachineHealthCheck", "openshift-machine-api", "aro-machinehealthcheck").Times(1)
				mdh.EXPECT().EnsureDeleted(gomock.Any(), "PrometheusRule", "openshift-machine-api", "mhc-remediation-alert").Times(1)
			},
			wantConditions: defaultConditions,
			wantErr:        "",
		},
		{
			name: "Managed Feature Flag is false: mhc fails to delete, an error is returned",
			instance: &arov1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: arov1alpha1.SingletonClusterName,
				},
				Spec: arov1alpha1.ClusterSpec{
					OperatorFlags: arov1alpha1.OperatorFlags{
						MHCEnabled: strconv.FormatBool(true),
						MHCManaged: strconv.FormatBool(false),
					},
				},
				Status: arov1alpha1.ClusterStatus{
					Conditions: defaultConditions,
				},
			},
			mocks: func(mdh *mock_dynamichelper.MockInterface) {
				mdh.EXPECT().EnsureDeleted(gomock.Any(), "MachineHealthCheck", "openshift-machine-api", "aro-machinehealthcheck").Return(errors.New("Could not delete mhc"))
			},
			wantErr: "Could not delete mhc",
			wantConditions: []operatorv1.OperatorCondition{
				defaultAvailable,
				defaultProgressing,
				{
					Type:               MHCControllerName + "Controller" + operatorv1.OperatorStatusTypeDegraded,
					Status:             operatorv1.ConditionTrue,
					LastTransitionTime: transitionTime,
					Message:            "Could not delete mhc",
				},
			},
			wantRequeueAfter: time.Hour,
		},
		{
			name: "Managed Feature Flag is false: mhc deletes but mhc alert fails to delete, an error is returned",
			instance: &arov1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: arov1alpha1.SingletonClusterName,
				},
				Spec: arov1alpha1.ClusterSpec{
					OperatorFlags: arov1alpha1.OperatorFlags{
						MHCEnabled: strconv.FormatBool(true),
						MHCManaged: strconv.FormatBool(false),
					},
				},
				Status: arov1alpha1.ClusterStatus{
					Conditions: defaultConditions,
				},
			},
			mocks: func(mdh *mock_dynamichelper.MockInterface) {
				mdh.EXPECT().EnsureDeleted(gomock.Any(), "MachineHealthCheck", "openshift-machine-api", "aro-machinehealthcheck").Times(1)
				mdh.EXPECT().EnsureDeleted(gomock.Any(), "PrometheusRule", "openshift-machine-api", "mhc-remediation-alert").Return(errors.New("Could not delete mhc alert"))
			},
			wantErr: "Could not delete mhc alert",
			wantConditions: []operatorv1.OperatorCondition{
				defaultAvailable,
				defaultProgressing,
				{
					Type:               MHCControllerName + "Controller" + operatorv1.OperatorStatusTypeDegraded,
					Status:             operatorv1.ConditionTrue,
					LastTransitionTime: transitionTime,
					Message:            "Could not delete mhc alert",
				},
			},
			wantRequeueAfter: time.Hour,
		},
		{
			name: "Managed Feature Flag is true: dynamic helper ensures resources",
			instance: &arov1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: arov1alpha1.SingletonClusterName,
				},
				Spec: arov1alpha1.ClusterSpec{
					OperatorFlags: arov1alpha1.OperatorFlags{
						MHCEnabled: strconv.FormatBool(true),
						MHCManaged: strconv.FormatBool(true),
					},
				},
			},
			mocks: func(mdh *mock_dynamichelper.MockInterface) {
				mdh.EXPECT().Ensure(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantConditions: defaultConditions,
			wantErr:        "",
		},
		{
			name: "When ensuring resources fails, an error is returned",
			instance: &arov1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: arov1alpha1.SingletonClusterName,
				},
				Spec: arov1alpha1.ClusterSpec{
					OperatorFlags: arov1alpha1.OperatorFlags{
						MHCEnabled: strconv.FormatBool(true),
						MHCManaged: strconv.FormatBool(true),
					},
				},
				Status: arov1alpha1.ClusterStatus{
					Conditions: defaultConditions,
				},
			},
			mocks: func(mdh *mock_dynamichelper.MockInterface) {
				mdh.EXPECT().Ensure(gomock.Any(), gomock.Any()).Return(errors.New("failed to ensure"))
			},
			wantErr: "failed to ensure",
			wantConditions: []operatorv1.OperatorCondition{
				defaultAvailable,
				defaultProgressing,
				{
					Type:               MHCControllerName + "Controller" + operatorv1.OperatorStatusTypeDegraded,
					Status:             operatorv1.ConditionTrue,
					LastTransitionTime: transitionTime,
					Message:            "failed to ensure",
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mdh := mock_dynamichelper.NewMockInterface(controller)

			tt.mocks(mdh)

			clientBuilder := ctrlfake.NewClientBuilder()
			if tt.instance != nil {
				clientBuilder = clientBuilder.WithObjects(tt.instance)
			}

			ctx := context.Background()

			r := NewMachineHealthCheckReconciler(
				logrus.NewEntry(logrus.StandardLogger()),
				clientBuilder.Build(),
				mdh,
			)

			request := ctrl.Request{}
			request.Name = "cluster"

			result, err := r.Reconcile(ctx, request)

			if tt.wantRequeueAfter != result.RequeueAfter {
				t.Errorf("wanted to requeue after %v but was set to %v", tt.wantRequeueAfter, result.RequeueAfter)
			}

			if tt.instance != nil {
				utilconditions.AssertControllerConditions(t, ctx, r.AROController.Client, tt.wantConditions)
			}

			utilerror.AssertErrorMessage(t, err, tt.wantErr)
		})
	}
}
