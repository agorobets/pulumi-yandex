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
    public sealed class AlbBackendGroupGrpcBackendTls
    {
        /// <summary>
        /// [SNI](https://en.wikipedia.org/wiki/Server_Name_Indication) string for TLS connections.
        /// * `validation_context.0.trusted_ca_id` - (Optional) Trusted CA certificate ID in the Certificate Manager.
        /// * `validation_context.0.trusted_ca_bytes` - (Optional) PEM-encoded trusted CA certificate chain.
        /// 
        /// &gt; Only one of `validation_context.0.trusted_ca_id` or `validation_context.0.trusted_ca_bytes` should be specified.
        /// </summary>
        public readonly string? Sni;
        public readonly Outputs.AlbBackendGroupGrpcBackendTlsValidationContext? ValidationContext;

        [OutputConstructor]
        private AlbBackendGroupGrpcBackendTls(
            string? sni,

            Outputs.AlbBackendGroupGrpcBackendTlsValidationContext? validationContext)
        {
            Sni = sni;
            ValidationContext = validationContext;
        }
    }
}
