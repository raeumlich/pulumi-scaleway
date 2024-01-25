// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer.Outputs
{

    [OutputType]
    public sealed class CertficateCustomCertificate
    {
        /// <summary>
        /// Full PEM-formatted certificate chain.
        /// 
        /// &gt; **Important:** Updates to `custom_certificate` will recreate the load-balancer certificate.
        /// </summary>
        public readonly string CertificateChain;

        [OutputConstructor]
        private CertficateCustomCertificate(string certificateChain)
        {
            CertificateChain = certificateChain;
        }
    }
}