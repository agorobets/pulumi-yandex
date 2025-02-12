// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The ID of the [Cloud](https://yandex.cloud/docs/resource-manager/concepts/resources-hierarchy#cloud) to apply any
// resources to. This can also be specified using environment variable `YC_CLOUD_ID`.
func GetCloudId(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:cloudId")
}

// The endpoint for API calls, default value is **api.cloud.yandex.net:443**. This can also be defined by environment
// variable `YC_ENDPOINT`.
func GetEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:endpoint")
}

// The ID of the [Folder](https://yandex.cloud/docs/resource-manager/concepts/resources-hierarchy#folder) to operate under,
// if not specified by a given resource. This can also be specified using environment variable `YC_FOLDER_ID`.
func GetFolderId(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:folderId")
}

// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`.
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "yandex:insecure")
}

// This is the maximum number of times an API call is retried, in the case where requests are being throttled or
// experiencing transient failures. The delay between the subsequent API calls increases exponentially.
func GetMaxRetries(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "yandex:maxRetries")
}

// The ID of the [Cloud Organization](https://yandex.cloud/docs/organization/quickstart) to operate under.
func GetOrganizationId(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:organizationId")
}

// Disable use of TLS. Default value is `false`.
func GetPlaintext(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "yandex:plaintext")
}

// Profile name to use in the shared credentials file. Default value is `default`.
func GetProfile(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:profile")
}

// [The region](https://yandex.cloud/docs/overview/concepts/region) where operations will take place. For example
// `ru-central1`.
func GetRegionId(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:regionId")
}

// Contains either a path to or the contents of the [Service Account
// file](https://yandex.cloud/docs/iam/concepts/authorization/key) in JSON format. This can also be specified using
// environment variable `YC_SERVICE_ACCOUNT_KEY_FILE`. You can read how to create service account key file
// [here](https://yandex.cloud/docs/iam/operations/iam-token/create-for-sa#keys-create). > Only one of `token` or
// `serviceAccountKeyFile` must be specified. > One can authenticate via instance service account from inside a compute
// instance. In order to use this method, omit both `token`/`serviceAccountKeyFile` and attach service account to the
// instance. [Working with Yandex Cloud from inside an
// instance](https://yandex.cloud/docs/compute/operations/vm-connect/auth-inside-vm).
func GetServiceAccountKeyFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:serviceAccountKeyFile")
}

// Shared credentials file path. Supported keys: `storageAccessKey` and `storageSecretKey`. > The `storageAccessKey` and
// `storageSecretKey` attributes from the shared credentials file are used only when the provider and a storage
// data/resource do not have an access/secret keys explicitly specified.
func GetSharedCredentialsFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:sharedCredentialsFile")
}

// Yandex Cloud Object Storage access key, which is used when a storage data/resource doesn't have an access key explicitly
// specified. This can also be specified using environment variable `YC_STORAGE_ACCESS_KEY`.
func GetStorageAccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:storageAccessKey")
}

// Yandex Cloud [Object Storage Endpoint](https://yandex.cloud/docs/storage/s3/#request-url), which is used to connect to
// `S3 API`. Default value is **storage.yandexcloud.net**.
func GetStorageEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:storageEndpoint")
}

// Yandex Cloud Object Storage secret key, which is used when a storage data/resource doesn't have a secret key explicitly
// specified. This can also be specified using environment variable `YC_STORAGE_SECRET_KEY`.
func GetStorageSecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:storageSecretKey")
}

// Security token or IAM token used for authentication in Yandex Cloud. Check
// [documentation](https://yandex.cloud/docs/iam/operations/iam-token/create) about how to create IAM token. This can also
// be specified using environment variable `YC_TOKEN`.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:token")
}

// Yandex Cloud Message Queue service access key, which is used when a YMQ queue resource doesn't have an access key
// explicitly specified. This can also be specified using environment variable `YC_MESSAGE_QUEUE_ACCESS_KEY`.
func GetYmqAccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:ymqAccessKey")
}

// Yandex Cloud Message Queue service endpoint. Default value is **message-queue.api.cloud.yandex.net**.
func GetYmqEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:ymqEndpoint")
}

// Yandex Cloud Message Queue service secret key, which is used when a YMQ queue resource doesn't have a secret key
// explicitly specified. This can also be specified using environment variable `YC_MESSAGE_QUEUE_SECRET_KEY`.
func GetYmqSecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:ymqSecretKey")
}

// The default [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) to operate under, if not
// specified by a given resource. This can also be specified using environment variable `YC_ZONE`.
func GetZone(ctx *pulumi.Context) string {
	return config.Get(ctx, "yandex:zone")
}
