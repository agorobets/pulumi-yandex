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
    public sealed class VpcRouteTableStaticRoute
    {
        /// <summary>
        /// Route prefix in CIDR notation.
        /// </summary>
        public readonly string? DestinationPrefix;
        /// <summary>
        /// ID of the gateway used ad next hop.
        /// 
        /// &gt; Only one of `next_hop_address` or `gateway_id` should be specified.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// Address of the next hop.
        /// </summary>
        public readonly string? NextHopAddress;

        [OutputConstructor]
        private VpcRouteTableStaticRoute(
            string? destinationPrefix,

            string? gatewayId,

            string? nextHopAddress)
        {
            DestinationPrefix = destinationPrefix;
            GatewayId = gatewayId;
            NextHopAddress = nextHopAddress;
        }
    }
}
