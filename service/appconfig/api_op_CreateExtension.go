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

// Creates an AppConfig extension. An extension augments your ability to inject
// logic or behavior at different points during the AppConfig workflow of creating
// or deploying a configuration.
//
// You can create your own extensions or use the Amazon Web Services authored
// extensions provided by AppConfig. For an AppConfig extension that uses Lambda,
// you must create a Lambda function to perform any computation and processing
// defined in the extension. If you plan to create custom versions of the Amazon
// Web Services authored notification extensions, you only need to specify an
// Amazon Resource Name (ARN) in the Uri field for the new extension version.
//
//   - For a custom EventBridge notification extension, enter the ARN of the
//     EventBridge default events in the Uri field.
//
//   - For a custom Amazon SNS notification extension, enter the ARN of an Amazon
//     SNS topic in the Uri field.
//
//   - For a custom Amazon SQS notification extension, enter the ARN of an Amazon
//     SQS message queue in the Uri field.
//
// For more information about extensions, see [Extending workflows] in the AppConfig User Guide.
//
// [Extending workflows]: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html
func (c *Client) CreateExtension(ctx context.Context, params *CreateExtensionInput, optFns ...func(*Options)) (*CreateExtensionOutput, error) {
	if params == nil {
		params = &CreateExtensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExtension", params, optFns, c.addOperationCreateExtensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExtensionInput struct {

	// The actions defined in the extension.
	//
	// This member is required.
	Actions map[string][]types.Action

	// A name for the extension. Each extension name in your account must be unique.
	// Extension versions use the same name.
	//
	// This member is required.
	Name *string

	// Information about the extension.
	Description *string

	// You can omit this field when you create an extension. When you create a new
	// version, specify the most recent current version number. For example, you create
	// version 3, enter 2 for this field.
	LatestVersionNumber *int32

	// The parameters accepted by the extension. You specify parameter values when you
	// associate the extension to an AppConfig resource by using the
	// CreateExtensionAssociation API action. For Lambda extension actions, these
	// parameters are included in the Lambda request object.
	Parameters map[string]types.Parameter

	// Adds one or more tags for the specified extension. Tags are metadata that help
	// you categorize resources in different ways, for example, by purpose, owner, or
	// environment. Each tag consists of a key and an optional value, both of which you
	// define.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateExtensionOutput struct {

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

func (c *Client) addOperationCreateExtensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateExtension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateExtension{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateExtension"); err != nil {
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
	if err = addOpCreateExtensionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExtension(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExtension(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateExtension",
	}
}
