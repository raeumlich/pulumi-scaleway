// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Compute Instance security group rules. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89).
 *
 * This resource can be used to externalize rules from a `scaleway.instance.SecurityGroup` to solve circular dependency problems. When using this resource do not forget to set `externalRules = true` on the security group.
 *
 * > **Warning:** In order to guaranty rules order in a given security group only one scaleway.instance.SecurityGroupRules is allowed per security group.
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const sg01 = new scaleway.instance.SecurityGroup("sg01", {externalRules: true});
 * const sgrs01 = new scaleway.instance.SecurityGroupRules("sgrs01", {
 *     securityGroupId: sg01.id,
 *     inboundRules: [{
 *         action: "accept",
 *         port: 80,
 *         ipRange: "0.0.0.0/0",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Instance security group rules can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:instance/securityGroupRules:SecurityGroupRules web fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class SecurityGroupRules extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRulesState, opts?: pulumi.CustomResourceOptions): SecurityGroupRules {
        return new SecurityGroupRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:instance/securityGroupRules:SecurityGroupRules';

    /**
     * Returns true if the given object is an instance of SecurityGroupRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRules.__pulumiType;
    }

    /**
     * A list of inbound rule to add to the security group. (Structure is documented below.)
     */
    public readonly inboundRules!: pulumi.Output<outputs.instance.SecurityGroupRulesInboundRule[] | undefined>;
    /**
     * A list of outbound rule to add to the security group. (Structure is documented below.)
     */
    public readonly outboundRules!: pulumi.Output<outputs.instance.SecurityGroupRulesOutboundRule[] | undefined>;
    /**
     * The ID of the security group.
     */
    public readonly securityGroupId!: pulumi.Output<string>;

    /**
     * Create a SecurityGroupRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRulesArgs | SecurityGroupRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupRulesState | undefined;
            resourceInputs["inboundRules"] = state ? state.inboundRules : undefined;
            resourceInputs["outboundRules"] = state ? state.outboundRules : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
        } else {
            const args = argsOrState as SecurityGroupRulesArgs | undefined;
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            resourceInputs["inboundRules"] = args ? args.inboundRules : undefined;
            resourceInputs["outboundRules"] = args ? args.outboundRules : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroupRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRules resources.
 */
export interface SecurityGroupRulesState {
    /**
     * A list of inbound rule to add to the security group. (Structure is documented below.)
     */
    inboundRules?: pulumi.Input<pulumi.Input<inputs.instance.SecurityGroupRulesInboundRule>[]>;
    /**
     * A list of outbound rule to add to the security group. (Structure is documented below.)
     */
    outboundRules?: pulumi.Input<pulumi.Input<inputs.instance.SecurityGroupRulesOutboundRule>[]>;
    /**
     * The ID of the security group.
     */
    securityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRules resource.
 */
export interface SecurityGroupRulesArgs {
    /**
     * A list of inbound rule to add to the security group. (Structure is documented below.)
     */
    inboundRules?: pulumi.Input<pulumi.Input<inputs.instance.SecurityGroupRulesInboundRule>[]>;
    /**
     * A list of outbound rule to add to the security group. (Structure is documented below.)
     */
    outboundRules?: pulumi.Input<pulumi.Input<inputs.instance.SecurityGroupRulesOutboundRule>[]>;
    /**
     * The ID of the security group.
     */
    securityGroupId: pulumi.Input<string>;
}
