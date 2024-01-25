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
    public sealed class GetBackendHealthCheckHttpResult
    {
        public readonly int Code;
        public readonly string HostHeader;
        public readonly string Method;
        public readonly string Sni;
        public readonly string Uri;

        [OutputConstructor]
        private GetBackendHealthCheckHttpResult(
            int code,

            string hostHeader,

            string method,

            string sni,

            string uri)
        {
            Code = code;
            HostHeader = hostHeader;
            Method = method;
            Sni = sni;
            Uri = uri;
        }
    }
}