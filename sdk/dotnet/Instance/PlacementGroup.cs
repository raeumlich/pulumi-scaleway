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
    /// Creates and manages Compute Instance Placement Groups. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653).
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
    ///     var availabilityGroup = new Scaleway.Instance.PlacementGroup("availabilityGroup");
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Placement groups can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:instance/placementGroup:PlacementGroup availability_group fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:instance/placementGroup:PlacementGroup")]
    public partial class PlacementGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the placement group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the placement group is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `optional` or `enforced`.
        /// </summary>
        [Output("policyMode")]
        public Output<string?> PolicyMode { get; private set; } = null!;

        /// <summary>
        /// Is true when the policy is respected.
        /// </summary>
        [Output("policyRespected")]
        public Output<bool> PolicyRespected { get; private set; } = null!;

        /// <summary>
        /// The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `low_latency` or `max_availability`.
        /// </summary>
        [Output("policyType")]
        public Output<string?> PolicyType { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the placement group is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// A list of tags to apply to the placement group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the placement group should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a PlacementGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PlacementGroup(string name, PlacementGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:instance/placementGroup:PlacementGroup", name, args ?? new PlacementGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PlacementGroup(string name, Input<string> id, PlacementGroupState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:instance/placementGroup:PlacementGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PlacementGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PlacementGroup Get(string name, Input<string> id, PlacementGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new PlacementGroup(name, id, state, options);
        }
    }

    public sealed class PlacementGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the placement group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `optional` or `enforced`.
        /// </summary>
        [Input("policyMode")]
        public Input<string>? PolicyMode { get; set; }

        /// <summary>
        /// The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `low_latency` or `max_availability`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the placement group is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the placement group.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the placement group should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public PlacementGroupArgs()
        {
        }
        public static new PlacementGroupArgs Empty => new PlacementGroupArgs();
    }

    public sealed class PlacementGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the placement group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the placement group is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `optional` or `enforced`.
        /// </summary>
        [Input("policyMode")]
        public Input<string>? PolicyMode { get; set; }

        /// <summary>
        /// Is true when the policy is respected.
        /// </summary>
        [Input("policyRespected")]
        public Input<bool>? PolicyRespected { get; set; }

        /// <summary>
        /// The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group. Possible values are: `low_latency` or `max_availability`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the placement group is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the placement group.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the placement group should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public PlacementGroupState()
        {
        }
        public static new PlacementGroupState Empty => new PlacementGroupState();
    }
}
