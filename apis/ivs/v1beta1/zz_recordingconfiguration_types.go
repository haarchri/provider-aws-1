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

type DestinationConfigurationObservation struct {
}

type DestinationConfigurationParameters struct {

	// S3 destination configuration where recorded videos will be stored.
	// +kubebuilder:validation:Required
	S3 []S3Parameters `json:"s3" tf:"s3,omitempty"`
}

type RecordingConfigurationObservation struct {

	// ARN of the Recording Configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current state of the Recording Configuration.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RecordingConfigurationParameters struct {

	// Object containing destination configuration for where recorded video will be stored.
	// +kubebuilder:validation:Required
	DestinationConfiguration []DestinationConfigurationParameters `json:"destinationConfiguration" tf:"destination_configuration,omitempty"`

	// Recording Configuration name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	// +kubebuilder:validation:Optional
	RecordingReconnectWindowSeconds *float64 `json:"recordingReconnectWindowSeconds,omitempty" tf:"recording_reconnect_window_seconds,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
	// +kubebuilder:validation:Optional
	ThumbnailConfiguration []ThumbnailConfigurationParameters `json:"thumbnailConfiguration,omitempty" tf:"thumbnail_configuration,omitempty"`
}

type S3Observation struct {
}

type S3Parameters struct {

	// S3 bucket name where recorded videos will be stored.
	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`
}

type ThumbnailConfigurationObservation struct {
}

type ThumbnailConfigurationParameters struct {

	// Thumbnail recording mode. Valid values: DISABLED, INTERVAL.
	// +kubebuilder:validation:Optional
	RecordingMode *string `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// The targeted thumbnail-generation interval in seconds.
	// +kubebuilder:validation:Optional
	TargetIntervalSeconds *float64 `json:"targetIntervalSeconds,omitempty" tf:"target_interval_seconds,omitempty"`
}

// RecordingConfigurationSpec defines the desired state of RecordingConfiguration
type RecordingConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordingConfigurationParameters `json:"forProvider"`
}

// RecordingConfigurationStatus defines the observed state of RecordingConfiguration.
type RecordingConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecordingConfiguration is the Schema for the RecordingConfigurations API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RecordingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordingConfigurationSpec   `json:"spec"`
	Status            RecordingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordingConfigurationList contains a list of RecordingConfigurations
type RecordingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecordingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	RecordingConfiguration_Kind             = "RecordingConfiguration"
	RecordingConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecordingConfiguration_Kind}.String()
	RecordingConfiguration_KindAPIVersion   = RecordingConfiguration_Kind + "." + CRDGroupVersion.String()
	RecordingConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(RecordingConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&RecordingConfiguration{}, &RecordingConfigurationList{})
}