// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage.Inputs
{

    public sealed class BucketACLAccessControlPolicyOwnerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The `region`,`bucket` and `acl` separated by (`/`).
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public BucketACLAccessControlPolicyOwnerGetArgs()
        {
        }
        public static new BucketACLAccessControlPolicyOwnerGetArgs Empty => new BucketACLAccessControlPolicyOwnerGetArgs();
    }
}
