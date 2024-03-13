// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Instance.Inputs
{

    public sealed class SecurityGroupOutboundRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take when rule match. Possible values are: `accept` or `drop`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The ip this rule apply to. If no `ip` nor `ip_range` are specified, rule will apply to all ip. Only one of `ip` and `ip_range` should be specified.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The ip range (e.g `192.168.1.0/24`) this rule applies to. If no `ip` nor `ip_range` are specified, rule will apply to all ip. Only one of `ip` and `ip_range` should be specified.
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// The port this rule applies to. If no `port` nor `port_range` are specified, the rule will apply to all port. Only one of `port` and `port_range` should be specified.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Computed port range for this rule (e.g: 1-1024, 22-22)
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        /// <summary>
        /// The protocol this rule apply to. Possible values are: `TCP`, `UDP`, `ICMP` or `ANY`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public SecurityGroupOutboundRuleArgs()
        {
        }
        public static new SecurityGroupOutboundRuleArgs Empty => new SecurityGroupOutboundRuleArgs();
    }
}
