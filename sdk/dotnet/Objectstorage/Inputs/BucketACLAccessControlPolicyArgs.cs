// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage.Inputs
{

    public sealed class BucketACLAccessControlPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("grants")]
        private InputList<Inputs.BucketACLAccessControlPolicyGrantArgs>? _grants;
        public InputList<Inputs.BucketACLAccessControlPolicyGrantArgs> Grants
        {
            get => _grants ?? (_grants = new InputList<Inputs.BucketACLAccessControlPolicyGrantArgs>());
            set => _grants = value;
        }

        /// <summary>
        /// Configuration block of the bucket project owner's display organization ID.
        /// </summary>
        [Input("owner", required: true)]
        public Input<Inputs.BucketACLAccessControlPolicyOwnerArgs> Owner { get; set; } = null!;

        public BucketACLAccessControlPolicyArgs()
        {
        }
        public static new BucketACLAccessControlPolicyArgs Empty => new BucketACLAccessControlPolicyArgs();
    }
}
