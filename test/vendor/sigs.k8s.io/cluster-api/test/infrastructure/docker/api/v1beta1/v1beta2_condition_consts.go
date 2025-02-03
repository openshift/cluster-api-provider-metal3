/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"

// Conditions that will be used for the DockerCluster object in v1Beta2 API version.
const (
	// DockerClusterReadyV1Beta2Condition is true if the DockerCluster is not deleted,
	// and LoadBalancerAvailable condition is true.
	DockerClusterReadyV1Beta2Condition = clusterv1.ReadyV1Beta2Condition

	// DockerClusterLoadBalancerAvailableV1Beta2Condition documents the availability of the container that implements the cluster load balancer.
	DockerClusterLoadBalancerAvailableV1Beta2Condition clusterv1.ConditionType = "LoadBalancerAvailable"

	// DockerClusterDeletingV1Beta2Condition surfaces details about ongoing deletion of the DockerCluster.
	DockerClusterDeletingV1Beta2Condition = clusterv1.DeletingV1Beta2Condition
)

// Conditions that will be used for the DockerMachine object in v1Beta2 API version.
const (
	// DockerMachineReadyV1Beta2Condition is true if the DockerMachine is not deleted,
	// and both BootstrapExecSucceeded and ContainerProvisioned conditions are true.
	DockerMachineReadyV1Beta2Condition = clusterv1.ReadyV1Beta2Condition

	// DockerMachineContainerProvisionedV1Beta2Condition documents the status of the provisioning of the container
	// generated by a DockerMachine.
	//
	// NOTE as a difference from other providers, container provisioning and bootstrap are directly managed
	// by the DockerMachine controller (not by cloud-init).
	DockerMachineContainerProvisionedV1Beta2Condition clusterv1.ConditionType = "ContainerProvisioned"

	// DockerMachineBootstrapExecSucceededV1Beta2Condition provides an observation of the DockerMachine bootstrap process.
	// It is set based on successful execution of bootstrap commands and on the existence of
	// the /run/cluster-api/bootstrap-success.complete file.
	// The condition gets generated after ContainerProvisionedCondition is True.
	//
	// NOTE as a difference from other providers, container provisioning and bootstrap are directly managed
	// by the DockerMachine controller (not by cloud-init).
	DockerMachineBootstrapExecSucceededV1Beta2Condition clusterv1.ConditionType = "BootstrapExecSucceeded"

	// DockerMachineDeletingV1Beta2Condition surfaces details about ongoing deletion of the DockerMachine.
	DockerMachineDeletingV1Beta2Condition = clusterv1.DeletingV1Beta2Condition
)
