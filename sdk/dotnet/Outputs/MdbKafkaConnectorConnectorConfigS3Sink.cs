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
    public sealed class MdbKafkaConnectorConnectorConfigS3Sink
    {
        /// <summary>
        /// Сompression type for messages. Cannot be changed.
        /// </summary>
        public readonly string FileCompressionType;
        /// <summary>
        /// Max records per file.
        /// </summary>
        public readonly int? FileMaxRecords;
        /// <summary>
        /// Settings for connection to s3-compatible storage. The structure is documented below.
        /// 
        /// The `s3_connection` block supports:
        /// </summary>
        public readonly Outputs.MdbKafkaConnectorConnectorConfigS3SinkS3Connection S3Connection;
        public readonly string Topics;

        [OutputConstructor]
        private MdbKafkaConnectorConnectorConfigS3Sink(
            string fileCompressionType,

            int? fileMaxRecords,

            Outputs.MdbKafkaConnectorConnectorConfigS3SinkS3Connection s3Connection,

            string topics)
        {
            FileCompressionType = fileCompressionType;
            FileMaxRecords = fileMaxRecords;
            S3Connection = s3Connection;
            Topics = topics;
        }
    }
}
