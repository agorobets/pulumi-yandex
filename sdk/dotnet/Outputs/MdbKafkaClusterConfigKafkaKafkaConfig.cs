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
    public sealed class MdbKafkaClusterConfigKafkaKafkaConfig
    {
        public readonly bool? AutoCreateTopicsEnable;
        /// <summary>
        /// , `log_flush_interval_messages`, `log_flush_interval_ms`, `log_flush_scheduler_interval_ms`, `log_retention_bytes`, `log_retention_hours`, `log_retention_minutes`, `log_retention_ms`, `log_segment_bytes`, `log_preallocate`, `socket_send_buffer_bytes`, `socket_receive_buffer_bytes`, `auto_create_topics_enable`, `num_partitions`, `default_replication_factor`, `message_max_bytes`, `replica_fetch_max_bytes`, `ssl_cipher_suites`, `offsets_retention_minutes`, `sasl_enabled_mechanisms` - (Optional) Kafka server settings. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kafka/operations/cluster-update) and [the Kafka documentation](https://kafka.apache.org/documentation/#configuration).
        /// </summary>
        public readonly string? CompressionType;
        public readonly string? DefaultReplicationFactor;
        public readonly string? LogFlushIntervalMessages;
        public readonly string? LogFlushIntervalMs;
        public readonly string? LogFlushSchedulerIntervalMs;
        public readonly bool? LogPreallocate;
        public readonly string? LogRetentionBytes;
        public readonly string? LogRetentionHours;
        public readonly string? LogRetentionMinutes;
        public readonly string? LogRetentionMs;
        public readonly string? LogSegmentBytes;
        public readonly string? MessageMaxBytes;
        public readonly string? NumPartitions;
        public readonly string? OffsetsRetentionMinutes;
        public readonly string? ReplicaFetchMaxBytes;
        public readonly ImmutableArray<string> SaslEnabledMechanisms;
        public readonly string? SocketReceiveBufferBytes;
        public readonly string? SocketSendBufferBytes;
        public readonly ImmutableArray<string> SslCipherSuites;

        [OutputConstructor]
        private MdbKafkaClusterConfigKafkaKafkaConfig(
            bool? autoCreateTopicsEnable,

            string? compressionType,

            string? defaultReplicationFactor,

            string? logFlushIntervalMessages,

            string? logFlushIntervalMs,

            string? logFlushSchedulerIntervalMs,

            bool? logPreallocate,

            string? logRetentionBytes,

            string? logRetentionHours,

            string? logRetentionMinutes,

            string? logRetentionMs,

            string? logSegmentBytes,

            string? messageMaxBytes,

            string? numPartitions,

            string? offsetsRetentionMinutes,

            string? replicaFetchMaxBytes,

            ImmutableArray<string> saslEnabledMechanisms,

            string? socketReceiveBufferBytes,

            string? socketSendBufferBytes,

            ImmutableArray<string> sslCipherSuites)
        {
            AutoCreateTopicsEnable = autoCreateTopicsEnable;
            CompressionType = compressionType;
            DefaultReplicationFactor = defaultReplicationFactor;
            LogFlushIntervalMessages = logFlushIntervalMessages;
            LogFlushIntervalMs = logFlushIntervalMs;
            LogFlushSchedulerIntervalMs = logFlushSchedulerIntervalMs;
            LogPreallocate = logPreallocate;
            LogRetentionBytes = logRetentionBytes;
            LogRetentionHours = logRetentionHours;
            LogRetentionMinutes = logRetentionMinutes;
            LogRetentionMs = logRetentionMs;
            LogSegmentBytes = logSegmentBytes;
            MessageMaxBytes = messageMaxBytes;
            NumPartitions = numPartitions;
            OffsetsRetentionMinutes = offsetsRetentionMinutes;
            ReplicaFetchMaxBytes = replicaFetchMaxBytes;
            SaslEnabledMechanisms = saslEnabledMechanisms;
            SocketReceiveBufferBytes = socketReceiveBufferBytes;
            SocketSendBufferBytes = socketSendBufferBytes;
            SslCipherSuites = sslCipherSuites;
        }
    }
}
