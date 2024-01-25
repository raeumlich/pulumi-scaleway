// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Instance
{
    /// <summary>
    /// Creates and manages Scaleway Compute Snapshots.
    /// For more information,
    /// see [the documentation](https://developers.scaleway.com/en/products/instance/api/#snapshots-756fae).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Instance.Snapshot("main", new()
    ///     {
    ///         VolumeId = "11111111-1111-1111-1111-111111111111",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Example with Unified type snapshot
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainVolume = new Scaleway.Instance.Volume("mainVolume", new()
    ///     {
    ///         Type = "l_ssd",
    ///         SizeInGb = 10,
    ///     });
    /// 
    ///     var mainServer = new Scaleway.Instance.Server("mainServer", new()
    ///     {
    ///         Image = "ubuntu_jammy",
    ///         Type = "DEV1-S",
    ///         RootVolume = new Scaleway.Instance.Inputs.ServerRootVolumeArgs
    ///         {
    ///             SizeInGb = 10,
    ///             VolumeType = "l_ssd",
    ///         },
    ///         AdditionalVolumeIds = new[]
    ///         {
    ///             mainVolume.Id,
    ///         },
    ///     });
    /// 
    ///     var mainSnapshot = new Scaleway.Instance.Snapshot("mainSnapshot", new()
    ///     {
    ///         VolumeId = mainVolume.Id,
    ///         Type = "unified",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             mainServer,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Example importing a local qcow2 file
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Scaleway.Objectstorage.Bucket("bucket");
    /// 
    ///     var qcow = new Scaleway.Objectstorage.Item("qcow", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Key = "server.qcow2",
    ///         File = "myqcow.qcow2",
    ///     });
    /// 
    ///     var snapshot = new Scaleway.Instance.Snapshot("snapshot", new()
    ///     {
    ///         Type = "unified",
    ///         Import = new Scaleway.Instance.Inputs.SnapshotImportArgs
    ///         {
    ///             Bucket = qcow.Bucket,
    ///             Key = qcow.Key,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Snapshots can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:instance/snapshot:Snapshot main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:instance/snapshot:Snapshot")]
    public partial class Snapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The snapshot creation time.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Import a snapshot from a qcow2 file located in a bucket
        /// </summary>
        [Output("import")]
        public Output<Outputs.SnapshotImport?> Import { get; private set; } = null!;

        /// <summary>
        /// The name of the snapshot. If not provided it will be randomly generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the snapshot is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the snapshot is
        /// associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// (Optional) The size of the snapshot.
        /// </summary>
        [Output("sizeInGb")]
        public Output<int> SizeInGb { get; private set; } = null!;

        /// <summary>
        /// A list of tags to apply to the snapshot.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
        /// Updates to this field will recreate a new resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The ID of the volume to take a snapshot from.
        /// </summary>
        [Output("volumeId")]
        public Output<string?> VolumeId { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which
        /// the snapshot should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:instance/snapshot:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:instance/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Import a snapshot from a qcow2 file located in a bucket
        /// </summary>
        [Input("import")]
        public Input<Inputs.SnapshotImportArgs>? Import { get; set; }

        /// <summary>
        /// The name of the snapshot. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the snapshot is
        /// associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the snapshot.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
        /// Updates to this field will recreate a new resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The ID of the volume to take a snapshot from.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        /// <summary>
        /// `zone`) The zone in which
        /// the snapshot should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public SnapshotArgs()
        {
        }
        public static new SnapshotArgs Empty => new SnapshotArgs();
    }

    public sealed class SnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The snapshot creation time.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Import a snapshot from a qcow2 file located in a bucket
        /// </summary>
        [Input("import")]
        public Input<Inputs.SnapshotImportGetArgs>? Import { get; set; }

        /// <summary>
        /// The name of the snapshot. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the snapshot is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the snapshot is
        /// associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (Optional) The size of the snapshot.
        /// </summary>
        [Input("sizeInGb")]
        public Input<int>? SizeInGb { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the snapshot.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
        /// Updates to this field will recreate a new resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The ID of the volume to take a snapshot from.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        /// <summary>
        /// `zone`) The zone in which
        /// the snapshot should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public SnapshotState()
        {
        }
        public static new SnapshotState Empty => new SnapshotState();
    }
}
