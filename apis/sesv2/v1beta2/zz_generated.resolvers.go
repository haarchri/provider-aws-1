// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ConfigurationSetEventDestination) ResolveReferences( // ResolveReferences of this ConfigurationSetEventDestination.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sesv2.aws.upbound.io", "v1beta2", "ConfigurationSet", "ConfigurationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ConfigurationSetNameRef,
			Selector:     mg.Spec.ForProvider.ConfigurationSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationSetName")
	}
	mg.Spec.ForProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationSetNameRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.EventDestination != nil {
		if mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArn),
					Extract:      resource.ExtractParamPath("arn", false),
					Reference:    mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArnRef,
					Selector:     mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArn")
			}
			mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.EventDestination != nil {
		if mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArnRef,
					Selector:     mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArn")
			}
			mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.EventDestination != nil {
		if mg.Spec.ForProvider.EventDestination.PinpointDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("pinpoint.aws.upbound.io", "v1beta2", "App", "AppList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventDestination.PinpointDestination.ApplicationArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.EventDestination.PinpointDestination.ApplicationArnRef,
					Selector:     mg.Spec.ForProvider.EventDestination.PinpointDestination.ApplicationArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventDestination.PinpointDestination.ApplicationArn")
			}
			mg.Spec.ForProvider.EventDestination.PinpointDestination.ApplicationArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventDestination.PinpointDestination.ApplicationArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.EventDestination != nil {
		if mg.Spec.ForProvider.EventDestination.SnsDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventDestination.SnsDestination.TopicArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.EventDestination.SnsDestination.TopicArnRef,
					Selector:     mg.Spec.ForProvider.EventDestination.SnsDestination.TopicArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventDestination.SnsDestination.TopicArn")
			}
			mg.Spec.ForProvider.EventDestination.SnsDestination.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventDestination.SnsDestination.TopicArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("sesv2.aws.upbound.io", "v1beta2", "ConfigurationSet", "ConfigurationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ConfigurationSetNameRef,
			Selector:     mg.Spec.InitProvider.ConfigurationSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationSetName")
	}
	mg.Spec.InitProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationSetNameRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.EventDestination != nil {
		if mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArn),
					Extract:      resource.ExtractParamPath("arn", false),
					Reference:    mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArnRef,
					Selector:     mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArn")
			}
			mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.DeliveryStreamArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EventDestination != nil {
		if mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArnRef,
					Selector:     mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArn")
			}
			mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventDestination.KinesisFirehoseDestination.IAMRoleArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EventDestination != nil {
		if mg.Spec.InitProvider.EventDestination.PinpointDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("pinpoint.aws.upbound.io", "v1beta2", "App", "AppList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventDestination.PinpointDestination.ApplicationArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.EventDestination.PinpointDestination.ApplicationArnRef,
					Selector:     mg.Spec.InitProvider.EventDestination.PinpointDestination.ApplicationArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventDestination.PinpointDestination.ApplicationArn")
			}
			mg.Spec.InitProvider.EventDestination.PinpointDestination.ApplicationArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventDestination.PinpointDestination.ApplicationArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.EventDestination != nil {
		if mg.Spec.InitProvider.EventDestination.SnsDestination != nil {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventDestination.SnsDestination.TopicArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.EventDestination.SnsDestination.TopicArnRef,
					Selector:     mg.Spec.InitProvider.EventDestination.SnsDestination.TopicArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventDestination.SnsDestination.TopicArn")
			}
			mg.Spec.InitProvider.EventDestination.SnsDestination.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventDestination.SnsDestination.TopicArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this EmailIdentity.
func (mg *EmailIdentity) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sesv2.aws.upbound.io", "v1beta2", "ConfigurationSet", "ConfigurationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ConfigurationSetNameRef,
			Selector:     mg.Spec.ForProvider.ConfigurationSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationSetName")
	}
	mg.Spec.ForProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationSetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sesv2.aws.upbound.io", "v1beta2", "ConfigurationSet", "ConfigurationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ConfigurationSetNameRef,
			Selector:     mg.Spec.InitProvider.ConfigurationSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationSetName")
	}
	mg.Spec.InitProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationSetNameRef = rsp.ResolvedReference

	return nil
}
