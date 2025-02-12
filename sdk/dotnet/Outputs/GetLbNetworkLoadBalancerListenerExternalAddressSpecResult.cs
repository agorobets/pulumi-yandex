// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetLbNetworkLoadBalancerListenerExternalAddressSpecResult
    {
        /// <summary>
        /// External IP address of a listener.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// IP version of the external addresses.
        /// </summary>
        public readonly string IpVersion;

        [OutputConstructor]
        private GetLbNetworkLoadBalancerListenerExternalAddressSpecResult(
            string address,

            string ipVersion)
        {
            Address = address;
            IpVersion = ipVersion;
        }
    }
}
