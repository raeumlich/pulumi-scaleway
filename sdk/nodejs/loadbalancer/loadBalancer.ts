// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Load-Balancers.
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api).
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.loadbalancer.IP("main", {zone: "fr-par-1"});
 * const base = new scaleway.loadbalancer.LoadBalancer("base", {
 *     ipId: main.id,
 *     zone: main.zone,
 *     type: "LB-S",
 * });
 * ```
 * ### Private LB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const base = new scaleway.loadbalancer.LoadBalancer("base", {
 *     ipId: scaleway_lb_ip.main.id,
 *     zone: scaleway_lb_ip.main.zone,
 *     type: "LB-S",
 *     assignFlexibleIp: false,
 * });
 * ```
 * ### IP for Public Gateway
 * resource "scaleway_vpc_public_gateway_ip" "main" {
 * }
 *
 * ### Scaleway Private Network
 * resource scaleway_vpc_private_network main {
 * }
 *
 * ### VPC Public Gateway Network
 * resource "scaleway_vpc_public_gateway" "main" {
 *     name  = "tf-test-public-gw"
 *     type  = "VPC-GW-S"
 *     ip_id = scaleway_vpc_public_gateway_ip.main.id
 * }
 *
 * ### VPC Public Gateway Network DHCP config
 * resource "scaleway_vpc_public_gateway_dhcp" "main" {
 *     subnet = "10.0.0.0/24"
 * }
 *
 * ### VPC Gateway Network
 * resource "scaleway_vpc_gateway_network" "main" {
 *     gateway_id         = scaleway_vpc_public_gateway.main.id
 *     private_network_id = scaleway_vpc_private_network.main.id
 *     dhcp_id            = scaleway_vpc_public_gateway_dhcp.main.id
 *     cleanup_dhcp       = true
 *     enable_masquerade  = true
 * }
 *
 * ### Scaleway Instance
 * resource "scaleway_instance_server" "main" {
 *     name        = "Scaleway Terraform Provider"
 *     type        = "DEV1-S"
 *     image       = "debian_bullseye"
 *     enable_ipv6 = false
 *
 *     private_network {
 *         pn_id = scaleway_vpc_private_network.main.id
 *     }
 * }
 *
 * ### IP for LB IP
 * resource scaleway_lb_ip main {
 * }
 *
 * ### Scaleway Private Network
 * resource scaleway_vpc_private_network "main" {
 *     name = "private network with static config"
 * }
 * ## Migration
 *
 * In order to migrate to other types you can check the migration up or down via our CLI `scw lb lb-types list`.
 * this change will not recreate your Load Balancer.
 *
 * Please check our [documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-migrate-a-load-balancer) for further details
 *
 * ## IP ID
 *
 * Since v1.15.0, `ipId` is a required field. This means that now a separate `scaleway.loadbalancer.IP` is required.
 * When importing, the IP needs to be imported as well as the LB.
 * When upgrading to v1.15.0, you will need to create a new `scaleway.loadbalancer.IP` resource and import it.
 *
 * For instance, if you had the following:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.loadbalancer.LoadBalancer("main", {
 *     type: "LB-S",
 *     zone: "fr-par-1",
 * });
 * ```
 *
 * You will need to update it to:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainIP = new scaleway.loadbalancer.IP("mainIP", {});
 * const mainLoadBalancer = new scaleway.loadbalancer.LoadBalancer("mainLoadBalancer", {
 *     ipId: mainIP.id,
 *     zone: "fr-par-1",
 *     type: "LB-S",
 *     releaseIp: false,
 * });
 * ```
 *
 * ## Import
 *
 * Load-Balancer can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:loadbalancer/loadBalancer:LoadBalancer main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  Be aware that you will also need to import the `scaleway_lb_ip` resource.
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:loadbalancer/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * Defines whether to automatically assign a flexible public IP to the load-balancer.
     */
    public readonly assignFlexibleIp!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the load-balancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The load-balance public IP Address
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * The ID of the associated LB IP. See below.
     *
     * > **Important:** Updates to `ipId` will recreate the load-balancer.
     */
    public readonly ipId!: pulumi.Output<string | undefined>;
    /**
     * The name of the load-balancer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the load-balancer is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * List of private network to connect with your load balancer
     */
    public readonly privateNetworks!: pulumi.Output<outputs.loadbalancer.LoadBalancerPrivateNetwork[] | undefined>;
    /**
     * `projectId`) The ID of the project the load-balancer is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region of the resource
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * The releaseIp allow release the ip address associated with the load-balancers.
     *
     * @deprecated The resource ip will be destroyed by it's own resource. Please set this to `false`
     */
    public readonly releaseIp!: pulumi.Output<boolean | undefined>;
    /**
     * Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
     */
    public readonly sslCompatibilityLevel!: pulumi.Output<string | undefined>;
    /**
     * The tags associated with the load-balancers.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of the load-balancer. Please check the migration section to upgrade the type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * `zone`) The zone of the load-balancer.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            resourceInputs["assignFlexibleIp"] = state ? state.assignFlexibleIp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["ipId"] = state ? state.ipId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["privateNetworks"] = state ? state.privateNetworks : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["releaseIp"] = state ? state.releaseIp : undefined;
            resourceInputs["sslCompatibilityLevel"] = state ? state.sslCompatibilityLevel : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["assignFlexibleIp"] = args ? args.assignFlexibleIp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipId"] = args ? args.ipId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateNetworks"] = args ? args.privateNetworks : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["releaseIp"] = args ? args.releaseIp : undefined;
            resourceInputs["sslCompatibilityLevel"] = args ? args.sslCompatibilityLevel : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * Defines whether to automatically assign a flexible public IP to the load-balancer.
     */
    assignFlexibleIp?: pulumi.Input<boolean>;
    /**
     * The description of the load-balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The load-balance public IP Address
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The ID of the associated LB IP. See below.
     *
     * > **Important:** Updates to `ipId` will recreate the load-balancer.
     */
    ipId?: pulumi.Input<string>;
    /**
     * The name of the load-balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the load-balancer is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * List of private network to connect with your load balancer
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.loadbalancer.LoadBalancerPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the project the load-balancer is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region of the resource
     */
    region?: pulumi.Input<string>;
    /**
     * The releaseIp allow release the ip address associated with the load-balancers.
     *
     * @deprecated The resource ip will be destroyed by it's own resource. Please set this to `false`
     */
    releaseIp?: pulumi.Input<boolean>;
    /**
     * Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
     */
    sslCompatibilityLevel?: pulumi.Input<string>;
    /**
     * The tags associated with the load-balancers.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the load-balancer. Please check the migration section to upgrade the type.
     */
    type?: pulumi.Input<string>;
    /**
     * `zone`) The zone of the load-balancer.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * Defines whether to automatically assign a flexible public IP to the load-balancer.
     */
    assignFlexibleIp?: pulumi.Input<boolean>;
    /**
     * The description of the load-balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the associated LB IP. See below.
     *
     * > **Important:** Updates to `ipId` will recreate the load-balancer.
     */
    ipId?: pulumi.Input<string>;
    /**
     * The name of the load-balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * List of private network to connect with your load balancer
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.loadbalancer.LoadBalancerPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the project the load-balancer is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The releaseIp allow release the ip address associated with the load-balancers.
     *
     * @deprecated The resource ip will be destroyed by it's own resource. Please set this to `false`
     */
    releaseIp?: pulumi.Input<boolean>;
    /**
     * Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
     */
    sslCompatibilityLevel?: pulumi.Input<string>;
    /**
     * The tags associated with the load-balancers.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the load-balancer. Please check the migration section to upgrade the type.
     */
    type: pulumi.Input<string>;
    /**
     * `zone`) The zone of the load-balancer.
     */
    zone?: pulumi.Input<string>;
}