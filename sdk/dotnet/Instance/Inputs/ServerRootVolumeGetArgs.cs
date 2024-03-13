// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Instance.Inputs
{

    public sealed class ServerRootVolumeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the volume where the boot the server
        /// </summary>
        [Input("boot")]
        public Input<bool>? Boot { get; set; }

        /// <summary>
        /// Forces deletion of the root volume on instance termination.
        /// 
        /// &gt; **Important:** Updates to `root_volume.size_in_gb` will be ignored after the creation of the server.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Size of the root volume in gigabytes.
        /// To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
        /// check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
        /// Updates to this field will recreate a new resource.
        /// </summary>
        [Input("sizeInGb")]
        public Input<int>? SizeInGb { get; set; }

        /// <summary>
        /// The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        /// <summary>
        /// Volume type of root volume, can be `b_ssd` or `l_ssd`, default value depends on server type
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public ServerRootVolumeGetArgs()
        {
        }
        public static new ServerRootVolumeGetArgs Empty => new ServerRootVolumeGetArgs();
    }
}
