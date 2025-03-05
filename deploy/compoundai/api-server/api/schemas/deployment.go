/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schemas

type DeploymentSchema struct {
	ResourceSchema
	Creator        *UserSchema               `json:"creator"`
	Cluster        *ClusterFullSchema        `json:"cluster"`
	Status         DeploymentStatus          `json:"status" enum:"unknown,non-deployed,running,unhealthy,failed,deploying"`
	URLs           []string                  `json:"urls"`
	LatestRevision *DeploymentRevisionSchema `json:"latest_revision"`
	KubeNamespace  string                    `json:"kube_namespace"`
}

type DeploymentListSchema struct {
	BaseListSchema
	Items []*DeploymentSchema `json:"items"`
}

type UpdateDeploymentSchema struct {
	Targets     []*CreateDeploymentTargetSchema `json:"targets"`
	Description *string                         `json:"description,omitempty"`
	DoNotDeploy bool                            `json:"do_not_deploy,omitempty"`
}

type CreateDeploymentSchema struct {
	Name          string `json:"name"`
	KubeNamespace string `json:"kube_namespace"`
	UpdateDeploymentSchema
}

type GetDeploymentSchema struct {
	GetClusterSchema
	DeploymentName string `uri:"deploymentName" binding:"required"`
	KubeNamespace  string `uri:"kubeNamespace" binding:"required"`
}
