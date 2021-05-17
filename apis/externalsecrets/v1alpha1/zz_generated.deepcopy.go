// +build !ignore_autogenerated

/*
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

package v1alpha1

import (
	metav1 "github.com/external-secrets/external-secrets/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAuth) DeepCopyInto(out *AWSAuth) {
	*out = *in
	in.SecretRef.DeepCopyInto(&out.SecretRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAuth.
func (in *AWSAuth) DeepCopy() *AWSAuth {
	if in == nil {
		return nil
	}
	out := new(AWSAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAuthSecretRef) DeepCopyInto(out *AWSAuthSecretRef) {
	*out = *in
	in.AccessKeyID.DeepCopyInto(&out.AccessKeyID)
	in.SecretAccessKey.DeepCopyInto(&out.SecretAccessKey)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAuthSecretRef.
func (in *AWSAuthSecretRef) DeepCopy() *AWSAuthSecretRef {
	if in == nil {
		return nil
	}
	out := new(AWSAuthSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSProvider) DeepCopyInto(out *AWSProvider) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(AWSAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSProvider.
func (in *AWSProvider) DeepCopy() *AWSProvider {
	if in == nil {
		return nil
	}
	out := new(AWSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKVAuth) DeepCopyInto(out *AzureKVAuth) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(metav1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientSecret != nil {
		in, out := &in.ClientSecret, &out.ClientSecret
		*out = new(metav1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKVAuth.
func (in *AzureKVAuth) DeepCopy() *AzureKVAuth {
	if in == nil {
		return nil
	}
	out := new(AzureKVAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKVProvider) DeepCopyInto(out *AzureKVProvider) {
	*out = *in
	if in.VaultURL != nil {
		in, out := &in.VaultURL, &out.VaultURL
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.AuthSecretRef != nil {
		in, out := &in.AuthSecretRef, &out.AuthSecretRef
		*out = new(AzureKVAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKVProvider.
func (in *AzureKVProvider) DeepCopy() *AzureKVProvider {
	if in == nil {
		return nil
	}
	out := new(AzureKVProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretStore) DeepCopyInto(out *ClusterSecretStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretStore.
func (in *ClusterSecretStore) DeepCopy() *ClusterSecretStore {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSecretStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretStoreList) DeepCopyInto(out *ClusterSecretStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSecretStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretStoreList.
func (in *ClusterSecretStoreList) DeepCopy() *ClusterSecretStoreList {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSecretStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecret) DeepCopyInto(out *ExternalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecret.
func (in *ExternalSecret) DeepCopy() *ExternalSecret {
	if in == nil {
		return nil
	}
	out := new(ExternalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretData) DeepCopyInto(out *ExternalSecretData) {
	*out = *in
	out.RemoteRef = in.RemoteRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretData.
func (in *ExternalSecretData) DeepCopy() *ExternalSecretData {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretDataRemoteRef) DeepCopyInto(out *ExternalSecretDataRemoteRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretDataRemoteRef.
func (in *ExternalSecretDataRemoteRef) DeepCopy() *ExternalSecretDataRemoteRef {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretDataRemoteRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretList) DeepCopyInto(out *ExternalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretList.
func (in *ExternalSecretList) DeepCopy() *ExternalSecretList {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretSpec) DeepCopyInto(out *ExternalSecretSpec) {
	*out = *in
	out.SecretStoreRef = in.SecretStoreRef
	in.Target.DeepCopyInto(&out.Target)
	if in.RefreshInterval != nil {
		in, out := &in.RefreshInterval, &out.RefreshInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]ExternalSecretData, len(*in))
		copy(*out, *in)
	}
	if in.DataFrom != nil {
		in, out := &in.DataFrom, &out.DataFrom
		*out = make([]ExternalSecretDataRemoteRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretSpec.
func (in *ExternalSecretSpec) DeepCopy() *ExternalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretStatus) DeepCopyInto(out *ExternalSecretStatus) {
	*out = *in
	in.RefreshTime.DeepCopyInto(&out.RefreshTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ExternalSecretStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretStatus.
func (in *ExternalSecretStatus) DeepCopy() *ExternalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretStatusCondition) DeepCopyInto(out *ExternalSecretStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretStatusCondition.
func (in *ExternalSecretStatusCondition) DeepCopy() *ExternalSecretStatusCondition {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretTarget) DeepCopyInto(out *ExternalSecretTarget) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(ExternalSecretTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretTarget.
func (in *ExternalSecretTarget) DeepCopy() *ExternalSecretTarget {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretTemplate) DeepCopyInto(out *ExternalSecretTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretTemplate.
func (in *ExternalSecretTemplate) DeepCopy() *ExternalSecretTemplate {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretTemplateMetadata) DeepCopyInto(out *ExternalSecretTemplateMetadata) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretTemplateMetadata.
func (in *ExternalSecretTemplateMetadata) DeepCopy() *ExternalSecretTemplateMetadata {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretTemplateMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPSMAuth) DeepCopyInto(out *GCPSMAuth) {
	*out = *in
	in.SecretRef.DeepCopyInto(&out.SecretRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPSMAuth.
func (in *GCPSMAuth) DeepCopy() *GCPSMAuth {
	if in == nil {
		return nil
	}
	out := new(GCPSMAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPSMAuthSecretRef) DeepCopyInto(out *GCPSMAuthSecretRef) {
	*out = *in
	in.SecretAccessKey.DeepCopyInto(&out.SecretAccessKey)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPSMAuthSecretRef.
func (in *GCPSMAuthSecretRef) DeepCopy() *GCPSMAuthSecretRef {
	if in == nil {
		return nil
	}
	out := new(GCPSMAuthSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPSMProvider) DeepCopyInto(out *GCPSMProvider) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPSMProvider.
func (in *GCPSMProvider) DeepCopy() *GCPSMProvider {
	if in == nil {
		return nil
	}
	out := new(GCPSMProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStore) DeepCopyInto(out *SecretStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStore.
func (in *SecretStore) DeepCopy() *SecretStore {
	if in == nil {
		return nil
	}
	out := new(SecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreList) DeepCopyInto(out *SecretStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreList.
func (in *SecretStoreList) DeepCopy() *SecretStoreList {
	if in == nil {
		return nil
	}
	out := new(SecretStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreProvider) DeepCopyInto(out *SecretStoreProvider) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.AzureKV != nil {
		in, out := &in.AzureKV, &out.AzureKV
		*out = new(AzureKVProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(VaultProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.GCPSM != nil {
		in, out := &in.GCPSM, &out.GCPSM
		*out = new(GCPSMProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreProvider.
func (in *SecretStoreProvider) DeepCopy() *SecretStoreProvider {
	if in == nil {
		return nil
	}
	out := new(SecretStoreProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreRef) DeepCopyInto(out *SecretStoreRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreRef.
func (in *SecretStoreRef) DeepCopy() *SecretStoreRef {
	if in == nil {
		return nil
	}
	out := new(SecretStoreRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreSpec) DeepCopyInto(out *SecretStoreSpec) {
	*out = *in
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(SecretStoreProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreSpec.
func (in *SecretStoreSpec) DeepCopy() *SecretStoreSpec {
	if in == nil {
		return nil
	}
	out := new(SecretStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreStatus) DeepCopyInto(out *SecretStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]SecretStoreStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreStatus.
func (in *SecretStoreStatus) DeepCopy() *SecretStoreStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreStatusCondition) DeepCopyInto(out *SecretStoreStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreStatusCondition.
func (in *SecretStoreStatusCondition) DeepCopy() *SecretStoreStatusCondition {
	if in == nil {
		return nil
	}
	out := new(SecretStoreStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAppRole) DeepCopyInto(out *VaultAppRole) {
	*out = *in
	in.SecretRef.DeepCopyInto(&out.SecretRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAppRole.
func (in *VaultAppRole) DeepCopy() *VaultAppRole {
	if in == nil {
		return nil
	}
	out := new(VaultAppRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuth) DeepCopyInto(out *VaultAuth) {
	*out = *in
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(metav1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AppRole != nil {
		in, out := &in.AppRole, &out.AppRole
		*out = new(VaultAppRole)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(VaultKubernetesAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.Ldap != nil {
		in, out := &in.Ldap, &out.Ldap
		*out = new(VaultLdapAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.Jwt != nil {
		in, out := &in.Jwt, &out.Jwt
		*out = new(VaultJwtAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuth.
func (in *VaultAuth) DeepCopy() *VaultAuth {
	if in == nil {
		return nil
	}
	out := new(VaultAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultJwtAuth) DeepCopyInto(out *VaultJwtAuth) {
	*out = *in
	in.SecretRef.DeepCopyInto(&out.SecretRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultJwtAuth.
func (in *VaultJwtAuth) DeepCopy() *VaultJwtAuth {
	if in == nil {
		return nil
	}
	out := new(VaultJwtAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultKubernetesAuth) DeepCopyInto(out *VaultKubernetesAuth) {
	*out = *in
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(metav1.ServiceAccountSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(metav1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultKubernetesAuth.
func (in *VaultKubernetesAuth) DeepCopy() *VaultKubernetesAuth {
	if in == nil {
		return nil
	}
	out := new(VaultKubernetesAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultLdapAuth) DeepCopyInto(out *VaultLdapAuth) {
	*out = *in
	in.SecretRef.DeepCopyInto(&out.SecretRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultLdapAuth.
func (in *VaultLdapAuth) DeepCopy() *VaultLdapAuth {
	if in == nil {
		return nil
	}
	out := new(VaultLdapAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultProvider) DeepCopyInto(out *VaultProvider) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultProvider.
func (in *VaultProvider) DeepCopy() *VaultProvider {
	if in == nil {
		return nil
	}
	out := new(VaultProvider)
	in.DeepCopyInto(out)
	return out
}
