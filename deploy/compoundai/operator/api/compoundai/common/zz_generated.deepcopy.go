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

//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package common

import (
	"k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTargetHPAConf) DeepCopyInto(out *DeploymentTargetHPAConf) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(int32)
		**out = **in
	}
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(int32)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
	if in.QPS != nil {
		in, out := &in.QPS, &out.QPS
		*out = new(int64)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTargetHPAConf.
func (in *DeploymentTargetHPAConf) DeepCopy() *DeploymentTargetHPAConf {
	if in == nil {
		return nil
	}
	out := new(DeploymentTargetHPAConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraPodMetadata) DeepCopyInto(out *ExtraPodMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraPodMetadata.
func (in *ExtraPodMetadata) DeepCopy() *ExtraPodMetadata {
	if in == nil {
		return nil
	}
	out := new(ExtraPodMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraPodSpec) DeepCopyInto(out *ExtraPodSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraPodSpec.
func (in *ExtraPodSpec) DeepCopy() *ExtraPodSpec {
	if in == nil {
		return nil
	}
	out := new(ExtraPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelItemSchema) DeepCopyInto(out *LabelItemSchema) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelItemSchema.
func (in *LabelItemSchema) DeepCopy() *LabelItemSchema {
	if in == nil {
		return nil
	}
	out := new(LabelItemSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorExporterMountSpec) DeepCopyInto(out *MonitorExporterMountSpec) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorExporterMountSpec.
func (in *MonitorExporterMountSpec) DeepCopy() *MonitorExporterMountSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorExporterMountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorExporterSpec) DeepCopyInto(out *MonitorExporterSpec) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StructureOptions != nil {
		in, out := &in.StructureOptions, &out.StructureOptions
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]MonitorExporterMountSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorExporterSpec.
func (in *MonitorExporterSpec) DeepCopy() *MonitorExporterSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceItem) DeepCopyInto(out *ResourceItem) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceItem.
func (in *ResourceItem) DeepCopy() *ResourceItem {
	if in == nil {
		return nil
	}
	out := new(ResourceItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceItem)
		(*in).DeepCopyInto(*out)
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceItem)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}
