// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Mnq
{
    /// <summary>
    /// Creates and manages Scaleway Messaging and queuing Nats Accounts.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)
    /// To use Scaleway's provider with official nats jetstream provider, check out the corresponding guide
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
    ///     var main = new Scaleway.Mnq.NATSAccount("main");
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Namespaces can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:mnq/nATSAccount:NATSAccount main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:mnq/nATSAccount:NATSAccount")]
    public partial class NATSAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the NATS service for this account.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The unique name of the nats account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the
        /// account is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which the account should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a NATSAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NATSAccount(string name, NATSAccountArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/nATSAccount:NATSAccount", name, args ?? new NATSAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NATSAccount(string name, Input<string> id, NATSAccountState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/nATSAccount:NATSAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NATSAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NATSAccount Get(string name, Input<string> id, NATSAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new NATSAccount(name, id, state, options);
        }
    }

    public sealed class NATSAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the nats account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the
        /// account is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the account should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public NATSAccountArgs()
        {
        }
        public static new NATSAccountArgs Empty => new NATSAccountArgs();
    }

    public sealed class NATSAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the NATS service for this account.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The unique name of the nats account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the
        /// account is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the account should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public NATSAccountState()
        {
        }
        public static new NATSAccountState Empty => new NATSAccountState();
    }
}
