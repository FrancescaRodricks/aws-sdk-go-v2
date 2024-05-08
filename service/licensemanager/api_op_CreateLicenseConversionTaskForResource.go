// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new license conversion task.
func (c *Client) CreateLicenseConversionTaskForResource(ctx context.Context, params *CreateLicenseConversionTaskForResourceInput, optFns ...func(*Options)) (*CreateLicenseConversionTaskForResourceOutput, error) {
	if params == nil {
		params = &CreateLicenseConversionTaskForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLicenseConversionTaskForResource", params, optFns, c.addOperationCreateLicenseConversionTaskForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLicenseConversionTaskForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLicenseConversionTaskForResourceInput struct {

	// Information that identifies the license type you are converting to. For the
	// structure of the destination license, see [Convert a license type using the CLI]in the License Manager User Guide.
	//
	// [Convert a license type using the CLI]: https://docs.aws.amazon.com/license-manager/latest/userguide/conversion-procedures.html#conversion-cli
	//
	// This member is required.
	DestinationLicenseContext *types.LicenseConversionContext

	// Amazon Resource Name (ARN) of the resource you are converting the license type
	// for.
	//
	// This member is required.
	ResourceArn *string

	// Information that identifies the license type you are converting from.
	//
	// For the structure of the source license, see [Convert a license type using the CLI] in the License Manager User Guide.
	//
	// [Convert a license type using the CLI]: https://docs.aws.amazon.com/license-manager/latest/userguide/conversion-procedures.html#conversion-cli
	//
	// This member is required.
	SourceLicenseContext *types.LicenseConversionContext

	noSmithyDocumentSerde
}

type CreateLicenseConversionTaskForResourceOutput struct {

	// The ID of the created license type conversion task.
	LicenseConversionTaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLicenseConversionTaskForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLicenseConversionTaskForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLicenseConversionTaskForResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLicenseConversionTaskForResource"); err != nil {
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
	if err = addOpCreateLicenseConversionTaskForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLicenseConversionTaskForResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLicenseConversionTaskForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLicenseConversionTaskForResource",
	}
}
