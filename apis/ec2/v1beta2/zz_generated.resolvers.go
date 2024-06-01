// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"

	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Route.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "ManagedPrefixList", "ManagedPrefixListList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationPrefixListID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DestinationPrefixListIDRef,
			Selector:     mg.Spec.ForProvider.DestinationPrefixListIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DestinationPrefixListID")
	}
	mg.Spec.ForProvider.DestinationPrefixListID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DestinationPrefixListIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "EgressOnlyInternetGateway", "EgressOnlyInternetGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EgressOnlyGatewayID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EgressOnlyGatewayIDRef,
			Selector:     mg.Spec.ForProvider.EgressOnlyGatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EgressOnlyGatewayID")
	}
	mg.Spec.ForProvider.EgressOnlyGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EgressOnlyGatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "InternetGateway", "InternetGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.GatewayIDRef,
			Selector:     mg.Spec.ForProvider.GatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GatewayID")
	}
	mg.Spec.ForProvider.GatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "NATGateway", "NATGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NATGatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NATGatewayIDRef,
			Selector:     mg.Spec.ForProvider.NATGatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NATGatewayID")
	}
	mg.Spec.ForProvider.NATGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NATGatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "NetworkInterface", "NetworkInterfaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterfaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterfaceIDRef,
			Selector:     mg.Spec.ForProvider.NetworkInterfaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterfaceID")
	}
	mg.Spec.ForProvider.NetworkInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkInterfaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "RouteTable", "RouteTableList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouteTableID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RouteTableIDRef,
			Selector:     mg.Spec.ForProvider.RouteTableIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouteTableID")
	}
	mg.Spec.ForProvider.RouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouteTableIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "TransitGateway", "TransitGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TransitGatewayIDRef,
			Selector:     mg.Spec.ForProvider.TransitGatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayID")
	}
	mg.Spec.ForProvider.TransitGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPCEndpoint", "VPCEndpointList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCEndpointID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCEndpointIDRef,
			Selector:     mg.Spec.ForProvider.VPCEndpointIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCEndpointID")
	}
	mg.Spec.ForProvider.VPCEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCEndpointIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPCPeeringConnection", "VPCPeeringConnectionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCPeeringConnectionID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCPeeringConnectionIDRef,
			Selector:     mg.Spec.ForProvider.VPCPeeringConnectionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCPeeringConnectionID")
	}
	mg.Spec.ForProvider.VPCPeeringConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCPeeringConnectionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "ManagedPrefixList", "ManagedPrefixListList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DestinationPrefixListID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DestinationPrefixListIDRef,
			Selector:     mg.Spec.InitProvider.DestinationPrefixListIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DestinationPrefixListID")
	}
	mg.Spec.InitProvider.DestinationPrefixListID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DestinationPrefixListIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "EgressOnlyInternetGateway", "EgressOnlyInternetGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EgressOnlyGatewayID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EgressOnlyGatewayIDRef,
			Selector:     mg.Spec.InitProvider.EgressOnlyGatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EgressOnlyGatewayID")
	}
	mg.Spec.InitProvider.EgressOnlyGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EgressOnlyGatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "InternetGateway", "InternetGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.GatewayIDRef,
			Selector:     mg.Spec.InitProvider.GatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GatewayID")
	}
	mg.Spec.InitProvider.GatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "NATGateway", "NATGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NATGatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NATGatewayIDRef,
			Selector:     mg.Spec.InitProvider.NATGatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NATGatewayID")
	}
	mg.Spec.InitProvider.NATGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NATGatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "NetworkInterface", "NetworkInterfaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkInterfaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NetworkInterfaceIDRef,
			Selector:     mg.Spec.InitProvider.NetworkInterfaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkInterfaceID")
	}
	mg.Spec.InitProvider.NetworkInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkInterfaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "RouteTable", "RouteTableList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RouteTableID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RouteTableIDRef,
			Selector:     mg.Spec.InitProvider.RouteTableIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RouteTableID")
	}
	mg.Spec.InitProvider.RouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RouteTableIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "TransitGateway", "TransitGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TransitGatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TransitGatewayIDRef,
			Selector:     mg.Spec.InitProvider.TransitGatewayIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TransitGatewayID")
	}
	mg.Spec.InitProvider.TransitGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TransitGatewayIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPCEndpoint", "VPCEndpointList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCEndpointID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VPCEndpointIDRef,
			Selector:     mg.Spec.InitProvider.VPCEndpointIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCEndpointID")
	}
	mg.Spec.InitProvider.VPCEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCEndpointIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPCPeeringConnection", "VPCPeeringConnectionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCPeeringConnectionID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VPCPeeringConnectionIDRef,
			Selector:     mg.Spec.InitProvider.VPCPeeringConnectionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCPeeringConnectionID")
	}
	mg.Spec.InitProvider.VPCPeeringConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCPeeringConnectionIDRef = rsp.ResolvedReference

	return nil
}
