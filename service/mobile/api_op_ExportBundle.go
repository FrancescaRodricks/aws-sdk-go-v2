// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Generates customized software development kit (SDK) and or tool packages used
//
// to integrate mobile web or mobile app clients with backend AWS resources.
func (c *Client) ExportBundle(ctx context.Context, params *ExportBundleInput, optFns ...func(*Options)) (*ExportBundleOutput, error) {
	if params == nil {
		params = &ExportBundleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportBundle", params, optFns, c.addOperationExportBundleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//	Request structure used to request generation of custom SDK and tool packages
//
// required to integrate mobile web or app clients with backed AWS resources.
type ExportBundleInput struct {

	//  Unique bundle identifier.
	//
	// This member is required.
	BundleId *string

	//  Developer desktop or target application platform.
	Platform types.Platform

	//  Unique project identifier.
	ProjectId *string

	noSmithyDocumentSerde
}

//	Result structure which contains link to download custom-generated SDK and tool
//
// packages used to integrate mobile web or app clients with backed AWS resources.
type ExportBundleOutput struct {

	//  URL which contains the custom-generated SDK and tool packages used to
	// integrate the client mobile app or web app with the AWS resources created by the
	// AWS Mobile Hub project.
	DownloadUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportBundleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportBundle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportBundle{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportBundle"); err != nil {
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
	if err = addOpExportBundleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportBundle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportBundle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportBundle",
	}
}
