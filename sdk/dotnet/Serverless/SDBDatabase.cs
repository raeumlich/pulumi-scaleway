// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Serverless
{
    /// <summary>
    /// Creates and manages Scaleway Serverless SQL Databases. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/serverless-databases/).
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
    ///     var database = new Scaleway.Serverless.SDBDatabase("database", new()
    ///     {
    ///         MaxCpu = 8,
    ///         MinCpu = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Serverless SQL Database can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:serverless/sDBDatabase:SDBDatabase database fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:serverless/sDBDatabase:SDBDatabase")]
    public partial class SDBDatabase : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Endpoint of the database
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The maximum number of CPU units for your database. Defaults to 15.
        /// </summary>
        [Output("maxCpu")]
        public Output<int?> MaxCpu { get; private set; } = null!;

        /// <summary>
        /// The minimum number of CPU units for your database. Defaults to 0.
        /// </summary>
        [Output("minCpu")]
        public Output<int?> MinCpu { get; private set; } = null!;

        /// <summary>
        /// Name of the database (e.g. `my-new-database`).
        /// 
        /// &gt; **Important:** Updates to `name` will recreate the database.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a SDBDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SDBDatabase(string name, SDBDatabaseArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:serverless/sDBDatabase:SDBDatabase", name, args ?? new SDBDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SDBDatabase(string name, Input<string> id, SDBDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:serverless/sDBDatabase:SDBDatabase", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SDBDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SDBDatabase Get(string name, Input<string> id, SDBDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new SDBDatabase(name, id, state, options);
        }
    }

    public sealed class SDBDatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of CPU units for your database. Defaults to 15.
        /// </summary>
        [Input("maxCpu")]
        public Input<int>? MaxCpu { get; set; }

        /// <summary>
        /// The minimum number of CPU units for your database. Defaults to 0.
        /// </summary>
        [Input("minCpu")]
        public Input<int>? MinCpu { get; set; }

        /// <summary>
        /// Name of the database (e.g. `my-new-database`).
        /// 
        /// &gt; **Important:** Updates to `name` will recreate the database.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SDBDatabaseArgs()
        {
        }
        public static new SDBDatabaseArgs Empty => new SDBDatabaseArgs();
    }

    public sealed class SDBDatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint of the database
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The maximum number of CPU units for your database. Defaults to 15.
        /// </summary>
        [Input("maxCpu")]
        public Input<int>? MaxCpu { get; set; }

        /// <summary>
        /// The minimum number of CPU units for your database. Defaults to 0.
        /// </summary>
        [Input("minCpu")]
        public Input<int>? MinCpu { get; set; }

        /// <summary>
        /// Name of the database (e.g. `my-new-database`).
        /// 
        /// &gt; **Important:** Updates to `name` will recreate the database.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SDBDatabaseState()
        {
        }
        public static new SDBDatabaseState Empty => new SDBDatabaseState();
    }
}