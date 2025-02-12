// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class FunctionTriggerMailArgs : global::Pulumi.ResourceArgs
    {
        [Input("attachmentsBucketId")]
        public Input<string>? AttachmentsBucketId { get; set; }

        [Input("batchCutoff", required: true)]
        public Input<string> BatchCutoff { get; set; } = null!;

        [Input("batchSize")]
        public Input<string>? BatchSize { get; set; }

        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        public FunctionTriggerMailArgs()
        {
        }
        public static new FunctionTriggerMailArgs Empty => new FunctionTriggerMailArgs();
    }
}
