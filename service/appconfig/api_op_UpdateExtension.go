// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an AppConfig extension. For more information about extensions, see [Extending workflows] in
// the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func (c *Client) UpdateExtension(ctx context.Context, params *UpdateExtensionInput, optFns ...func(*Options)) (*UpdateExtensionOutput, error) {
	if params == nil {
		params = &UpdateExtensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateExtension", params, optFns, c.addOperationUpdateExtensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExtensionInput struct {

	// The name, the ID, or the Amazon Resource Name (ARN) of the extension.
	//
	// This member is required.
	ExtensionIdentifier *string

	// The actions defined in the extension.
	Actions map[string][]types.Action

	// Information about the extension.
	Description *string

	// One or more parameters for the actions called by the extension.
	Parameters map[string]types.Parameter

	// The extension version number.
	VersionNumber *int32

	noSmithyDocumentSerde
}

type UpdateExtensionOutput struct {

	// The actions defined in the extension.
	Actions map[string][]types.Action

	// The system-generated Amazon Resource Name (ARN) for the extension.
	Arn *string

	// Information about the extension.
	Description *string

	// The system-generated ID of the extension.
	Id *string

	// The extension name.
	Name *string

	// The parameters accepted by the extension. You specify parameter values when you
	// associate the extension to an AppConfig resource by using the
	// CreateExtensionAssociation API action. For Lambda extension actions, these
	// parameters are included in the Lambda request object.
	Parameters map[string]types.Parameter

	// The extension version number.
	VersionNumber int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateExtensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateExtension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateExtension{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateExtension"); err != nil {
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
	if err = addOpUpdateExtensionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExtension(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateExtension(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateExtension",
	}
}
