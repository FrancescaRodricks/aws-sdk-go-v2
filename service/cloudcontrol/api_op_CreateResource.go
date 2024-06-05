// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the specified resource. For more information, see [Creating a resource] in the Amazon Web
// Services Cloud Control API User Guide.
//
// After you have initiated a resource creation request, you can monitor the
// progress of your request by calling [GetResourceRequestStatus]using the RequestToken of the ProgressEvent
// type returned by CreateResource .
//
// [Creating a resource]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-create.html
// [GetResourceRequestStatus]: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
func (c *Client) CreateResource(ctx context.Context, params *CreateResourceInput, optFns ...func(*Options)) (*CreateResourceOutput, error) {
	if params == nil {
		params = &CreateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResource", params, optFns, c.addOperationCreateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceInput struct {

	// Structured data format representing the desired state of the resource,
	// consisting of that resource's properties and their desired values.
	//
	// Cloud Control API currently supports JSON as a structured data format.
	//
	// Specify the desired state as one of the following:
	//
	//   - A JSON blob
	//
	//   - A local path containing the desired state in JSON data format
	//
	// For more information, see [Composing the desired state of the resource] in the Amazon Web Services Cloud Control API User
	// Guide.
	//
	// For more information about the properties of a specific resource, refer to the
	// related topic for the resource in the [Resource and property types reference]in the CloudFormation Users Guide.
	//
	// [Resource and property types reference]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
	// [Composing the desired state of the resource]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-create.html#resource-operations-create-desiredstate
	//
	// This member is required.
	DesiredState *string

	// The name of the resource type.
	//
	// This member is required.
	TypeName *string

	// A unique identifier to ensure the idempotency of the resource request. As a
	// best practice, specify this token to ensure idempotency, so that Amazon Web
	// Services Cloud Control API can accurately distinguish between request retries
	// and new resource requests. You might retry a resource request to ensure that it
	// was successfully received.
	//
	// A client token is valid for 36 hours once used. After that, a resource request
	// with the same client token is treated as a new request.
	//
	// If you do not specify a client token, one is generated for inclusion in the
	// request.
	//
	// For more information, see [Ensuring resource operation requests are unique] in the Amazon Web Services Cloud Control API User
	// Guide.
	//
	// [Ensuring resource operation requests are unique]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-idempotency
	ClientToken *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// for Cloud Control API to use when performing this resource operation. The role
	// specified must have the permissions required for this operation. The necessary
	// permissions for each event handler are defined in the [handlers]section of the [resource type definition schema].
	//
	// If you do not specify a role, Cloud Control API uses a temporary session
	// created using your Amazon Web Services user credentials.
	//
	// For more information, see [Specifying credentials] in the Amazon Web Services Cloud Control API User
	// Guide.
	//
	// [handlers]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-handlers
	// [Specifying credentials]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-permissions
	// [resource type definition schema]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
	RoleArn *string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version.
	TypeVersionId *string

	noSmithyDocumentSerde
}

type CreateResourceOutput struct {

	// Represents the current status of the resource creation request.
	//
	// After you have initiated a resource creation request, you can monitor the
	// progress of your request by calling [GetResourceRequestStatus]using the RequestToken of the ProgressEvent
	// returned by CreateResource .
	//
	// [GetResourceRequestStatus]: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
	ProgressEvent *types.ProgressEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateResource"); err != nil {
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
	if err = addIdempotencyToken_opCreateResourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateResource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateResource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateResourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateResourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateResource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateResource",
	}
}
