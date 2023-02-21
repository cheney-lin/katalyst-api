//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Katalyst Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfig) DeepCopyInto(out *CustomNodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfig.
func (in *CustomNodeConfig) DeepCopy() *CustomNodeConfig {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomNodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfigList) DeepCopyInto(out *CustomNodeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomNodeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfigList.
func (in *CustomNodeConfigList) DeepCopy() *CustomNodeConfigList {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomNodeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfigSpec) DeepCopyInto(out *CustomNodeConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfigSpec.
func (in *CustomNodeConfigSpec) DeepCopy() *CustomNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfigStatus) DeepCopyInto(out *CustomNodeConfigStatus) {
	*out = *in
	if in.KatalystCustomConfigList != nil {
		in, out := &in.KatalystCustomConfigList, &out.KatalystCustomConfigList
		*out = make([]TargetConfig, len(*in))
		copy(*out, *in)
	}
	if in.ServiceProfileConfigList != nil {
		in, out := &in.ServiceProfileConfigList, &out.ServiceProfileConfigList
		*out = make([]TargetConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfigStatus.
func (in *CustomNodeConfigStatus) DeepCopy() *CustomNodeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralSelector) DeepCopyInto(out *EphemeralSelector) {
	*out = *in
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastDuration != nil {
		in, out := &in.LastDuration, &out.LastDuration
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralSelector.
func (in *EphemeralSelector) DeepCopy() *EphemeralSelector {
	if in == nil {
		return nil
	}
	out := new(EphemeralSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionConfig) DeepCopyInto(out *EvictionConfig) {
	*out = *in
	in.EvictionPluginsConfig.DeepCopyInto(&out.EvictionPluginsConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionConfig.
func (in *EvictionConfig) DeepCopy() *EvictionConfig {
	if in == nil {
		return nil
	}
	out := new(EvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionConfiguration) DeepCopyInto(out *EvictionConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionConfiguration.
func (in *EvictionConfiguration) DeepCopy() *EvictionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EvictionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EvictionConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionConfigurationList) DeepCopyInto(out *EvictionConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EvictionConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionConfigurationList.
func (in *EvictionConfigurationList) DeepCopy() *EvictionConfigurationList {
	if in == nil {
		return nil
	}
	out := new(EvictionConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EvictionConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionConfigurationSpec) DeepCopyInto(out *EvictionConfigurationSpec) {
	*out = *in
	in.GenericConfigSpec.DeepCopyInto(&out.GenericConfigSpec)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionConfigurationSpec.
func (in *EvictionConfigurationSpec) DeepCopy() *EvictionConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(EvictionConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionPluginsConfig) DeepCopyInto(out *EvictionPluginsConfig) {
	*out = *in
	in.ReclaimedResourcesEvictionPluginConfig.DeepCopyInto(&out.ReclaimedResourcesEvictionPluginConfig)
	in.MemoryEvictionPluginConfig.DeepCopyInto(&out.MemoryEvictionPluginConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionPluginsConfig.
func (in *EvictionPluginsConfig) DeepCopy() *EvictionPluginsConfig {
	if in == nil {
		return nil
	}
	out := new(EvictionPluginsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConfigCondition) DeepCopyInto(out *GenericConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConfigCondition.
func (in *GenericConfigCondition) DeepCopy() *GenericConfigCondition {
	if in == nil {
		return nil
	}
	out := new(GenericConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConfigSpec) DeepCopyInto(out *GenericConfigSpec) {
	*out = *in
	in.EphemeralSelector.DeepCopyInto(&out.EphemeralSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConfigSpec.
func (in *GenericConfigSpec) DeepCopy() *GenericConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GenericConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConfigStatus) DeepCopyInto(out *GenericConfigStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GenericConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConfigStatus.
func (in *GenericConfigStatus) DeepCopy() *GenericConfigStatus {
	if in == nil {
		return nil
	}
	out := new(GenericConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfig) DeepCopyInto(out *KatalystCustomConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfig.
func (in *KatalystCustomConfig) DeepCopy() *KatalystCustomConfig {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KatalystCustomConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigCondition) DeepCopyInto(out *KatalystCustomConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigCondition.
func (in *KatalystCustomConfigCondition) DeepCopy() *KatalystCustomConfigCondition {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigList) DeepCopyInto(out *KatalystCustomConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KatalystCustomConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigList.
func (in *KatalystCustomConfigList) DeepCopy() *KatalystCustomConfigList {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KatalystCustomConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigSpec) DeepCopyInto(out *KatalystCustomConfigSpec) {
	*out = *in
	out.TargetType = in.TargetType
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigSpec.
func (in *KatalystCustomConfigSpec) DeepCopy() *KatalystCustomConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigStatus) DeepCopyInto(out *KatalystCustomConfigStatus) {
	*out = *in
	if in.InvalidTargetConfigList != nil {
		in, out := &in.InvalidTargetConfigList, &out.InvalidTargetConfigList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KatalystCustomConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigStatus.
func (in *KatalystCustomConfigStatus) DeepCopy() *KatalystCustomConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryEvictionPluginConfig) DeepCopyInto(out *MemoryEvictionPluginConfig) {
	*out = *in
	if in.EnableNumaLevelDetection != nil {
		in, out := &in.EnableNumaLevelDetection, &out.EnableNumaLevelDetection
		*out = new(bool)
		**out = **in
	}
	if in.EnableSystemLevelDetection != nil {
		in, out := &in.EnableSystemLevelDetection, &out.EnableSystemLevelDetection
		*out = new(bool)
		**out = **in
	}
	if in.NumaFreeBelowWatermarkTimesThreshold != nil {
		in, out := &in.NumaFreeBelowWatermarkTimesThreshold, &out.NumaFreeBelowWatermarkTimesThreshold
		*out = new(int)
		**out = **in
	}
	if in.SystemKswapdRateThreshold != nil {
		in, out := &in.SystemKswapdRateThreshold, &out.SystemKswapdRateThreshold
		*out = new(int)
		**out = **in
	}
	if in.SystemKswapdRateExceedTimesThreshold != nil {
		in, out := &in.SystemKswapdRateExceedTimesThreshold, &out.SystemKswapdRateExceedTimesThreshold
		*out = new(int)
		**out = **in
	}
	if in.NumaEvictionRankingMetrics != nil {
		in, out := &in.NumaEvictionRankingMetrics, &out.NumaEvictionRankingMetrics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SystemEvictionRankingMetrics != nil {
		in, out := &in.SystemEvictionRankingMetrics, &out.SystemEvictionRankingMetrics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryEvictionPluginConfig.
func (in *MemoryEvictionPluginConfig) DeepCopy() *MemoryEvictionPluginConfig {
	if in == nil {
		return nil
	}
	out := new(MemoryEvictionPluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimedResourcesEvictionPluginConfig) DeepCopyInto(out *ReclaimedResourcesEvictionPluginConfig) {
	*out = *in
	if in.EvictionThreshold != nil {
		in, out := &in.EvictionThreshold, &out.EvictionThreshold
		*out = make(map[corev1.ResourceName]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimedResourcesEvictionPluginConfig.
func (in *ReclaimedResourcesEvictionPluginConfig) DeepCopy() *ReclaimedResourcesEvictionPluginConfig {
	if in == nil {
		return nil
	}
	out := new(ReclaimedResourcesEvictionPluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetConfig) DeepCopyInto(out *TargetConfig) {
	*out = *in
	out.ConfigType = in.ConfigType
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetConfig.
func (in *TargetConfig) DeepCopy() *TargetConfig {
	if in == nil {
		return nil
	}
	out := new(TargetConfig)
	in.DeepCopyInto(out)
	return out
}
