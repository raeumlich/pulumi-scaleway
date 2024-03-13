// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer
{
    public static class GetLoadBalancer
    {
        /// <summary>
        /// Gets information about a Load Balancer.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.Loadbalancer.GetLoadBalancer.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.Loadbalancer.GetLoadBalancer.Invoke(new()
        ///     {
        ///         LbId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("scaleway:loadbalancer/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Load Balancer.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.Loadbalancer.GetLoadBalancer.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.Loadbalancer.GetLoadBalancer.Invoke(new()
        ///     {
        ///         LbId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLoadBalancerResult> Invoke(GetLoadBalancerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadBalancerResult>("scaleway:loadbalancer/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancerArgs : global::Pulumi.InvokeArgs
    {
        [Input("lbId")]
        public string? LbId { get; set; }

        /// <summary>
        /// The load balancer name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the LB is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("releaseIp")]
        public bool? ReleaseIp { get; set; }

        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the LB exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetLoadBalancerArgs()
        {
        }
        public static new GetLoadBalancerArgs Empty => new GetLoadBalancerArgs();
    }

    public sealed class GetLoadBalancerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The load balancer name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the LB is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("releaseIp")]
        public Input<bool>? ReleaseIp { get; set; }

        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the LB exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetLoadBalancerInvokeArgs()
        {
        }
        public static new GetLoadBalancerInvokeArgs Empty => new GetLoadBalancerInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        public readonly bool AssignFlexibleIp;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The load-balancer public IP Address.
        /// </summary>
        public readonly string IpAddress;
        public readonly string IpId;
        public readonly string? LbId;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly ImmutableArray<Outputs.GetLoadBalancerPrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        public readonly string Region;
        public readonly bool? ReleaseIp;
        public readonly string SslCompatibilityLevel;
        /// <summary>
        /// The tags associated with the load-balancer.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of the load-balancer.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the LB exists.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private GetLoadBalancerResult(
            bool assignFlexibleIp,

            string description,

            string id,

            string ipAddress,

            string ipId,

            string? lbId,

            string? name,

            string organizationId,

            ImmutableArray<Outputs.GetLoadBalancerPrivateNetworkResult> privateNetworks,

            string? projectId,

            string region,

            bool? releaseIp,

            string sslCompatibilityLevel,

            ImmutableArray<string> tags,

            string type,

            string? zone)
        {
            AssignFlexibleIp = assignFlexibleIp;
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            IpId = ipId;
            LbId = lbId;
            Name = name;
            OrganizationId = organizationId;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            Region = region;
            ReleaseIp = releaseIp;
            SslCompatibilityLevel = sslCompatibilityLevel;
            Tags = tags;
            Type = type;
            Zone = zone;
        }
    }
}
