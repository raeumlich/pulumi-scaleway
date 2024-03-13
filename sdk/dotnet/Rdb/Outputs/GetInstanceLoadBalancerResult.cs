// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Rdb.Outputs
{

    [OutputType]
    public sealed class GetInstanceLoadBalancerResult
    {
        /// <summary>
        /// The endpoint ID
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The hostname of your endpoint
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The IP of your load balancer service
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The name of the RDB instance.
        /// Only one of `name` and `instance_id` should be specified.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port of your load balancer service
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private GetInstanceLoadBalancerResult(
            string endpointId,

            string hostname,

            string ip,

            string name,

            int port)
        {
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            Name = name;
            Port = port;
        }
    }
}
