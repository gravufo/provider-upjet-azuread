// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RoleAssignmentInitParameters struct {

	// The ID of the app role to be assigned, or the default role ID 00000000-0000-0000-0000-000000000000. Changing this forces a new resource to be created.
	// The ID of the app role to be assigned
	AppRoleID *string `json:"appRoleId,omitempty" tf:"app_role_id,omitempty"`
}

type RoleAssignmentObservation struct {

	// The ID of the app role to be assigned, or the default role ID 00000000-0000-0000-0000-000000000000. Changing this forces a new resource to be created.
	// The ID of the app role to be assigned
	AppRoleID *string `json:"appRoleId,omitempty" tf:"app_role_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of the principal to which the app role is assigned.
	// The display name of the principal to which the app role is assigned
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty" tf:"principal_display_name,omitempty"`

	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the user, group or service principal to be assigned this app role
	PrincipalObjectID *string `json:"principalObjectId,omitempty" tf:"principal_object_id,omitempty"`

	// The object type of the principal to which the app role is assigned.
	// The object type of the principal to which the app role is assigned
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The display name of the application representing the resource.
	// The display name of the application representing the resource
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty" tf:"resource_display_name,omitempty"`

	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	// The object ID of the service principal representing the resource
	ResourceObjectID *string `json:"resourceObjectId,omitempty" tf:"resource_object_id,omitempty"`
}

type RoleAssignmentParameters struct {

	// The ID of the app role to be assigned, or the default role ID 00000000-0000-0000-0000-000000000000. Changing this forces a new resource to be created.
	// The ID of the app role to be assigned
	// +kubebuilder:validation:Optional
	AppRoleID *string `json:"appRoleId,omitempty" tf:"app_role_id,omitempty"`

	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the user, group or service principal to be assigned this app role
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal
	// +kubebuilder:validation:Optional
	PrincipalObjectID *string `json:"principalObjectId,omitempty" tf:"principal_object_id,omitempty"`

	// Reference to a Principal in serviceprincipals to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDRef *v1.Reference `json:"principalObjectIdRef,omitempty" tf:"-"`

	// Selector for a Principal in serviceprincipals to populate principalObjectId.
	// +kubebuilder:validation:Optional
	PrincipalObjectIDSelector *v1.Selector `json:"principalObjectIdSelector,omitempty" tf:"-"`

	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	// The object ID of the service principal representing the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal
	// +kubebuilder:validation:Optional
	ResourceObjectID *string `json:"resourceObjectId,omitempty" tf:"resource_object_id,omitempty"`

	// Reference to a Principal in serviceprincipals to populate resourceObjectId.
	// +kubebuilder:validation:Optional
	ResourceObjectIDRef *v1.Reference `json:"resourceObjectIdRef,omitempty" tf:"-"`

	// Selector for a Principal in serviceprincipals to populate resourceObjectId.
	// +kubebuilder:validation:Optional
	ResourceObjectIDSelector *v1.Selector `json:"resourceObjectIdSelector,omitempty" tf:"-"`
}

// RoleAssignmentSpec defines the desired state of RoleAssignment
type RoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssignmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RoleAssignmentInitParameters `json:"initProvider,omitempty"`
}

// RoleAssignmentStatus defines the observed state of RoleAssignment.
type RoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignment is the Schema for the RoleAssignments API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.appRoleId) || (has(self.initProvider) && has(self.initProvider.appRoleId))",message="spec.forProvider.appRoleId is a required parameter"
	Spec   RoleAssignmentSpec   `json:"spec"`
	Status RoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignmentList contains a list of RoleAssignments
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	RoleAssignment_Kind             = "RoleAssignment"
	RoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAssignment_Kind}.String()
	RoleAssignment_KindAPIVersion   = RoleAssignment_Kind + "." + CRDGroupVersion.String()
	RoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(RoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
