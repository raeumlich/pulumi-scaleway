// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Rdb
{
    /// <summary>
    /// Creates and manages Scaleway Database read replicas.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
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
    ///     var instance = new Scaleway.Rdb.Instance("instance", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         Tags = new[]
    ///         {
    ///             "terraform-test",
    ///             "scaleway_rdb_read_replica",
    ///             "minimal",
    ///         },
    ///     });
    /// 
    ///     var replica = new Scaleway.Rdb.ReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         DirectAccess = null,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Private network with static endpoint
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
    ///     var instance = new Scaleway.Rdb.Instance("instance", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var pn = new Scaleway.Vpc.PrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.Rdb.ReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         PrivateNetwork = new Scaleway.Rdb.Inputs.ReadReplicaPrivateNetworkArgs
    ///         {
    ///             PrivateNetworkId = pn.Id,
    ///             ServiceIp = "192.168.1.254/24",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Private network with IPAM
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
    ///     var instance = new Scaleway.Rdb.Instance("instance", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var pn = new Scaleway.Vpc.PrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.Rdb.ReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         PrivateNetwork = new Scaleway.Rdb.Inputs.ReadReplicaPrivateNetworkArgs
    ///         {
    ///             PrivateNetworkId = pn.Id,
    ///             EnableIpam = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Database Read replica can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:rdb/readReplica:ReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:rdb/readReplica:ReadReplica")]
    public partial class ReadReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Output("directAccess")]
        public Output<Outputs.ReadReplicaDirectAccess?> DirectAccess { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.ReadReplicaPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Output("sameZone")]
        public Output<bool?> SameZone { get; private set; } = null!;


        /// <summary>
        /// Create a ReadReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadReplica(string name, ReadReplicaArgs args, CustomResourceOptions? options = null)
            : base("scaleway:rdb/readReplica:ReadReplica", name, args ?? new ReadReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadReplica(string name, Input<string> id, ReadReplicaState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:rdb/readReplica:ReadReplica", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/raeumlich/pulumi-scaleway/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReadReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReadReplica Get(string name, Input<string> id, ReadReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new ReadReplica(name, id, state, options);
        }
    }

    public sealed class ReadReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.ReadReplicaDirectAccessArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.ReadReplicaPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Input("sameZone")]
        public Input<bool>? SameZone { get; set; }

        public ReadReplicaArgs()
        {
        }
        public static new ReadReplicaArgs Empty => new ReadReplicaArgs();
    }

    public sealed class ReadReplicaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.ReadReplicaDirectAccessGetArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.ReadReplicaPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Input("sameZone")]
        public Input<bool>? SameZone { get; set; }

        public ReadReplicaState()
        {
        }
        public static new ReadReplicaState Empty => new ReadReplicaState();
    }
}
