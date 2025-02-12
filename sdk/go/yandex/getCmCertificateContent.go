// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get content (certificate, private key) from a Yandex Certificate Manager Certificate. For more information, see [the official documentation](https://yandex.cloud/docs/certificate-manager/concepts/).
//
// ## Example Usage
//
// {{ tffile "examples/cm_certificate_content/d_cm_certificate_content_1.tf" }}
//
// This data source is used to define contents of [Certificate Manager Certificate](https://yandex.cloud/docs/certificate-manager/concepts/) that can be used by other resources. Can also be used to wait for certificate validation.
func GetCmCertificateContent(ctx *pulumi.Context, args *GetCmCertificateContentArgs, opts ...pulumi.InvokeOption) (*GetCmCertificateContentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCmCertificateContentResult
	err := ctx.Invoke("yandex:index/getCmCertificateContent:getCmCertificateContent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCmCertificateContent.
type GetCmCertificateContentArgs struct {
	// Certificate Id.
	CertificateId *string `pulumi:"certificateId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Certificate name.
	Name *string `pulumi:"name"`
	// Format in which you want to export the private_key: `"PKCS1"` or `"PKCS8"`.
	//
	// > One of `certificateId` or `name` should be specified.
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// If `true`, the operation won't be completed while the certificate is in `VALIDATING`.
	WaitValidation *bool `pulumi:"waitValidation"`
}

// A collection of values returned by getCmCertificateContent.
type GetCmCertificateContentResult struct {
	CertificateId *string `pulumi:"certificateId"`
	// List of certificates in chain.
	Certificates []string `pulumi:"certificates"`
	FolderId     *string  `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// Private key in specified format.
	PrivateKey       string  `pulumi:"privateKey"`
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	WaitValidation   *bool   `pulumi:"waitValidation"`
}

func GetCmCertificateContentOutput(ctx *pulumi.Context, args GetCmCertificateContentOutputArgs, opts ...pulumi.InvokeOption) GetCmCertificateContentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCmCertificateContentResultOutput, error) {
			args := v.(GetCmCertificateContentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getCmCertificateContent:getCmCertificateContent", args, GetCmCertificateContentResultOutput{}, options).(GetCmCertificateContentResultOutput), nil
		}).(GetCmCertificateContentResultOutput)
}

// A collection of arguments for invoking getCmCertificateContent.
type GetCmCertificateContentOutputArgs struct {
	// Certificate Id.
	CertificateId pulumi.StringPtrInput `pulumi:"certificateId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// Certificate name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Format in which you want to export the private_key: `"PKCS1"` or `"PKCS8"`.
	//
	// > One of `certificateId` or `name` should be specified.
	PrivateKeyFormat pulumi.StringPtrInput `pulumi:"privateKeyFormat"`
	// If `true`, the operation won't be completed while the certificate is in `VALIDATING`.
	WaitValidation pulumi.BoolPtrInput `pulumi:"waitValidation"`
}

func (GetCmCertificateContentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCmCertificateContentArgs)(nil)).Elem()
}

// A collection of values returned by getCmCertificateContent.
type GetCmCertificateContentResultOutput struct{ *pulumi.OutputState }

func (GetCmCertificateContentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCmCertificateContentResult)(nil)).Elem()
}

func (o GetCmCertificateContentResultOutput) ToGetCmCertificateContentResultOutput() GetCmCertificateContentResultOutput {
	return o
}

func (o GetCmCertificateContentResultOutput) ToGetCmCertificateContentResultOutputWithContext(ctx context.Context) GetCmCertificateContentResultOutput {
	return o
}

func (o GetCmCertificateContentResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// List of certificates in chain.
func (o GetCmCertificateContentResultOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o GetCmCertificateContentResultOutput) FolderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.FolderId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCmCertificateContentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCmCertificateContentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Private key in specified format.
func (o GetCmCertificateContentResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o GetCmCertificateContentResultOutput) PrivateKeyFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.PrivateKeyFormat }).(pulumi.StringPtrOutput)
}

func (o GetCmCertificateContentResultOutput) WaitValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *bool { return v.WaitValidation }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCmCertificateContentResultOutput{})
}
