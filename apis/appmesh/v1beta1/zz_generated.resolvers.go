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

	// ResolveReferences of this GatewayRoute.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *GatewayRoute) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target); i6++ {
					for i7 := 0; i7 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService); i7++ {
						{
							m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualService", "VirtualServiceList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceName),
								Extract:      resource.ExtractParamPath("name", false),
								Reference:    mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceNameRef,
								Selector:     mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceNameSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceName")
						}
						mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceName = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceNameRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualGateway", "VirtualGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualGatewayName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.VirtualGatewayNameRef,
			Selector:     mg.Spec.ForProvider.VirtualGatewayNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualGatewayName")
	}
	mg.Spec.ForProvider.VirtualGatewayName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualGatewayNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target); i6++ {
					for i7 := 0; i7 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService); i7++ {
						{
							m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualService", "VirtualServiceList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceName),
								Extract:      resource.ExtractParamPath("name", false),
								Reference:    mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceNameRef,
								Selector:     mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceNameSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceName")
						}
						mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceName = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].Target[i6].VirtualService[i7].VirtualServiceNameRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualGateway", "VirtualGatewayList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VirtualGatewayName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.InitProvider.VirtualGatewayNameRef,
			Selector:     mg.Spec.InitProvider.VirtualGatewayNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VirtualGatewayName")
	}
	mg.Spec.InitProvider.VirtualGatewayName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VirtualGatewayNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MeshNameRef,
			Selector:     mg.Spec.ForProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MeshName")
	}
	mg.Spec.ForProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MeshNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualNode", "VirtualNodeList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode),
							Extract:      resource.ExtractParamPath("name", false),
							Reference:    mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef,
							Selector:     mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode")
					}
					mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].TCPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualNode", "VirtualNodeList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode),
							Extract:      resource.ExtractParamPath("name", false),
							Reference:    mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef,
							Selector:     mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode")
					}
					mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualRouter", "VirtualRouterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualRouterName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.VirtualRouterNameRef,
			Selector:     mg.Spec.ForProvider.VirtualRouterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualRouterName")
	}
	mg.Spec.ForProvider.VirtualRouterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualRouterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.MeshNameRef,
			Selector:     mg.Spec.InitProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MeshName")
	}
	mg.Spec.InitProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MeshNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualNode", "VirtualNodeList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode),
							Extract:      resource.ExtractParamPath("name", false),
							Reference:    mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef,
							Selector:     mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode")
					}
					mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Spec[i3].HTTPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].TCPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualNode", "VirtualNodeList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode),
							Extract:      resource.ExtractParamPath("name", false),
							Reference:    mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef,
							Selector:     mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode")
					}
					mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNode = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Spec[i3].TCPRoute[i4].Action[i5].WeightedTarget[i6].VirtualNodeRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualRouter", "VirtualRouterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VirtualRouterName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.InitProvider.VirtualRouterNameRef,
			Selector:     mg.Spec.InitProvider.VirtualRouterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VirtualRouterName")
	}
	mg.Spec.InitProvider.VirtualRouterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VirtualRouterNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualGateway.
func (mg *VirtualGateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].Listener); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate); i6++ {
					for i7 := 0; i7 < len(mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm); i7++ {
						{
							m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArn),
								Extract:      resource.ExtractParamPath("arn", true),
								Reference:    mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArnRef,
								Selector:     mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArnSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArn")
						}
						mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArnRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].Listener); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate); i6++ {
					for i7 := 0; i7 < len(mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm); i7++ {
						{
							m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArn),
								Extract:      resource.ExtractParamPath("arn", true),
								Reference:    mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArnRef,
								Selector:     mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArnSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArn")
						}
						mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.InitProvider.Spec[i3].Listener[i4].TLS[i5].Certificate[i6].Acm[i7].CertificateArnRef = rsp.ResolvedReference

					}
				}
			}
		}
	}

	return nil
}

// ResolveReferences of this VirtualNode.
func (mg *VirtualNode) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MeshNameRef,
			Selector:     mg.Spec.ForProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MeshName")
	}
	mg.Spec.ForProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MeshNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].ServiceDiscovery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("servicediscovery.aws.upbound.io", "v1beta1", "HTTPNamespace", "HTTPNamespaceList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceName),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceNameRef,
						Selector:     mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceNameSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceName")
				}
				mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceNameRef = rsp.ResolvedReference

			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.MeshNameRef,
			Selector:     mg.Spec.InitProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MeshName")
	}
	mg.Spec.InitProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MeshNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].ServiceDiscovery); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("servicediscovery.aws.upbound.io", "v1beta1", "HTTPNamespace", "HTTPNamespaceList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceName),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceNameRef,
						Selector:     mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceNameSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceName")
				}
				mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Spec[i3].ServiceDiscovery[i4].AwsCloudMap[i5].NamespaceNameRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}

// ResolveReferences of this VirtualRouter.
func (mg *VirtualRouter) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MeshNameRef,
			Selector:     mg.Spec.ForProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MeshName")
	}
	mg.Spec.ForProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MeshNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.MeshNameRef,
			Selector:     mg.Spec.InitProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MeshName")
	}
	mg.Spec.InitProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MeshNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualService.
func (mg *VirtualService) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MeshNameRef,
			Selector:     mg.Spec.ForProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MeshName")
	}
	mg.Spec.ForProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MeshNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].Provider); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualNode", "VirtualNodeList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeName),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeNameRef,
						Selector:     mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeNameSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeName")
				}
				mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeNameRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Spec[i3].Provider); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualRouter", "VirtualRouterList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterName),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterNameRef,
						Selector:     mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterNameSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterName")
				}
				mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterNameRef = rsp.ResolvedReference

			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "Mesh", "MeshList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MeshName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.MeshNameRef,
			Selector:     mg.Spec.InitProvider.MeshNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MeshName")
	}
	mg.Spec.InitProvider.MeshName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MeshNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].Provider); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualNode", "VirtualNodeList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeName),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeNameRef,
						Selector:     mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeNameSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeName")
				}
				mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualNode[i5].VirtualNodeNameRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Spec[i3].Provider); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("appmesh.aws.upbound.io", "v1beta1", "VirtualRouter", "VirtualRouterList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterName),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterNameRef,
						Selector:     mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterNameSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterName")
				}
				mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Spec[i3].Provider[i4].VirtualRouter[i5].VirtualRouterNameRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
