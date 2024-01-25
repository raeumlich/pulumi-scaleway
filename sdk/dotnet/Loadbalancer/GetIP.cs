// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer
{
    public static class GetIP
    {
        /// <summary>
        /// Gets information about a Load Balancer IP.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myIp = Scaleway.Loadbalancer.GetIP.Invoke(new()
        ///     {
        ///         IpId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIPResult> InvokeAsync(GetIPArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIPResult>("scaleway:loadbalancer/getIP:getIP", args ?? new GetIPArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Load Balancer IP.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myIp = Scaleway.Loadbalancer.GetIP.Invoke(new()
        ///     {
        ///         IpId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIPResult> Invoke(GetIPInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIPResult>("scaleway:loadbalancer/getIP:getIP", args ?? new GetIPInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIPArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        /// <summary>
        /// The IP ID.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipId")]
        public string? IpId { get; set; }

        /// <summary>
        /// The ID of the project the LB IP associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetIPArgs()
        {
        }
        public static new GetIPArgs Empty => new GetIPArgs();
    }

    public sealed class GetIPInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP ID.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The ID of the project the LB IP associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetIPInvokeArgs()
        {
        }
        public static new GetIPInvokeArgs Empty => new GetIPInvokeArgs();
    }


    [OutputType]
    public sealed class GetIPResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IpAddress;
        public readonly string? IpId;
        /// <summary>
        /// The associated load-balancer ID if any
        /// </summary>
        public readonly string LbId;
        /// <summary>
        /// (Defaults to provider `organization_id`) The ID of the organization the LB IP is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string Region;
        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        public readonly string Reverse;
        public readonly string Zone;

        [OutputConstructor]
        private GetIPResult(
            string id,

            string? ipAddress,

            string? ipId,

            string lbId,

            string organizationId,

            string projectId,

            string region,

            string reverse,

            string zone)
        {
            Id = id;
            IpAddress = ipAddress;
            IpId = ipId;
            LbId = lbId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Reverse = reverse;
            Zone = zone;
        }
    }
}