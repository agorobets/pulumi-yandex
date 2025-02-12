// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupStreamBackendLoadBalancingConfigInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Percent of traffic to be sent to the same availability zone. The rest will be equally divided between other zones.
        /// </summary>
        [Input("localityAwareRoutingPercent", required: true)]
        public Input<int> LocalityAwareRoutingPercent { get; set; } = null!;

        /// <summary>
        /// Load balancing mode for the backend.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// If percentage of healthy hosts in the backend is lower than panic_threshold, traffic will be routed to all backends no matter what the health status is. This helps to avoid healthy backends overloading when everything is bad. Zero means no panic threshold.
        /// </summary>
        [Input("panicThreshold", required: true)]
        public Input<int> PanicThreshold { get; set; } = null!;

        /// <summary>
        /// If set, will route requests only to the same availability zone. Balancer won't know about endpoints in other zones.
        /// </summary>
        [Input("strictLocality", required: true)]
        public Input<bool> StrictLocality { get; set; } = null!;

        public GetAlbBackendGroupStreamBackendLoadBalancingConfigInputArgs()
        {
        }
        public static new GetAlbBackendGroupStreamBackendLoadBalancingConfigInputArgs Empty => new GetAlbBackendGroupStreamBackendLoadBalancingConfigInputArgs();
    }
}
