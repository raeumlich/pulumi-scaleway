// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "scaleway:loadbalancer/aCL:ACL":
		r = &ACL{}
	case "scaleway:loadbalancer/backend:Backend":
		r = &Backend{}
	case "scaleway:loadbalancer/certficate:Certficate":
		r = &Certficate{}
	case "scaleway:loadbalancer/frontend:Frontend":
		r = &Frontend{}
	case "scaleway:loadbalancer/iP:IP":
		r = &IP{}
	case "scaleway:loadbalancer/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "scaleway:loadbalancer/route:Route":
		r = &Route{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/aCL",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/backend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/certficate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/frontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/iP",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"loadbalancer/route",
		&module{version},
	)
}
