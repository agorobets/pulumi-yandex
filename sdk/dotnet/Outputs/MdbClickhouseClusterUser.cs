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
    public sealed class MdbClickhouseClusterUser
    {
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The password of the user.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.MdbClickhouseClusterUserPermission> Permissions;
        /// <summary>
        /// Set of user quotas. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.MdbClickhouseClusterUserQuota> Quotas;
        /// <summary>
        /// Custom settings for user. The list is documented below.
        /// </summary>
        public readonly Outputs.MdbClickhouseClusterUserSettings? Settings;

        [OutputConstructor]
        private MdbClickhouseClusterUser(
            string name,

            string password,

            ImmutableArray<Outputs.MdbClickhouseClusterUserPermission> permissions,

            ImmutableArray<Outputs.MdbClickhouseClusterUserQuota> quotas,

            Outputs.MdbClickhouseClusterUserSettings? settings)
        {
            Name = name;
            Password = password;
            Permissions = permissions;
            Quotas = quotas;
            Settings = settings;
        }
    }
}
