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
    public sealed class AlbLoadBalancerListener
    {
        /// <summary>
        /// Network endpoints (addresses and ports) of the listener. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlbLoadBalancerListenerEndpoint> Endpoints;
        /// <summary>
        /// HTTP listener resource. The structure is documented below.
        /// </summary>
        public readonly Outputs.AlbLoadBalancerListenerHttp? Http;
        /// <summary>
        /// name of the listener.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Stream listener resource. The structure is documented below.
        /// </summary>
        public readonly Outputs.AlbLoadBalancerListenerStream? Stream;
        /// <summary>
        /// TLS listener resource. The structure is documented below.
        /// 
        /// &gt; Exactly one listener type: `http` or `tls` or `stream` should be specified.
        /// </summary>
        public readonly Outputs.AlbLoadBalancerListenerTls? Tls;

        [OutputConstructor]
        private AlbLoadBalancerListener(
            ImmutableArray<Outputs.AlbLoadBalancerListenerEndpoint> endpoints,

            Outputs.AlbLoadBalancerListenerHttp? http,

            string name,

            Outputs.AlbLoadBalancerListenerStream? stream,

            Outputs.AlbLoadBalancerListenerTls? tls)
        {
            Endpoints = endpoints;
            Http = http;
            Name = name;
            Stream = stream;
            Tls = tls;
        }
    }
}
