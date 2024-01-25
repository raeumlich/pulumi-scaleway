// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Iot
{
    /// <summary>
    /// ## Example Usage
    /// ### S3 Route
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainHub = new Scaleway.Iot.Hub("mainHub", new()
    ///     {
    ///         ProductPlan = "plan_shared",
    ///     });
    /// 
    ///     var mainBucket = new Scaleway.Objectstorage.Bucket("mainBucket", new()
    ///     {
    ///         Region = "fr-par",
    ///     });
    /// 
    ///     var mainRoute = new Scaleway.Iot.Route("mainRoute", new()
    ///     {
    ///         HubId = mainHub.Id,
    ///         Topic = "#",
    ///         S3 = new Scaleway.Iot.Inputs.RouteS3Args
    ///         {
    ///             BucketRegion = mainBucket.Region,
    ///             BucketName = mainBucket.Name,
    ///             ObjectPrefix = "foo",
    ///             Strategy = "per_topic",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Rest Route
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainHub = new Scaleway.Iot.Hub("mainHub", new()
    ///     {
    ///         ProductPlan = "plan_shared",
    ///     });
    /// 
    ///     var mainRoute = new Scaleway.Iot.Route("mainRoute", new()
    ///     {
    ///         HubId = mainHub.Id,
    ///         Topic = "#",
    ///         Rest = new Scaleway.Iot.Inputs.RouteRestArgs
    ///         {
    ///             Verb = "get",
    ///             Uri = "http://scaleway.com",
    ///             Headers = 
    ///             {
    ///                 { "X-awesome-header", "my-awesome-value" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoT Routes can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:iot/route:Route route01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:iot/route:Route")]
    public partial class Route : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time the Route was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        /// </summary>
        [Output("database")]
        public Output<Outputs.RouteDatabase?> Database { get; private set; } = null!;

        /// <summary>
        /// The hub ID to which the Route will be attached to.
        /// </summary>
        [Output("hubId")]
        public Output<string> HubId { get; private set; } = null!;

        /// <summary>
        /// The name of the IoT Route you want to create (e.g. `my-route`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Defaults to provider `region`) The region in which the Route is attached to.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        /// </summary>
        [Output("rest")]
        public Output<Outputs.RouteRest?> Rest { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        /// </summary>
        [Output("s3")]
        public Output<Outputs.RouteS3?> S3 { get; private set; } = null!;

        /// <summary>
        /// The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("scaleway:iot/route:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:iot/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        /// </summary>
        [Input("database")]
        public Input<Inputs.RouteDatabaseArgs>? Database { get; set; }

        /// <summary>
        /// The hub ID to which the Route will be attached to.
        /// </summary>
        [Input("hubId", required: true)]
        public Input<string> HubId { get; set; } = null!;

        /// <summary>
        /// The name of the IoT Route you want to create (e.g. `my-route`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Defaults to provider `region`) The region in which the Route is attached to.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        /// </summary>
        [Input("rest")]
        public Input<Inputs.RouteRestArgs>? Rest { get; set; }

        /// <summary>
        /// Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        /// </summary>
        [Input("s3")]
        public Input<Inputs.RouteS3Args>? S3 { get; set; }

        /// <summary>
        /// The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public RouteArgs()
        {
        }
        public static new RouteArgs Empty => new RouteArgs();
    }

    public sealed class RouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time the Route was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        /// </summary>
        [Input("database")]
        public Input<Inputs.RouteDatabaseGetArgs>? Database { get; set; }

        /// <summary>
        /// The hub ID to which the Route will be attached to.
        /// </summary>
        [Input("hubId")]
        public Input<string>? HubId { get; set; }

        /// <summary>
        /// The name of the IoT Route you want to create (e.g. `my-route`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Defaults to provider `region`) The region in which the Route is attached to.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        /// </summary>
        [Input("rest")]
        public Input<Inputs.RouteRestGetArgs>? Rest { get; set; }

        /// <summary>
        /// Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        /// </summary>
        [Input("s3")]
        public Input<Inputs.RouteS3GetArgs>? S3 { get; set; }

        /// <summary>
        /// The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public RouteState()
        {
        }
        public static new RouteState Empty => new RouteState();
    }
}