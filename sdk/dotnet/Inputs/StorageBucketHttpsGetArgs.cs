// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class StorageBucketHttpsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// — Id of the certificate in Certificate Manager, that will be used for bucket.
        /// 
        /// The `tags` object for setting tags (or labels) for bucket. See [tags](https://yandex.cloud/docs/storage/concepts/tags) for more information.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        public StorageBucketHttpsGetArgs()
        {
        }
        public static new StorageBucketHttpsGetArgs Empty => new StorageBucketHttpsGetArgs();
    }
}
