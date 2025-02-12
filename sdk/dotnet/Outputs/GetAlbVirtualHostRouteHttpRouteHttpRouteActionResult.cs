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
    public sealed class GetAlbVirtualHostRouteHttpRouteHttpRouteActionResult
    {
        /// <summary>
        /// If set, will automatically rewrite host.
        /// </summary>
        public readonly bool AutoHostRewrite;
        /// <summary>
        /// Backend group to route requests.
        /// </summary>
        public readonly string BackendGroupId;
        /// <summary>
        /// Host rewrite specifier.
        /// </summary>
        public readonly string HostRewrite;
        /// <summary>
        /// Specifies the idle timeout (time without any data transfer for the active request) for the route. It is useful for streaming scenarios - one should set idle_timeout to something meaningful and max_timeout to the maximum time the stream is allowed to be alive. If not specified, there is no per-route idle timeout.
        /// </summary>
        public readonly string IdleTimeout;
        /// <summary>
        /// If not empty, matched path prefix will be replaced by this value.
        /// </summary>
        public readonly string PrefixRewrite;
        /// <summary>
        /// Rate limit configuration applied for a whole virtual host
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlbVirtualHostRouteHttpRouteHttpRouteActionRateLimitResult> RateLimits;
        /// <summary>
        /// Specifies the request timeout (overall time request processing is allowed to take) for the route. If not set, default is 60 seconds.
        /// </summary>
        public readonly string Timeout;
        /// <summary>
        /// List of upgrade types. Only specified upgrade types will be allowed. For example, "websocket".
        /// </summary>
        public readonly ImmutableArray<string> UpgradeTypes;

        [OutputConstructor]
        private GetAlbVirtualHostRouteHttpRouteHttpRouteActionResult(
            bool autoHostRewrite,

            string backendGroupId,

            string hostRewrite,

            string idleTimeout,

            string prefixRewrite,

            ImmutableArray<Outputs.GetAlbVirtualHostRouteHttpRouteHttpRouteActionRateLimitResult> rateLimits,

            string timeout,

            ImmutableArray<string> upgradeTypes)
        {
            AutoHostRewrite = autoHostRewrite;
            BackendGroupId = backendGroupId;
            HostRewrite = hostRewrite;
            IdleTimeout = idleTimeout;
            PrefixRewrite = prefixRewrite;
            RateLimits = rateLimits;
            Timeout = timeout;
            UpgradeTypes = upgradeTypes;
        }
    }
}
