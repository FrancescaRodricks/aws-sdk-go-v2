// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a provisioning template version.
//
// Requires permission to access the [DescribeProvisioningTemplateVersion] action.
//
// [DescribeProvisioningTemplateVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) DescribeProvisioningTemplateVersion(ctx context.Context, params *DescribeProvisioningTemplateVersionInput, optFns ...func(*Options)) (*DescribeProvisioningTemplateVersionOutput, error) {
	if params == nil {
		params = &DescribeProvisioningTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProvisioningTemplateVersion", params, optFns, c.addOperationDescribeProvisioningTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProvisioningTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningTemplateVersionInput struct {

	// The template name.
	//
	// This member is required.
	TemplateName *string

	// The provisioning template version ID.
	//
	// This member is required.
	VersionId *int32

	noSmithyDocumentSerde
}

type DescribeProvisioningTemplateVersionOutput struct {

	// The date when the provisioning template version was created.
	CreationDate *time.Time

	// True if the provisioning template version is the default version.
	IsDefaultVersion bool

	// The JSON formatted contents of the provisioning template version.
	TemplateBody *string

	// The provisioning template version ID.
	VersionId *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProvisioningTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProvisioningTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProvisioningTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeProvisioningTemplateVersion"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpDescribeProvisioningTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningTemplateVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeProvisioningTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeProvisioningTemplateVersion",
	}
}
