// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressArgs : global::Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressArgs()
        {
        }
        public static new AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressArgs Empty => new AlbLoadBalancerListenerEndpointAddressInternalIpv4AddressArgs();
    }
}
