// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferTransferRuntimeYcRuntimeUploadShardParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of workers.
        /// </summary>
        [Input("jobCount")]
        public Input<int>? JobCount { get; set; }

        /// <summary>
        /// Number of threads.
        /// </summary>
        [Input("processCount")]
        public Input<int>? ProcessCount { get; set; }

        public DatatransferTransferRuntimeYcRuntimeUploadShardParamsArgs()
        {
        }
        public static new DatatransferTransferRuntimeYcRuntimeUploadShardParamsArgs Empty => new DatatransferTransferRuntimeYcRuntimeUploadShardParamsArgs();
    }
}
