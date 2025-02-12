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
    public sealed class GetAlbHttpRouterRouteOptionRbacPrincipalAndPrincipalHeaderResult
    {
        /// <summary>
        /// Name of the HTTP Router.
        /// 
        /// &gt; One of `http_router_id` or `name` should be specified.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetAlbHttpRouterRouteOptionRbacPrincipalAndPrincipalHeaderValueResult> Values;

        [OutputConstructor]
        private GetAlbHttpRouterRouteOptionRbacPrincipalAndPrincipalHeaderResult(
            string name,

            ImmutableArray<Outputs.GetAlbHttpRouterRouteOptionRbacPrincipalAndPrincipalHeaderValueResult> values)
        {
            Name = name;
            Values = values;
        }
    }
}
