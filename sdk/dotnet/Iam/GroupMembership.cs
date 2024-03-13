// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Iam
{
    /// <summary>
    /// Add members to an IAM group.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).
    /// 
    /// ## Example Usage
    /// 
    /// ### Application Membership
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
    ///     var @group = new Scaleway.Iam.Group("group", new()
    ///     {
    ///         ExternalMembership = true,
    ///     });
    /// 
    ///     var app = new Scaleway.Iam.Application("app");
    /// 
    ///     var member = new Scaleway.Iam.GroupMembership("member", new()
    ///     {
    ///         GroupId = @group.Id,
    ///         ApplicationId = app.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// IAM group memberships can be imported using two format:
    /// 
    /// - For user: `{group_id}/user/{user_id}`
    /// 
    /// - For application: `{group_id}/app/{application_id}`
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:iam/groupMembership:GroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:iam/groupMembership:GroupMembership")]
    public partial class GroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the application that will be added to the group.
        /// </summary>
        [Output("applicationId")]
        public Output<string?> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// ID of the group to add members to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the user that will be added to the group
        /// 
        /// - &gt; Only one of `application_id` or `user_id` must be specified
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMembership(string name, GroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("scaleway:iam/groupMembership:GroupMembership", name, args ?? new GroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMembership(string name, Input<string> id, GroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:iam/groupMembership:GroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMembership Get(string name, Input<string> id, GroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMembership(name, id, state, options);
        }
    }

    public sealed class GroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application that will be added to the group.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ID of the group to add members to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The ID of the user that will be added to the group
        /// 
        /// - &gt; Only one of `application_id` or `user_id` must be specified
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GroupMembershipArgs()
        {
        }
        public static new GroupMembershipArgs Empty => new GroupMembershipArgs();
    }

    public sealed class GroupMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application that will be added to the group.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ID of the group to add members to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The ID of the user that will be added to the group
        /// 
        /// - &gt; Only one of `application_id` or `user_id` must be specified
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GroupMembershipState()
        {
        }
        public static new GroupMembershipState Empty => new GroupMembershipState();
    }
}
