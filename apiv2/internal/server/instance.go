// SPDX-FileCopyrightText: (C) 2025 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"fmt"

	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	computev1 "github.com/open-edge-platform/infra-core/apiv2/v2/internal/pbapi/resources/compute/v1"
	osv1 "github.com/open-edge-platform/infra-core/apiv2/v2/internal/pbapi/resources/os/v1"
	statusv1 "github.com/open-edge-platform/infra-core/apiv2/v2/internal/pbapi/resources/status/v1"
	restv1 "github.com/open-edge-platform/infra-core/apiv2/v2/internal/pbapi/services/v1"
	inv_computev1 "github.com/open-edge-platform/infra-core/inventory/v2/pkg/api/compute/v1"
	inventory "github.com/open-edge-platform/infra-core/inventory/v2/pkg/api/inventory/v1"
	inv_osv1 "github.com/open-edge-platform/infra-core/inventory/v2/pkg/api/os/v1"
	"github.com/open-edge-platform/infra-core/inventory/v2/pkg/validator"
)

// OpenAPIInstanceToProto maps OpenAPI fields name to Proto fields name.
// The key is derived from the json property respectively of the
// structs Instance defined in edge-infra-manager-openapi-types.gen.go.
var OpenAPIInstanceToProto = map[string]string{
	"Name":   inv_computev1.InstanceResourceFieldName,
	"Kind":   inv_computev1.InstanceResourceFieldKind,
	"OsID":   inv_computev1.InstanceResourceEdgeDesiredOs,
	"HostID": inv_computev1.InstanceResourceEdgeHost,
}

var OpenAPIInstanceObjectsNames = map[string]struct{}{
	"Host":            {},
	"DesiredOs":       {},
	"CurrentOs":       {},
	"WorkloadMembers": {},
}

func toInvInstance(instance *computev1.InstanceResource) (*inv_computev1.InstanceResource, error) {
	if instance == nil {
		return &inv_computev1.InstanceResource{}, nil
	}

	invInstance := &inv_computev1.InstanceResource{
		Kind:            inv_computev1.InstanceKind(instance.GetKind()),
		Name:            instance.GetName(),
		DesiredState:    inv_computev1.InstanceState_INSTANCE_STATE_RUNNING,
		SecurityFeature: inv_osv1.SecurityFeature(instance.GetSecurityFeature()),
	}

	hostID := instance.GetHostId()
	if isSet(&hostID) {
		invInstance.Host = &inv_computev1.HostResource{
			ResourceId: hostID,
		}
	}

	osID := instance.GetOsId()
	if isSet(&osID) {
		invInstance.DesiredOs = &inv_osv1.OperatingSystemResource{
			ResourceId: osID,
		}
	}

	err := validator.ValidateMessage(invInstance)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to validate inventory resource")
		return nil, err
	}

	return invInstance, nil
}

func fromInvInstance(invInstance *inv_computev1.InstanceResource) (*computev1.InstanceResource, error) {
	if invInstance == nil {
		return &computev1.InstanceResource{}, nil
	}

	var err error
	var desiredOs *osv1.OperatingSystemResource
	var currentOs *osv1.OperatingSystemResource
	var host *computev1.HostResource
	if invInstance.GetDesiredOs() != nil {
		desiredOs = fromInvOSResource(invInstance.GetDesiredOs())
	}
	if invInstance.GetCurrentOs() != nil {
		currentOs = fromInvOSResource(invInstance.GetCurrentOs())
	}

	if invInstance.GetHost() != nil {
		host, err = fromInvHost(invInstance.GetHost(), nil)
		if err != nil {
			return nil, err
		}
	}

	workloadMembers := []*computev1.WorkloadMember{}
	for _, instWM := range invInstance.GetWorkloadMembers() {
		workloadMember, err := fromInvWorkloadMember(instWM)
		if err != nil {
			return nil, err
		}
		workloadMembers = append(workloadMembers, workloadMember)
	}
	instanceStatus := invInstance.GetInstanceStatus()
	instanceStatusIndicator := statusv1.StatusIndication(invInstance.GetInstanceStatusIndicator())
	instanceStatusTimestamp := fmt.Sprintf("%d", invInstance.GetInstanceStatusTimestamp())

	provisioningStatus := invInstance.GetProvisioningStatus()
	provisioningStatusIndicator := statusv1.StatusIndication(invInstance.GetProvisioningStatusIndicator())
	provisioningStatusTimestamp := fmt.Sprintf("%d", invInstance.GetProvisioningStatusTimestamp())

	updateStatus := invInstance.GetUpdateStatus()
	updateStatusIndicator := statusv1.StatusIndication(invInstance.GetUpdateStatusIndicator())
	updateStatusTimestamp := fmt.Sprintf("%d", invInstance.GetUpdateStatusTimestamp())

	instance := &computev1.InstanceResource{
		ResourceId:                  invInstance.GetResourceId(),
		InstanceId:                  invInstance.GetResourceId(),
		Kind:                        computev1.InstanceKind(invInstance.GetKind()),
		Name:                        invInstance.GetName(),
		DesiredState:                computev1.InstanceState(invInstance.GetDesiredState()),
		CurrentState:                computev1.InstanceState(invInstance.GetCurrentState()),
		Host:                        host,
		DesiredOs:                   desiredOs,
		CurrentOs:                   currentOs,
		SecurityFeature:             osv1.SecurityFeature(invInstance.GetSecurityFeature()),
		InstanceStatus:              instanceStatus,
		InstanceStatusIndicator:     instanceStatusIndicator,
		InstanceStatusTimestamp:     instanceStatusTimestamp,
		ProvisioningStatus:          provisioningStatus,
		ProvisioningStatusIndicator: provisioningStatusIndicator,
		ProvisioningStatusTimestamp: provisioningStatusTimestamp,
		UpdateStatus:                updateStatus,
		UpdateStatusIndicator:       updateStatusIndicator,
		UpdateStatusTimestamp:       updateStatusTimestamp,
		UpdateStatusDetail:          invInstance.GetUpdateStatusDetail(),
		WorkloadMembers:             workloadMembers,
	}

	return instance, nil
}

func (is *InventorygRPCServer) CreateInstance(
	ctx context.Context,
	req *restv1.CreateInstanceRequest,
) (*computev1.InstanceResource, error) {
	zlog.Debug().Msg("CreateInstance")

	instance := req.GetInstance()
	invInstance, err := toInvInstance(instance)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to convert to inventory instance")
		return nil, err
	}

	invRes := &inventory.Resource{
		Resource: &inventory.Resource_Instance{
			Instance: invInstance,
		},
	}

	invResp, err := is.InvClient.Create(ctx, invRes)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to create instance in inventory")
		return nil, err
	}

	instanceCreated, err := fromInvInstance(invResp.GetInstance())
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to convert from inventory instance")
		return nil, err
	}
	zlog.Debug().Msgf("Created %s", instanceCreated)
	return instanceCreated, nil
}

// Get a list of instances.
func (is *InventorygRPCServer) ListInstances(
	ctx context.Context,
	req *restv1.ListInstancesRequest,
) (*restv1.ListInstancesResponse, error) {
	zlog.Debug().Msg("ListInstances")

	filter := &inventory.ResourceFilter{
		Resource: &inventory.Resource{Resource: &inventory.Resource_Instance{Instance: &inv_computev1.InstanceResource{}}},
		Offset:   req.GetOffset(),
		Limit:    req.GetPageSize(),
		OrderBy:  req.GetOrderBy(),
		Filter:   req.GetFilter(),
	}

	invResp, err := is.InvClient.List(ctx, filter)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to list instances from inventory")
		return nil, err
	}

	invResources := invResp.GetResources()
	instances := make([]*computev1.InstanceResource, 0, len(invResources))
	for _, invRes := range invResources {
		instance, err := fromInvInstance(invRes.GetResource().GetInstance())
		if err != nil {
			zlog.InfraErr(err).Msg("Failed to convert from inventory instance")
			return nil, err
		}
		instances = append(instances, instance)
	}

	resp := &restv1.ListInstancesResponse{
		Instances:     instances,
		TotalElements: invResp.GetTotalElements(),
		HasNext:       invResp.GetHasNext(),
	}
	zlog.Debug().Msgf("Listed %s", resp)
	return resp, nil
}

// Get a specific instance.
func (is *InventorygRPCServer) GetInstance(
	ctx context.Context,
	req *restv1.GetInstanceRequest,
) (*computev1.InstanceResource, error) {
	zlog.Debug().Msg("GetInstance")

	invResp, err := is.InvClient.Get(ctx, req.GetResourceId())
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to get instance from inventory")
		return nil, err
	}

	invInstance := invResp.GetResource().GetInstance()
	instance, err := fromInvInstance(invInstance)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to convert from inventory instance")
		return nil, err
	}
	zlog.Debug().Msgf("Got %s", instance)
	return instance, nil
}

// Update a instance. (PUT).
func (is *InventorygRPCServer) UpdateInstance(
	ctx context.Context,
	req *restv1.UpdateInstanceRequest,
) (*computev1.InstanceResource, error) {
	zlog.Debug().Msg("UpdateInstance")

	instance := req.GetInstance()
	invInstance, err := toInvInstance(instance)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to convert to inventory instance")
		return nil, err
	}

	fieldmask, err := fieldmaskpb.New(invInstance, maps.Values(OpenAPIInstanceToProto)...)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to create field mask")
		return nil, err
	}

	invRes := &inventory.Resource{
		Resource: &inventory.Resource_Instance{
			Instance: invInstance,
		},
	}
	upRes, err := is.InvClient.Update(ctx, req.GetResourceId(), fieldmask, invRes)
	if err != nil {
		zlog.InfraErr(err).Msgf("failed to update inventory resource %s %s", req.GetResourceId(), invRes)
		return nil, err
	}
	invUp := upRes.GetInstance()
	invUpRes, err := fromInvInstance(invUp)
	if err != nil {
		return nil, err
	}

	zlog.Debug().Msgf("Updated %s", invUpRes)
	return invUpRes, nil
}

// Delete a instance.
func (is *InventorygRPCServer) DeleteInstance(
	ctx context.Context,
	req *restv1.DeleteInstanceRequest,
) (*restv1.DeleteInstanceResponse, error) {
	zlog.Debug().Msg("DeleteInstance")

	_, err := is.InvClient.Delete(ctx, req.GetResourceId())
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to delete instance from inventory")
		return nil, err
	}
	zlog.Debug().Msgf("Deleted %s", req.GetResourceId())
	return &restv1.DeleteInstanceResponse{}, nil
}

// Invalidate a instance.
func (is *InventorygRPCServer) InvalidateInstance(
	ctx context.Context,
	req *restv1.InvalidateInstanceRequest,
) (*restv1.InvalidateInstanceResponse, error) {
	zlog.Debug().Msg("InvalidateInstance")
	res := &inventory.Resource{
		Resource: &inventory.Resource_Instance{
			Instance: &inv_computev1.InstanceResource{
				DesiredState: inv_computev1.InstanceState_INSTANCE_STATE_UNTRUSTED,
			},
		},
	}

	fm, err := fieldmaskpb.New(res.GetInstance(), inv_computev1.InstanceResourceFieldDesiredState)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to create field mask")
		return nil, err
	}

	_, err = is.InvClient.Update(ctx, req.GetResourceId(), fm, res)
	if err != nil {
		zlog.InfraErr(err).Msg("Failed to invalidate instance in inventory")
		return nil, err
	}
	zlog.Debug().Msgf("Invalidated %s", req.GetResourceId())
	return &restv1.InvalidateInstanceResponse{}, nil
}
