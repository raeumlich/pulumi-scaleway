// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Documentdb
{
    /// <summary>
    /// Creates and manages Scaleway Database Instances.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
    /// 
    /// ## Example Usage
    /// 
    /// ### Example Basic
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
    ///     var main = new Scaleway.Documentdb.Instance("main", new()
    ///     {
    ///         Engine = "FerretDB-1",
    ///         NodeType = "docdb-play2-pico",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         Tags = new[]
    ///         {
    ///             "terraform-test",
    ///             "scaleway_documentdb_instance",
    ///             "minimal",
    ///         },
    ///         UserName = "my_initial_user",
    ///         VolumeSizeInGb = 20,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Database Instance can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:documentdb/instance:Instance db fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:documentdb/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Database Instance's engine version (e.g. `FerretDB-1`).
        /// 
        /// &gt; **Important:** Updates to `engine` will recreate the Database Instance.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Enable or disable high availability for the database instance.
        /// </summary>
        [Output("isHaCluster")]
        public Output<bool?> IsHaCluster { get; private set; } = null!;

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of database instance you want to create (e.g. `docdb-play2-pico`).
        /// 
        /// &gt; **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        /// interruption. Keep in mind that you cannot downgrade a Database Instance.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// Password for the first user of the database instance.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the Database
        /// Instance is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Database Instance should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the Database Instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
        /// 
        /// &gt; **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        /// </summary>
        [Output("telemetryEnabled")]
        public Output<bool?> TelemetryEnabled { get; private set; } = null!;

        /// <summary>
        /// Identifier for the first user of the database instance.
        /// 
        /// &gt; **Important:** Updates to `user_name` will recreate the Database Instance.
        /// </summary>
        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;

        /// <summary>
        /// Volume size (in GB) when `volume_type` is set to `bssd`.
        /// </summary>
        [Output("volumeSizeInGb")]
        public Output<int> VolumeSizeInGb { get; private set; } = null!;

        /// <summary>
        /// Type of volume where data are stored (`bssd` or `lssd`).
        /// </summary>
        [Output("volumeType")]
        public Output<string?> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("scaleway:documentdb/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:documentdb/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database Instance's engine version (e.g. `FerretDB-1`).
        /// 
        /// &gt; **Important:** Updates to `engine` will recreate the Database Instance.
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Enable or disable high availability for the database instance.
        /// </summary>
        [Input("isHaCluster")]
        public Input<bool>? IsHaCluster { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of database instance you want to create (e.g. `docdb-play2-pico`).
        /// 
        /// &gt; **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        /// interruption. Keep in mind that you cannot downgrade a Database Instance.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the first user of the database instance.
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

        /// <summary>
        /// `project_id`) The ID of the project the Database
        /// Instance is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the Database Instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
        /// 
        /// &gt; **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        /// </summary>
        [Input("telemetryEnabled")]
        public Input<bool>? TelemetryEnabled { get; set; }

        /// <summary>
        /// Identifier for the first user of the database instance.
        /// 
        /// &gt; **Important:** Updates to `user_name` will recreate the Database Instance.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Volume size (in GB) when `volume_type` is set to `bssd`.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Type of volume where data are stored (`bssd` or `lssd`).
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database Instance's engine version (e.g. `FerretDB-1`).
        /// 
        /// &gt; **Important:** Updates to `engine` will recreate the Database Instance.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Enable or disable high availability for the database instance.
        /// </summary>
        [Input("isHaCluster")]
        public Input<bool>? IsHaCluster { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of database instance you want to create (e.g. `docdb-play2-pico`).
        /// 
        /// &gt; **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        /// interruption. Keep in mind that you cannot downgrade a Database Instance.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the first user of the database instance.
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

        /// <summary>
        /// `project_id`) The ID of the project the Database
        /// Instance is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the Database Instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
        /// 
        /// &gt; **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        /// </summary>
        [Input("telemetryEnabled")]
        public Input<bool>? TelemetryEnabled { get; set; }

        /// <summary>
        /// Identifier for the first user of the database instance.
        /// 
        /// &gt; **Important:** Updates to `user_name` will recreate the Database Instance.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Volume size (in GB) when `volume_type` is set to `bssd`.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Type of volume where data are stored (`bssd` or `lssd`).
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
