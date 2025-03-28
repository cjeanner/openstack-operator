//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
	ovn_operatorapiv1beta1 "github.com/openstack-k8s-operators/ovn-operator/api/v1beta1"
	apiv1beta1 "github.com/rabbitmq/cluster-operator/api/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSection) DeepCopyInto(out *CinderSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSection.
func (in *CinderSection) DeepCopy() *CinderSection {
	if in == nil {
		return nil
	}
	out := new(CinderSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlanceSection) DeepCopyInto(out *GlanceSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlanceSection.
func (in *GlanceSection) DeepCopy() *GlanceSection {
	if in == nil {
		return nil
	}
	out := new(GlanceSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeystoneSection) DeepCopyInto(out *KeystoneSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeystoneSection.
func (in *KeystoneSection) DeepCopy() *KeystoneSection {
	if in == nil {
		return nil
	}
	out := new(KeystoneSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariadbSection) DeepCopyInto(out *MariadbSection) {
	*out = *in
	out.Template = in.Template
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariadbSection.
func (in *MariadbSection) DeepCopy() *MariadbSection {
	if in == nil {
		return nil
	}
	out := new(MariadbSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSection) DeepCopyInto(out *NeutronSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSection.
func (in *NeutronSection) DeepCopy() *NeutronSection {
	if in == nil {
		return nil
	}
	out := new(NeutronSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NovaSection) DeepCopyInto(out *NovaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NovaSection.
func (in *NovaSection) DeepCopy() *NovaSection {
	if in == nil {
		return nil
	}
	out := new(NovaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlane) DeepCopyInto(out *OpenStackControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlane.
func (in *OpenStackControlPlane) DeepCopy() *OpenStackControlPlane {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneList) DeepCopyInto(out *OpenStackControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneList.
func (in *OpenStackControlPlaneList) DeepCopy() *OpenStackControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneSpec) DeepCopyInto(out *OpenStackControlPlaneSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Keystone.DeepCopyInto(&out.Keystone)
	in.Placement.DeepCopyInto(&out.Placement)
	in.Glance.DeepCopyInto(&out.Glance)
	in.Cinder.DeepCopyInto(&out.Cinder)
	out.Mariadb = in.Mariadb
	in.Rabbitmq.DeepCopyInto(&out.Rabbitmq)
	in.Ovn.DeepCopyInto(&out.Ovn)
	in.Ovs.DeepCopyInto(&out.Ovs)
	in.Neutron.DeepCopyInto(&out.Neutron)
	in.Nova.DeepCopyInto(&out.Nova)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]OpenStackExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneSpec.
func (in *OpenStackControlPlaneSpec) DeepCopy() *OpenStackControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneStatus) DeepCopyInto(out *OpenStackControlPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneStatus.
func (in *OpenStackControlPlaneStatus) DeepCopy() *OpenStackControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackExtraVolMounts) DeepCopyInto(out *OpenStackExtraVolMounts) {
	*out = *in
	if in.VolMounts != nil {
		in, out := &in.VolMounts, &out.VolMounts
		*out = make([]storage.VolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackExtraVolMounts.
func (in *OpenStackExtraVolMounts) DeepCopy() *OpenStackExtraVolMounts {
	if in == nil {
		return nil
	}
	out := new(OpenStackExtraVolMounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvnResources) DeepCopyInto(out *OvnResources) {
	*out = *in
	if in.OVNDBCluster != nil {
		in, out := &in.OVNDBCluster, &out.OVNDBCluster
		*out = make(map[string]ovn_operatorapiv1beta1.OVNDBClusterSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.OVNNorthd.DeepCopyInto(&out.OVNNorthd)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvnResources.
func (in *OvnResources) DeepCopy() *OvnResources {
	if in == nil {
		return nil
	}
	out := new(OvnResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvnSection) DeepCopyInto(out *OvnSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvnSection.
func (in *OvnSection) DeepCopy() *OvnSection {
	if in == nil {
		return nil
	}
	out := new(OvnSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvsSection) DeepCopyInto(out *OvsSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvsSection.
func (in *OvsSection) DeepCopy() *OvsSection {
	if in == nil {
		return nil
	}
	out := new(OvsSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementSection) DeepCopyInto(out *PlacementSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementSection.
func (in *PlacementSection) DeepCopy() *PlacementSection {
	if in == nil {
		return nil
	}
	out := new(PlacementSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqSection) DeepCopyInto(out *RabbitmqSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]apiv1beta1.RabbitmqClusterSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqSection.
func (in *RabbitmqSection) DeepCopy() *RabbitmqSection {
	if in == nil {
		return nil
	}
	out := new(RabbitmqSection)
	in.DeepCopyInto(out)
	return out
}
