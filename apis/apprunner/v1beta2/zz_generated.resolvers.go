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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Service.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.NetworkConfiguration != nil {
		if mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration != nil {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "VPCConnector", "VPCConnectorList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArnRef,
					Selector:     mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArn")
			}
			mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.ObservabilityConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta2", "ObservabilityConfiguration", "ObservabilityConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ObservabilityConfiguration.ObservabilityConfigurationArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ObservabilityConfiguration.ObservabilityConfigurationArnRef,
				Selector:     mg.Spec.ForProvider.ObservabilityConfiguration.ObservabilityConfigurationArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ObservabilityConfiguration.ObservabilityConfigurationArn")
		}
		mg.Spec.ForProvider.ObservabilityConfiguration.ObservabilityConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ObservabilityConfiguration.ObservabilityConfigurationArnRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.SourceConfiguration != nil {
		if mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration != nil {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "Connection", "ConnectionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArnRef,
					Selector:     mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArn")
			}
			mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.NetworkConfiguration != nil {
		if mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration != nil {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "VPCConnector", "VPCConnectorList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArnRef,
					Selector:     mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArn")
			}
			mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.NetworkConfiguration.EgressConfiguration.VPCConnectorArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.ObservabilityConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta2", "ObservabilityConfiguration", "ObservabilityConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ObservabilityConfiguration.ObservabilityConfigurationArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.ObservabilityConfiguration.ObservabilityConfigurationArnRef,
				Selector:     mg.Spec.InitProvider.ObservabilityConfiguration.ObservabilityConfigurationArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ObservabilityConfiguration.ObservabilityConfigurationArn")
		}
		mg.Spec.InitProvider.ObservabilityConfiguration.ObservabilityConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ObservabilityConfiguration.ObservabilityConfigurationArnRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.SourceConfiguration != nil {
		if mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration != nil {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "Connection", "ConnectionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArnRef,
					Selector:     mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArn")
			}
			mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.SourceConfiguration.AuthenticationConfiguration.ConnectionArnRef = rsp.ResolvedReference

		}
	}

	return nil
}