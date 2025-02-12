// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerEndpointAddressGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// External IPv4 address. The structure is documented below.
        /// </summary>
        [Input("externalIpv4Address")]
        public Input<Inputs.AlbLoadBalancerListenerEndpointAddressExternalIpv4AddressGetArgs>? ExternalIpv4Address { get; set; }

        /// <summary>
        /// External IPv6 address. The structure is documented below.
        /// 
        /// &gt; Exactly one type of addresses `external_ipv4_address` or `internal_ipv4_address` or `external_ipv6_address` should be specified.
        /// </summary>
        [Input("externalIpv6Address")]
        public Input<Inputs.AlbLoadBalancerListenerEndpointAddressExternalIpv6AddressGetArgs>? ExternalIpv6Address { get; set; }

        /// <summary>
        /// Internal IPv4 address. The structure is documented below.
        /// </summary>
        [Input("internalIpv4Address")]
        public Input<Inputs.AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressGetArgs>? InternalIpv4Address { get; set; }

        public AlbLoadBalancerListenerEndpointAddressGetArgs()
        {
        }
        public static new AlbLoadBalancerListenerEndpointAddressGetArgs Empty => new AlbLoadBalancerListenerEndpointAddressGetArgs();
    }
}
