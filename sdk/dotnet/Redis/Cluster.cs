// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Redis
{
    /// <summary>
    /// Creates and manages Scaleway Redis Clusters.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/).
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Redis.Cluster("main", new()
    ///     {
    ///         Acls = new[]
    ///         {
    ///             new Scaleway.Redis.Inputs.ClusterAclArgs
    ///             {
    ///                 Description = "Allow all",
    ///                 Ip = "0.0.0.0/0",
    ///             },
    ///         },
    ///         ClusterSize = 1,
    ///         NodeType = "RED1-MICRO",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         Tags = new[]
    ///         {
    ///             "test",
    ///             "redis",
    ///         },
    ///         TlsEnabled = true,
    ///         UserName = "my_initial_user",
    ///         Version = "6.2.6",
    ///     });
    /// 
    /// });
    /// ```
    /// ### With settings
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Redis.Cluster("main", new()
    ///     {
    ///         NodeType = "RED1-MICRO",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         Settings = 
    ///         {
    ///             { "maxclients", "1000" },
    ///             { "tcp-keepalive", "120" },
    ///         },
    ///         UserName = "my_initial_user",
    ///         Version = "6.2.6",
    ///     });
    /// 
    /// });
    /// ```
    /// ### With a private network
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn = new Scaleway.Vpc.PrivateNetwork("pn");
    /// 
    ///     var main = new Scaleway.Redis.Cluster("main", new()
    ///     {
    ///         Version = "6.2.6",
    ///         NodeType = "RED1-MICRO",
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         ClusterSize = 1,
    ///         PrivateNetworks = new[]
    ///         {
    ///             new Scaleway.Redis.Inputs.ClusterPrivateNetworkArgs
    ///             {
    ///                 Id = pn.Id,
    ///                 ServiceIps = new[]
    ///                 {
    ///                     "10.12.1.1/20",
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             pn,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Redis Cluster can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:redis/cluster:Cluster main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:redis/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of acl rules, this is cluster's authorized IPs. More details on the ACL section.
        /// </summary>
        [Output("acls")]
        public Output<ImmutableArray<Outputs.ClusterAcl>> Acls { get; private set; } = null!;

        /// <summary>
        /// The PEM of the certificate used by redis, only when `tls_enabled` is true
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in the Redis Cluster.
        /// 
        /// &gt; **Important:** You cannot set `cluster_size` to 2, you either have to choose Standalone mode (1 node) or Cluster mode
        /// which is minimum 3 (1 main node + 2 secondary nodes)
        /// 
        /// &gt; **Important:** You can set a bigger `cluster_size` than you initially did, it will migrate the Redis Cluster, but
        /// keep in mind that you cannot downgrade a Redis Cluster so setting a smaller `cluster_size` will not have any effect.
        /// </summary>
        [Output("clusterSize")]
        public Output<int> ClusterSize { get; private set; } = null!;

        /// <summary>
        /// The date and time of creation of the Redis Cluster.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The name of the Redis Cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of Redis Cluster you want to create (e.g. `RED1-M`).
        /// 
        /// &gt; **Important:** Updates to `node_type` will migrate the Redis Cluster to the desired `node_type`. Keep in mind that
        /// you cannot downgrade a Redis Cluster.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// Password for the first user of the Redis Cluster.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Describes the private network you want to connect to your cluster. If not set, a public
        /// network will be provided. More details on the Private Network section
        /// 
        /// &gt; **Important:** The way to use private networks differs whether you are using redis in standalone or cluster mode.
        /// 
        /// - Standalone mode (`cluster_size` = 1) : you can attach as many private networks as you want (each must be a separate
        /// block). If you detach your only private network, your cluster won't be reachable until you define a new private or
        /// public network. You can modify your private_network and its specs, you can have both a private and public network side
        /// by side.
        /// 
        /// - Cluster mode (`cluster_size` &gt; 1) : you can define a single private network as you create your cluster, you won't be
        /// able to edit or detach it afterward, unless you create another cluster. Your `service_ips` must be listed as follows:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        [Output("privateNetworks")]
        public Output<ImmutableArray<Outputs.ClusterPrivateNetwork>> PrivateNetworks { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the Redis Cluster is
        /// associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// (Optional) Public network details. Only one of `private_network` and `public_network` may be set.
        /// &gt; The `public_network` block exports:
        /// </summary>
        [Output("publicNetwork")]
        public Output<Outputs.ClusterPublicNetwork> PublicNetwork { get; private set; } = null!;

        /// <summary>
        /// Map of settings for redis cluster. Available settings can be found by listing redis versions
        /// with scaleway API or CLI
        /// </summary>
        [Output("settings")]
        public Output<ImmutableDictionary<string, string>?> Settings { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the Redis Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether TLS is enabled or not.
        /// 
        /// &gt; The changes on `tls_enabled` will force the resource creation.
        /// </summary>
        [Output("tlsEnabled")]
        public Output<bool?> TlsEnabled { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the Redis Cluster.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Identifier for the first user of the Redis Cluster.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Redis's Cluster version (e.g. `6.2.6`).
        /// 
        /// &gt; **Important:** Updates to `version` will migrate the Redis Cluster to the desired `version`. Keep in mind that you
        /// cannot downgrade a Redis Cluster.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the
        /// Redis Cluster should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("scaleway:redis/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:redis/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/raeumlich/pulumi-scaleway/releases/",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.ClusterAclArgs>? _acls;

        /// <summary>
        /// List of acl rules, this is cluster's authorized IPs. More details on the ACL section.
        /// </summary>
        public InputList<Inputs.ClusterAclArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.ClusterAclArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// The number of nodes in the Redis Cluster.
        /// 
        /// &gt; **Important:** You cannot set `cluster_size` to 2, you either have to choose Standalone mode (1 node) or Cluster mode
        /// which is minimum 3 (1 main node + 2 secondary nodes)
        /// 
        /// &gt; **Important:** You can set a bigger `cluster_size` than you initially did, it will migrate the Redis Cluster, but
        /// keep in mind that you cannot downgrade a Redis Cluster so setting a smaller `cluster_size` will not have any effect.
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        /// <summary>
        /// The name of the Redis Cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of Redis Cluster you want to create (e.g. `RED1-M`).
        /// 
        /// &gt; **Important:** Updates to `node_type` will migrate the Redis Cluster to the desired `node_type`. Keep in mind that
        /// you cannot downgrade a Redis Cluster.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password for the first user of the Redis Cluster.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateNetworks")]
        private InputList<Inputs.ClusterPrivateNetworkArgs>? _privateNetworks;

        /// <summary>
        /// Describes the private network you want to connect to your cluster. If not set, a public
        /// network will be provided. More details on the Private Network section
        /// 
        /// &gt; **Important:** The way to use private networks differs whether you are using redis in standalone or cluster mode.
        /// 
        /// - Standalone mode (`cluster_size` = 1) : you can attach as many private networks as you want (each must be a separate
        /// block). If you detach your only private network, your cluster won't be reachable until you define a new private or
        /// public network. You can modify your private_network and its specs, you can have both a private and public network side
        /// by side.
        /// 
        /// - Cluster mode (`cluster_size` &gt; 1) : you can define a single private network as you create your cluster, you won't be
        /// able to edit or detach it afterward, unless you create another cluster. Your `service_ips` must be listed as follows:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        public InputList<Inputs.ClusterPrivateNetworkArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.ClusterPrivateNetworkArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the Redis Cluster is
        /// associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (Optional) Public network details. Only one of `private_network` and `public_network` may be set.
        /// &gt; The `public_network` block exports:
        /// </summary>
        [Input("publicNetwork")]
        public Input<Inputs.ClusterPublicNetworkArgs>? PublicNetwork { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of settings for redis cluster. Available settings can be found by listing redis versions
        /// with scaleway API or CLI
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the Redis Cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether TLS is enabled or not.
        /// 
        /// &gt; The changes on `tls_enabled` will force the resource creation.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

        /// <summary>
        /// Identifier for the first user of the Redis Cluster.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        /// <summary>
        /// Redis's Cluster version (e.g. `6.2.6`).
        /// 
        /// &gt; **Important:** Updates to `version` will migrate the Redis Cluster to the desired `version`. Keep in mind that you
        /// cannot downgrade a Redis Cluster.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the
        /// Redis Cluster should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.ClusterAclGetArgs>? _acls;

        /// <summary>
        /// List of acl rules, this is cluster's authorized IPs. More details on the ACL section.
        /// </summary>
        public InputList<Inputs.ClusterAclGetArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.ClusterAclGetArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// The PEM of the certificate used by redis, only when `tls_enabled` is true
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// The number of nodes in the Redis Cluster.
        /// 
        /// &gt; **Important:** You cannot set `cluster_size` to 2, you either have to choose Standalone mode (1 node) or Cluster mode
        /// which is minimum 3 (1 main node + 2 secondary nodes)
        /// 
        /// &gt; **Important:** You can set a bigger `cluster_size` than you initially did, it will migrate the Redis Cluster, but
        /// keep in mind that you cannot downgrade a Redis Cluster so setting a smaller `cluster_size` will not have any effect.
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        /// <summary>
        /// The date and time of creation of the Redis Cluster.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The name of the Redis Cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of Redis Cluster you want to create (e.g. `RED1-M`).
        /// 
        /// &gt; **Important:** Updates to `node_type` will migrate the Redis Cluster to the desired `node_type`. Keep in mind that
        /// you cannot downgrade a Redis Cluster.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the first user of the Redis Cluster.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateNetworks")]
        private InputList<Inputs.ClusterPrivateNetworkGetArgs>? _privateNetworks;

        /// <summary>
        /// Describes the private network you want to connect to your cluster. If not set, a public
        /// network will be provided. More details on the Private Network section
        /// 
        /// &gt; **Important:** The way to use private networks differs whether you are using redis in standalone or cluster mode.
        /// 
        /// - Standalone mode (`cluster_size` = 1) : you can attach as many private networks as you want (each must be a separate
        /// block). If you detach your only private network, your cluster won't be reachable until you define a new private or
        /// public network. You can modify your private_network and its specs, you can have both a private and public network side
        /// by side.
        /// 
        /// - Cluster mode (`cluster_size` &gt; 1) : you can define a single private network as you create your cluster, you won't be
        /// able to edit or detach it afterward, unless you create another cluster. Your `service_ips` must be listed as follows:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        public InputList<Inputs.ClusterPrivateNetworkGetArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.ClusterPrivateNetworkGetArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the Redis Cluster is
        /// associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (Optional) Public network details. Only one of `private_network` and `public_network` may be set.
        /// &gt; The `public_network` block exports:
        /// </summary>
        [Input("publicNetwork")]
        public Input<Inputs.ClusterPublicNetworkGetArgs>? PublicNetwork { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of settings for redis cluster. Available settings can be found by listing redis versions
        /// with scaleway API or CLI
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the Redis Cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether TLS is enabled or not.
        /// 
        /// &gt; The changes on `tls_enabled` will force the resource creation.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

        /// <summary>
        /// The date and time of the last update of the Redis Cluster.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Identifier for the first user of the Redis Cluster.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Redis's Cluster version (e.g. `6.2.6`).
        /// 
        /// &gt; **Important:** Updates to `version` will migrate the Redis Cluster to the desired `version`. Keep in mind that you
        /// cannot downgrade a Redis Cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// `zone`) The zone in which the
        /// Redis Cluster should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
