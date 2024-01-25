// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage
{
    /// <summary>
    /// Creates and manages Scaleway object storage objects.
    /// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
    /// 
    /// ## Import
    /// 
    /// Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:objectstorage/item:Item some_object fr-par/some-bucket/some-file
    /// ```
    /// 
    ///  ~&gt; **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project. If you are using a project different from the default one, you have to specify the project ID at the end of the import command. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:objectstorage/item:Item some_object fr-par/some-bucket/some-file@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:objectstorage/item:Item")]
    public partial class Item : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Output("contentBase64")]
        public Output<string?> ContentBase64 { get; private set; } = null!;

        /// <summary>
        /// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Output("file")]
        public Output<string?> File { get; private set; } = null!;

        /// <summary>
        /// Hash of the file, used to trigger upload on file change
        /// </summary>
        [Output("hash")]
        public Output<string?> Hash { get; private set; } = null!;

        /// <summary>
        /// The path of the object.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Map of metadata used for the object, keys must be lowercase
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Scaleway region this bucket resides in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        /// </summary>
        [Output("storageClass")]
        public Output<string?> StorageClass { get; private set; } = null!;

        /// <summary>
        /// Map of tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Visibility of the object, `public-read` or `private`
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Item resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Item(string name, ItemArgs args, CustomResourceOptions? options = null)
            : base("scaleway:objectstorage/item:Item", name, args ?? new ItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Item(string name, Input<string> id, ItemState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:objectstorage/item:Item", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Item resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Item Get(string name, Input<string> id, ItemState? state = null, CustomResourceOptions? options = null)
        {
            return new Item(name, id, state, options);
        }
    }

    public sealed class ItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Input("contentBase64")]
        public Input<string>? ContentBase64 { get; set; }

        /// <summary>
        /// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// Hash of the file, used to trigger upload on file change
        /// </summary>
        [Input("hash")]
        public Input<string>? Hash { get; set; }

        /// <summary>
        /// The path of the object.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map of metadata used for the object, keys must be lowercase
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Scaleway region this bucket resides in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Visibility of the object, `public-read` or `private`
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ItemArgs()
        {
        }
        public static new ItemArgs Empty => new ItemArgs();
    }

    public sealed class ItemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Input("contentBase64")]
        public Input<string>? ContentBase64 { get; set; }

        /// <summary>
        /// The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// Hash of the file, used to trigger upload on file change
        /// </summary>
        [Input("hash")]
        public Input<string>? Hash { get; set; }

        /// <summary>
        /// The path of the object.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map of metadata used for the object, keys must be lowercase
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Scaleway region this bucket resides in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Visibility of the object, `public-read` or `private`
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ItemState()
        {
        }
        public static new ItemState Empty => new ItemState();
    }
}
