// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Ipam
{
    public static class GetIPs
    {
        /// <summary>
        /// Gets information about multiple IPs managed by IPAM service.
        /// 
        /// ## Examples
        /// 
        /// ### By tag
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byTag = Scaleway.Ipam.GetIPs.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "tag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### By type and resource
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vpc01 = new Scaleway.Vpc.VPC("vpc01");
        /// 
        ///     var pn01 = new Scaleway.Vpc.PrivateNetwork("pn01", new()
        ///     {
        ///         VpcId = vpc01.Id,
        ///         Ipv4Subnet = new Scaleway.Vpc.Inputs.PrivateNetworkIpv4SubnetArgs
        ///         {
        ///             Subnet = "172.16.32.0/22",
        ///         },
        ///     });
        /// 
        ///     var redis01 = new Scaleway.Redis.Cluster("redis01", new()
        ///     {
        ///         Version = "7.0.5",
        ///         NodeType = "RED1-XS",
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         ClusterSize = 3,
        ///         PrivateNetworks = new[]
        ///         {
        ///             new Scaleway.Redis.Inputs.ClusterPrivateNetworkArgs
        ///             {
        ///                 Id = pn01.Id,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var byTypeAndResource = Scaleway.Ipam.GetIPs.Invoke(new()
        ///     {
        ///         Type = "ipv4",
        ///         Resource = new Scaleway.Ipam.Inputs.GetIPsResourceInputArgs
        ///         {
        ///             Id = redis01.Id,
        ///             Type = "redis_cluster",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIPsResult> InvokeAsync(GetIPsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIPsResult>("scaleway:ipam/getIPs:getIPs", args ?? new GetIPsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple IPs managed by IPAM service.
        /// 
        /// ## Examples
        /// 
        /// ### By tag
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byTag = Scaleway.Ipam.GetIPs.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "tag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### By type and resource
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vpc01 = new Scaleway.Vpc.VPC("vpc01");
        /// 
        ///     var pn01 = new Scaleway.Vpc.PrivateNetwork("pn01", new()
        ///     {
        ///         VpcId = vpc01.Id,
        ///         Ipv4Subnet = new Scaleway.Vpc.Inputs.PrivateNetworkIpv4SubnetArgs
        ///         {
        ///             Subnet = "172.16.32.0/22",
        ///         },
        ///     });
        /// 
        ///     var redis01 = new Scaleway.Redis.Cluster("redis01", new()
        ///     {
        ///         Version = "7.0.5",
        ///         NodeType = "RED1-XS",
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         ClusterSize = 3,
        ///         PrivateNetworks = new[]
        ///         {
        ///             new Scaleway.Redis.Inputs.ClusterPrivateNetworkArgs
        ///             {
        ///                 Id = pn01.Id,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var byTypeAndResource = Scaleway.Ipam.GetIPs.Invoke(new()
        ///     {
        ///         Type = "ipv4",
        ///         Resource = new Scaleway.Ipam.Inputs.GetIPsResourceInputArgs
        ///         {
        ///             Id = redis01.Id,
        ///             Type = "redis_cluster",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIPsResult> Invoke(GetIPsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIPsResult>("scaleway:ipam/getIPs:getIPs", args ?? new GetIPsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIPsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Defines whether to filter only for IPs which are attached to a resource.
        /// </summary>
        [Input("attached")]
        public bool? Attached { get; set; }

        /// <summary>
        /// The Mac Address used as filter.
        /// </summary>
        [Input("macAddress")]
        public string? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network used as filter.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the project used as filter.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region used as filter.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Filter by resource ID, type or name.
        /// </summary>
        [Input("resource")]
        public Inputs.GetIPsResourceArgs? Resource { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags used as filter.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public string? Zonal { get; set; }

        public GetIPsArgs()
        {
        }
        public static new GetIPsArgs Empty => new GetIPsArgs();
    }

    public sealed class GetIPsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Defines whether to filter only for IPs which are attached to a resource.
        /// </summary>
        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        /// <summary>
        /// The Mac Address used as filter.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network used as filter.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the project used as filter.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region used as filter.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Filter by resource ID, type or name.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.GetIPsResourceInputArgs>? Resource { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags used as filter.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public Input<string>? Zonal { get; set; }

        public GetIPsInvokeArgs()
        {
        }
        public static new GetIPsInvokeArgs Empty => new GetIPsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIPsResult
    {
        public readonly bool? Attached;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of found IPs
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPsIpResult> Ips;
        /// <summary>
        /// The mac address.
        /// </summary>
        public readonly string? MacAddress;
        public readonly string OrganizationId;
        public readonly string? PrivateNetworkId;
        /// <summary>
        /// The ID of the project the server is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The region in which the IP is.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The list of public IPs of the server.
        /// </summary>
        public readonly Outputs.GetIPsResourceResult? Resource;
        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of resource.
        /// </summary>
        public readonly string? Type;
        public readonly string Zonal;

        [OutputConstructor]
        private GetIPsResult(
            bool? attached,

            string id,

            ImmutableArray<Outputs.GetIPsIpResult> ips,

            string? macAddress,

            string organizationId,

            string? privateNetworkId,

            string projectId,

            string region,

            Outputs.GetIPsResourceResult? resource,

            ImmutableArray<string> tags,

            string? type,

            string zonal)
        {
            Attached = attached;
            Id = id;
            Ips = ips;
            MacAddress = macAddress;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Region = region;
            Resource = resource;
            Tags = tags;
            Type = type;
            Zonal = zonal;
        }
    }
}