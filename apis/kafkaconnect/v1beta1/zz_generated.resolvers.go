// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"

	common "github.com/crossplane-contrib/provider-upjet-aws/config/common"
	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Connector.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *Connector) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KafkaCluster); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
						CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroups),
						Extract:       reference.ExternalName(),
						References:    mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroupRefs,
						Selector:      mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroupSelector,
						To:            reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroups")
				}
				mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
				mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroupRefs = mrsp.ResolvedReferences

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.KafkaCluster); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
						CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].Subnets),
						Extract:       reference.ExternalName(),
						References:    mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SubnetRefs,
						Selector:      mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SubnetSelector,
						To:            reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].Subnets")
				}
				mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
				mg.Spec.ForProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SubnetRefs = mrsp.ResolvedReferences

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogDelivery); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io", "v1beta1", "Group", "GroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroup),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroupRef,
						Selector:     mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroupSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroup")
				}
				mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroup = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroupRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogDelivery); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStream),
						Extract:      resource.ExtractParamPath("name", true),
						Reference:    mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStreamRef,
						Selector:     mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStreamSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStream")
				}
				mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStream = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStreamRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogDelivery); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].BucketRef,
						Selector:     mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].Bucket")
				}
				mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Plugin); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Plugin[i3].CustomPlugin); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("kafkaconnect.aws.upbound.io", "v1beta1", "CustomPlugin", "CustomPluginList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Plugin[i3].CustomPlugin[i4].Arn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.ForProvider.Plugin[i3].CustomPlugin[i4].ArnRef,
					Selector:     mg.Spec.ForProvider.Plugin[i3].CustomPlugin[i4].ArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Plugin[i3].CustomPlugin[i4].Arn")
			}
			mg.Spec.ForProvider.Plugin[i3].CustomPlugin[i4].Arn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Plugin[i3].CustomPlugin[i4].ArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceExecutionRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceExecutionRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceExecutionRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceExecutionRoleArn")
	}
	mg.Spec.ForProvider.ServiceExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceExecutionRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.WorkerConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kafkaconnect.aws.upbound.io", "v1beta1", "WorkerConfiguration", "WorkerConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkerConfiguration[i3].Arn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.WorkerConfiguration[i3].ArnRef,
				Selector:     mg.Spec.ForProvider.WorkerConfiguration[i3].ArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.WorkerConfiguration[i3].Arn")
		}
		mg.Spec.ForProvider.WorkerConfiguration[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.WorkerConfiguration[i3].ArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.KafkaCluster); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
						CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroups),
						Extract:       reference.ExternalName(),
						References:    mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroupRefs,
						Selector:      mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroupSelector,
						To:            reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroups")
				}
				mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
				mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SecurityGroupRefs = mrsp.ResolvedReferences

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.KafkaCluster); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
						CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].Subnets),
						Extract:       reference.ExternalName(),
						References:    mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SubnetRefs,
						Selector:      mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SubnetSelector,
						To:            reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].Subnets")
				}
				mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
				mg.Spec.InitProvider.KafkaCluster[i3].ApacheKafkaCluster[i4].VPC[i5].SubnetRefs = mrsp.ResolvedReferences

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LogDelivery); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io", "v1beta1", "Group", "GroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroup),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroupRef,
						Selector:     mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroupSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroup")
				}
				mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroup = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].CloudwatchLogs[i5].LogGroupRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LogDelivery); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStream),
						Extract:      resource.ExtractParamPath("name", true),
						Reference:    mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStreamRef,
						Selector:     mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStreamSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStream")
				}
				mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStream = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].Firehose[i5].DeliveryStreamRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LogDelivery); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].BucketRef,
						Selector:     mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].Bucket")
				}
				mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.LogDelivery[i3].WorkerLogDelivery[i4].S3[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Plugin); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Plugin[i3].CustomPlugin); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("kafkaconnect.aws.upbound.io", "v1beta1", "CustomPlugin", "CustomPluginList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Plugin[i3].CustomPlugin[i4].Arn),
					Extract:      common.ARNExtractor(),
					Reference:    mg.Spec.InitProvider.Plugin[i3].CustomPlugin[i4].ArnRef,
					Selector:     mg.Spec.InitProvider.Plugin[i3].CustomPlugin[i4].ArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Plugin[i3].CustomPlugin[i4].Arn")
			}
			mg.Spec.InitProvider.Plugin[i3].CustomPlugin[i4].Arn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Plugin[i3].CustomPlugin[i4].ArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceExecutionRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceExecutionRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceExecutionRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceExecutionRoleArn")
	}
	mg.Spec.InitProvider.ServiceExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceExecutionRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.WorkerConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kafkaconnect.aws.upbound.io", "v1beta1", "WorkerConfiguration", "WorkerConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkerConfiguration[i3].Arn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.WorkerConfiguration[i3].ArnRef,
				Selector:     mg.Spec.InitProvider.WorkerConfiguration[i3].ArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.WorkerConfiguration[i3].Arn")
		}
		mg.Spec.InitProvider.WorkerConfiguration[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.WorkerConfiguration[i3].ArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this CustomPlugin.
func (mg *CustomPlugin) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Location); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Location[i3].S3); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Location[i3].S3[i4].BucketArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Location[i3].S3[i4].BucketArnRef,
					Selector:     mg.Spec.ForProvider.Location[i3].S3[i4].BucketArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Location[i3].S3[i4].BucketArn")
			}
			mg.Spec.ForProvider.Location[i3].S3[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Location[i3].S3[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Location); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Location[i3].S3); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Location[i3].S3[i4].FileKey),
					Extract:      resource.ExtractParamPath("key", false),
					Reference:    mg.Spec.ForProvider.Location[i3].S3[i4].FileKeyRef,
					Selector:     mg.Spec.ForProvider.Location[i3].S3[i4].FileKeySelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Location[i3].S3[i4].FileKey")
			}
			mg.Spec.ForProvider.Location[i3].S3[i4].FileKey = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Location[i3].S3[i4].FileKeyRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Location); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Location[i3].S3); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Location[i3].S3[i4].BucketArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.Location[i3].S3[i4].BucketArnRef,
					Selector:     mg.Spec.InitProvider.Location[i3].S3[i4].BucketArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Location[i3].S3[i4].BucketArn")
			}
			mg.Spec.InitProvider.Location[i3].S3[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Location[i3].S3[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Location); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Location[i3].S3); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Location[i3].S3[i4].FileKey),
					Extract:      resource.ExtractParamPath("key", false),
					Reference:    mg.Spec.InitProvider.Location[i3].S3[i4].FileKeyRef,
					Selector:     mg.Spec.InitProvider.Location[i3].S3[i4].FileKeySelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Location[i3].S3[i4].FileKey")
			}
			mg.Spec.InitProvider.Location[i3].S3[i4].FileKey = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Location[i3].S3[i4].FileKeyRef = rsp.ResolvedReference

		}
	}

	return nil
}
