// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Redis
{
    public static class GetCluster
    {
        /// <summary>
        /// Gets information about a Redis cluster. For further information check our [api documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816)
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
        ///     var myCluster = Scaleway.Redis.GetCluster.Invoke(new()
        ///     {
        ///         ClusterId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("scaleway:redis/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Redis cluster. For further information check our [api documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816)
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
        ///     var myCluster = Scaleway.Redis.GetCluster.Invoke(new()
        ///     {
        ///         ClusterId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("scaleway:redis/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Redis cluster ID.
        /// Only one of the `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// The name of the Redis cluster.
        /// Only one of the `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the Redis cluster is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Redis cluster ID.
        /// Only one of the `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The name of the Redis cluster.
        /// Only one of the `name` and `cluster_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the Redis cluster is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly ImmutableArray<Outputs.GetClusterAclResult> Acls;
        public readonly string Certificate;
        public readonly string? ClusterId;
        public readonly int ClusterSize;
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string NodeType;
        public readonly string Password;
        public readonly ImmutableArray<Outputs.GetClusterPrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        public readonly ImmutableArray<Outputs.GetClusterPublicNetworkResult> PublicNetworks;
        public readonly ImmutableDictionary<string, string> Settings;
        public readonly ImmutableArray<string> Tags;
        public readonly bool TlsEnabled;
        public readonly string UpdatedAt;
        public readonly string UserName;
        public readonly string Version;
        public readonly string? Zone;

        [OutputConstructor]
        private GetClusterResult(
            ImmutableArray<Outputs.GetClusterAclResult> acls,

            string certificate,

            string? clusterId,

            int clusterSize,

            string createdAt,

            string id,

            string? name,

            string nodeType,

            string password,

            ImmutableArray<Outputs.GetClusterPrivateNetworkResult> privateNetworks,

            string? projectId,

            ImmutableArray<Outputs.GetClusterPublicNetworkResult> publicNetworks,

            ImmutableDictionary<string, string> settings,

            ImmutableArray<string> tags,

            bool tlsEnabled,

            string updatedAt,

            string userName,

            string version,

            string? zone)
        {
            Acls = acls;
            Certificate = certificate;
            ClusterId = clusterId;
            ClusterSize = clusterSize;
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            NodeType = nodeType;
            Password = password;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            PublicNetworks = publicNetworks;
            Settings = settings;
            Tags = tags;
            TlsEnabled = tlsEnabled;
            UpdatedAt = updatedAt;
            UserName = userName;
            Version = version;
            Zone = zone;
        }
    }
}
