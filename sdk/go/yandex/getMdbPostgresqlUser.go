// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed PostgreSQL user. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/).
//
// ## Example Usage
//
// {{ tffile "examples/mdb_postgresql_user/d_mdb_postgresql_user_1.tf" }}
func LookupMdbPostgresqlUser(ctx *pulumi.Context, args *LookupMdbPostgresqlUserArgs, opts ...pulumi.InvokeOption) (*LookupMdbPostgresqlUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMdbPostgresqlUserResult
	err := ctx.Invoke("yandex:index/getMdbPostgresqlUser:getMdbPostgresqlUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbPostgresqlUser.
type LookupMdbPostgresqlUserArgs struct {
	// The ID of the PostgreSQL cluster.
	ClusterId string `pulumi:"clusterId"`
	// The maximum number of connections per user.
	ConnLimit *int `pulumi:"connLimit"`
	// Inhibits deletion of the user.
	DeletionProtection *string `pulumi:"deletionProtection"`
	// User's ability to login.
	Login *bool `pulumi:"login"`
	// The name of the PostgreSQL user.
	Name string `pulumi:"name"`
	// Map of user settings.
	Settings map[string]string `pulumi:"settings"`
}

// A collection of values returned by getMdbPostgresqlUser.
type LookupMdbPostgresqlUserResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The maximum number of connections per user.
	ConnLimit *int `pulumi:"connLimit"`
	// Inhibits deletion of the user.
	DeletionProtection *string `pulumi:"deletionProtection"`
	// List of the user's grants.
	Grants []string `pulumi:"grants"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// User's ability to login.
	Login *bool  `pulumi:"login"`
	Name  string `pulumi:"name"`
	// The password of the user.
	Password string `pulumi:"password"`
	// Set of permissions granted to the user. The structure is documented below.
	Permissions []GetMdbPostgresqlUserPermission `pulumi:"permissions"`
	// Map of user settings.
	Settings map[string]string `pulumi:"settings"`
}

func LookupMdbPostgresqlUserOutput(ctx *pulumi.Context, args LookupMdbPostgresqlUserOutputArgs, opts ...pulumi.InvokeOption) LookupMdbPostgresqlUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMdbPostgresqlUserResultOutput, error) {
			args := v.(LookupMdbPostgresqlUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getMdbPostgresqlUser:getMdbPostgresqlUser", args, LookupMdbPostgresqlUserResultOutput{}, options).(LookupMdbPostgresqlUserResultOutput), nil
		}).(LookupMdbPostgresqlUserResultOutput)
}

// A collection of arguments for invoking getMdbPostgresqlUser.
type LookupMdbPostgresqlUserOutputArgs struct {
	// The ID of the PostgreSQL cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The maximum number of connections per user.
	ConnLimit pulumi.IntPtrInput `pulumi:"connLimit"`
	// Inhibits deletion of the user.
	DeletionProtection pulumi.StringPtrInput `pulumi:"deletionProtection"`
	// User's ability to login.
	Login pulumi.BoolPtrInput `pulumi:"login"`
	// The name of the PostgreSQL user.
	Name pulumi.StringInput `pulumi:"name"`
	// Map of user settings.
	Settings pulumi.StringMapInput `pulumi:"settings"`
}

func (LookupMdbPostgresqlUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbPostgresqlUserArgs)(nil)).Elem()
}

// A collection of values returned by getMdbPostgresqlUser.
type LookupMdbPostgresqlUserResultOutput struct{ *pulumi.OutputState }

func (LookupMdbPostgresqlUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbPostgresqlUserResult)(nil)).Elem()
}

func (o LookupMdbPostgresqlUserResultOutput) ToLookupMdbPostgresqlUserResultOutput() LookupMdbPostgresqlUserResultOutput {
	return o
}

func (o LookupMdbPostgresqlUserResultOutput) ToLookupMdbPostgresqlUserResultOutputWithContext(ctx context.Context) LookupMdbPostgresqlUserResultOutput {
	return o
}

func (o LookupMdbPostgresqlUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The maximum number of connections per user.
func (o LookupMdbPostgresqlUserResultOutput) ConnLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) *int { return v.ConnLimit }).(pulumi.IntPtrOutput)
}

// Inhibits deletion of the user.
func (o LookupMdbPostgresqlUserResultOutput) DeletionProtection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) *string { return v.DeletionProtection }).(pulumi.StringPtrOutput)
}

// List of the user's grants.
func (o LookupMdbPostgresqlUserResultOutput) Grants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) []string { return v.Grants }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbPostgresqlUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// User's ability to login.
func (o LookupMdbPostgresqlUserResultOutput) Login() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) *bool { return v.Login }).(pulumi.BoolPtrOutput)
}

func (o LookupMdbPostgresqlUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// The password of the user.
func (o LookupMdbPostgresqlUserResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) string { return v.Password }).(pulumi.StringOutput)
}

// Set of permissions granted to the user. The structure is documented below.
func (o LookupMdbPostgresqlUserResultOutput) Permissions() GetMdbPostgresqlUserPermissionArrayOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) []GetMdbPostgresqlUserPermission { return v.Permissions }).(GetMdbPostgresqlUserPermissionArrayOutput)
}

// Map of user settings.
func (o LookupMdbPostgresqlUserResultOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbPostgresqlUserResult) map[string]string { return v.Settings }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbPostgresqlUserResultOutput{})
}
