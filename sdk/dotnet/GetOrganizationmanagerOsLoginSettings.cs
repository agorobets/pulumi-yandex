// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetOrganizationmanagerOsLoginSettings
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_os_login_settings/d_organizationmanager_os_login_settings_1.tf" }}
        /// </summary>
        public static Task<GetOrganizationmanagerOsLoginSettingsResult> InvokeAsync(GetOrganizationmanagerOsLoginSettingsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationmanagerOsLoginSettingsResult>("yandex:index/getOrganizationmanagerOsLoginSettings:getOrganizationmanagerOsLoginSettings", args ?? new GetOrganizationmanagerOsLoginSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_os_login_settings/d_organizationmanager_os_login_settings_1.tf" }}
        /// </summary>
        public static Output<GetOrganizationmanagerOsLoginSettingsResult> Invoke(GetOrganizationmanagerOsLoginSettingsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationmanagerOsLoginSettingsResult>("yandex:index/getOrganizationmanagerOsLoginSettings:getOrganizationmanagerOsLoginSettings", args ?? new GetOrganizationmanagerOsLoginSettingsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_os_login_settings/d_organizationmanager_os_login_settings_1.tf" }}
        /// </summary>
        public static Output<GetOrganizationmanagerOsLoginSettingsResult> Invoke(GetOrganizationmanagerOsLoginSettingsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationmanagerOsLoginSettingsResult>("yandex:index/getOrganizationmanagerOsLoginSettings:getOrganizationmanagerOsLoginSettings", args ?? new GetOrganizationmanagerOsLoginSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationmanagerOsLoginSettingsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the organization.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// The structure is documented below.
        /// </summary>
        [Input("sshCertificateSettings")]
        public Inputs.GetOrganizationmanagerOsLoginSettingsSshCertificateSettingsArgs? SshCertificateSettings { get; set; }

        /// <summary>
        /// The structure is documented below.
        /// </summary>
        [Input("userSshKeySettings")]
        public Inputs.GetOrganizationmanagerOsLoginSettingsUserSshKeySettingsArgs? UserSshKeySettings { get; set; }

        public GetOrganizationmanagerOsLoginSettingsArgs()
        {
        }
        public static new GetOrganizationmanagerOsLoginSettingsArgs Empty => new GetOrganizationmanagerOsLoginSettingsArgs();
    }

    public sealed class GetOrganizationmanagerOsLoginSettingsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the organization.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The structure is documented below.
        /// </summary>
        [Input("sshCertificateSettings")]
        public Input<Inputs.GetOrganizationmanagerOsLoginSettingsSshCertificateSettingsInputArgs>? SshCertificateSettings { get; set; }

        /// <summary>
        /// The structure is documented below.
        /// </summary>
        [Input("userSshKeySettings")]
        public Input<Inputs.GetOrganizationmanagerOsLoginSettingsUserSshKeySettingsInputArgs>? UserSshKeySettings { get; set; }

        public GetOrganizationmanagerOsLoginSettingsInvokeArgs()
        {
        }
        public static new GetOrganizationmanagerOsLoginSettingsInvokeArgs Empty => new GetOrganizationmanagerOsLoginSettingsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationmanagerOsLoginSettingsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        /// <summary>
        /// The structure is documented below.
        /// </summary>
        public readonly Outputs.GetOrganizationmanagerOsLoginSettingsSshCertificateSettingsResult? SshCertificateSettings;
        /// <summary>
        /// The structure is documented below.
        /// </summary>
        public readonly Outputs.GetOrganizationmanagerOsLoginSettingsUserSshKeySettingsResult? UserSshKeySettings;

        [OutputConstructor]
        private GetOrganizationmanagerOsLoginSettingsResult(
            string id,

            string organizationId,

            Outputs.GetOrganizationmanagerOsLoginSettingsSshCertificateSettingsResult? sshCertificateSettings,

            Outputs.GetOrganizationmanagerOsLoginSettingsUserSshKeySettingsResult? userSshKeySettings)
        {
            Id = id;
            OrganizationId = organizationId;
            SshCertificateSettings = sshCertificateSettings;
            UserSshKeySettings = userSshKeySettings;
        }
    }
}
