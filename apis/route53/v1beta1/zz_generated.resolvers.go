// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"

	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this HealthCheck.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *HealthCheck) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatch.aws.upbound.io", "v1beta1", "MetricAlarm", "MetricAlarmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudwatchAlarmName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CloudwatchAlarmNameRef,
			Selector:     mg.Spec.ForProvider.CloudwatchAlarmNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudwatchAlarmName")
	}
	mg.Spec.ForProvider.CloudwatchAlarmName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudwatchAlarmNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatch.aws.upbound.io", "v1beta1", "MetricAlarm", "MetricAlarmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CloudwatchAlarmName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CloudwatchAlarmNameRef,
			Selector:     mg.Spec.InitProvider.CloudwatchAlarmNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CloudwatchAlarmName")
	}
	mg.Spec.InitProvider.CloudwatchAlarmName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudwatchAlarmNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedZoneDNSSEC.
func (mg *HostedZoneDNSSEC) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostedZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.HostedZoneIDRef,
			Selector:     mg.Spec.ForProvider.HostedZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostedZoneID")
	}
	mg.Spec.ForProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostedZoneIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HostedZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.HostedZoneIDRef,
			Selector:     mg.Spec.InitProvider.HostedZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HostedZoneID")
	}
	mg.Spec.InitProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HostedZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Record.
func (mg *Record) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "HealthCheck", "HealthCheckList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HealthCheckID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.HealthCheckIDRef,
			Selector:     mg.Spec.ForProvider.HealthCheckIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HealthCheckID")
	}
	mg.Spec.ForProvider.HealthCheckID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HealthCheckIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ZoneIDRef,
			Selector:     mg.Spec.ForProvider.ZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "HealthCheck", "HealthCheckList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HealthCheckID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.HealthCheckIDRef,
			Selector:     mg.Spec.InitProvider.HealthCheckIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HealthCheckID")
	}
	mg.Spec.InitProvider.HealthCheckID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HealthCheckIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ZoneIDRef,
			Selector:     mg.Spec.InitProvider.ZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResolverConfig.
func (mg *ResolverConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ResourceIDRef,
			Selector:     mg.Spec.InitProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceID")
	}
	mg.Spec.InitProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TrafficPolicyInstance.
func (mg *TrafficPolicyInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostedZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.HostedZoneIDRef,
			Selector:     mg.Spec.ForProvider.HostedZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostedZoneID")
	}
	mg.Spec.ForProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostedZoneIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "TrafficPolicy", "TrafficPolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TrafficPolicyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TrafficPolicyIDRef,
			Selector:     mg.Spec.ForProvider.TrafficPolicyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TrafficPolicyID")
	}
	mg.Spec.ForProvider.TrafficPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TrafficPolicyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HostedZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.HostedZoneIDRef,
			Selector:     mg.Spec.InitProvider.HostedZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HostedZoneID")
	}
	mg.Spec.InitProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HostedZoneIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "TrafficPolicy", "TrafficPolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TrafficPolicyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TrafficPolicyIDRef,
			Selector:     mg.Spec.InitProvider.TrafficPolicyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TrafficPolicyID")
	}
	mg.Spec.InitProvider.TrafficPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TrafficPolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCAssociationAuthorization.
func (mg *VPCAssociationAuthorization) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCIDRef,
			Selector:     mg.Spec.ForProvider.VPCIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ZoneIDRef,
			Selector:     mg.Spec.ForProvider.ZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VPCIDRef,
			Selector:     mg.Spec.InitProvider.VPCIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ZoneIDRef,
			Selector:     mg.Spec.InitProvider.ZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Zone.
func (mg *Zone) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "DelegationSet", "DelegationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DelegationSetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DelegationSetIDRef,
			Selector:     mg.Spec.ForProvider.DelegationSetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DelegationSetID")
	}
	mg.Spec.ForProvider.DelegationSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DelegationSetIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPC); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPC[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VPC[i3].VPCIDRef,
				Selector:     mg.Spec.ForProvider.VPC[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPC[i3].VPCID")
		}
		mg.Spec.ForProvider.VPC[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPC[i3].VPCIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "DelegationSet", "DelegationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DelegationSetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DelegationSetIDRef,
			Selector:     mg.Spec.InitProvider.DelegationSetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DelegationSetID")
	}
	mg.Spec.InitProvider.DelegationSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DelegationSetIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPC); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPC[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VPC[i3].VPCIDRef,
				Selector:     mg.Spec.InitProvider.VPC[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPC[i3].VPCID")
		}
		mg.Spec.InitProvider.VPC[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VPC[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ZoneAssociation.
func (mg *ZoneAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCIDRef,
			Selector:     mg.Spec.ForProvider.VPCIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
			Extract:      resource.ExtractParamPath("zone_id", true),
			Reference:    mg.Spec.ForProvider.ZoneIDRef,
			Selector:     mg.Spec.ForProvider.ZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VPCIDRef,
			Selector:     mg.Spec.InitProvider.VPCIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53.aws.upbound.io", "v1beta1", "Zone", "ZoneList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
			Extract:      resource.ExtractParamPath("zone_id", true),
			Reference:    mg.Spec.InitProvider.ZoneIDRef,
			Selector:     mg.Spec.InitProvider.ZoneIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}
