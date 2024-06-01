// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type StreamInitParameters struct {

	// –  The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is 0, indicating that the stream does not persist data.
	DataRetentionInHours *float64 `json:"dataRetentionInHours,omitempty" tf:"data_retention_in_hours,omitempty"`

	// The name of the device that is writing to the stream. In the current implementation, Kinesis Video Streams does not use this name.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (aws/kinesisvideo) is used.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see Media Types. If you choose to specify the MediaType, see Naming Requirements for guidelines.
	MediaType *string `json:"mediaType,omitempty" tf:"media_type,omitempty"`

	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StreamObservation struct {

	// The Amazon Resource Name (ARN) specifying the Stream (same as id)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A time stamp that indicates when the stream was created.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// –  The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is 0, indicating that the stream does not persist data.
	DataRetentionInHours *float64 `json:"dataRetentionInHours,omitempty" tf:"data_retention_in_hours,omitempty"`

	// The name of the device that is writing to the stream. In the current implementation, Kinesis Video Streams does not use this name.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The unique Stream id
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (aws/kinesisvideo) is used.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see Media Types. If you choose to specify the MediaType, see Naming Requirements for guidelines.
	MediaType *string `json:"mediaType,omitempty" tf:"media_type,omitempty"`

	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The version of the stream.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type StreamParameters struct {

	// –  The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is 0, indicating that the stream does not persist data.
	// +kubebuilder:validation:Optional
	DataRetentionInHours *float64 `json:"dataRetentionInHours,omitempty" tf:"data_retention_in_hours,omitempty"`

	// The name of the device that is writing to the stream. In the current implementation, Kinesis Video Streams does not use this name.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (aws/kinesisvideo) is used.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see Media Types. If you choose to specify the MediaType, see Naming Requirements for guidelines.
	// +kubebuilder:validation:Optional
	MediaType *string `json:"mediaType,omitempty" tf:"media_type,omitempty"`

	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StreamSpec defines the desired state of Stream
type StreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamParameters `json:"forProvider"`
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
	InitProvider StreamInitParameters `json:"initProvider,omitempty"`
}

// StreamStatus defines the observed state of Stream.
type StreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Stream is the Schema for the Streams API. Provides a AWS Kinesis Video Stream
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StreamSpec   `json:"spec"`
	Status StreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamList contains a list of Streams
type StreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stream `json:"items"`
}

// Repository type metadata.
var (
	Stream_Kind             = "Stream"
	Stream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stream_Kind}.String()
	Stream_KindAPIVersion   = Stream_Kind + "." + CRDGroupVersion.String()
	Stream_GroupVersionKind = CRDGroupVersion.WithKind(Stream_Kind)
)

func init() {
	SchemeBuilder.Register(&Stream{}, &StreamList{})
}
