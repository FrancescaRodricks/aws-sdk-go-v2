// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds permissions to the resource-based policy of a version of an [Lambda layer]. Use this
// action to grant layer usage permission to other accounts. You can grant
// permission to a single account, all accounts in an organization, or all Amazon
// Web Services accounts.
//
// To revoke permission, call RemoveLayerVersionPermission with the statement ID that you specified when you
// added it.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func (c *Client) AddLayerVersionPermission(ctx context.Context, params *AddLayerVersionPermissionInput, optFns ...func(*Options)) (*AddLayerVersionPermissionOutput, error) {
	if params == nil {
		params = &AddLayerVersionPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddLayerVersionPermission", params, optFns, c.addOperationAddLayerVersionPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddLayerVersionPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddLayerVersionPermissionInput struct {

	// The API action that grants access to the layer. For example,
	// lambda:GetLayerVersion .
	//
	// This member is required.
	Action *string

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// This member is required.
	LayerName *string

	// An account ID, or * to grant layer usage permission to all accounts in an
	// organization, or all Amazon Web Services accounts (if organizationId is not
	// specified). For the last case, make sure that you really do want all Amazon Web
	// Services accounts to have usage permission to this layer.
	//
	// This member is required.
	Principal *string

	// An identifier that distinguishes the policy from others on the same layer
	// version.
	//
	// This member is required.
	StatementId *string

	// The version number.
	//
	// This member is required.
	VersionNumber *int64

	// With the principal set to * , grant permission to all accounts in the specified
	// organization.
	OrganizationId *string

	// Only update the policy if the revision ID matches the ID specified. Use this
	// option to avoid modifying a policy that has changed since you last read it.
	RevisionId *string

	noSmithyDocumentSerde
}

type AddLayerVersionPermissionOutput struct {

	// A unique identifier for the current revision of the policy.
	RevisionId *string

	// The permission statement.
	Statement *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddLayerVersionPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddLayerVersionPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddLayerVersionPermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddLayerVersionPermission"); err != nil {
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
	if err = addOpAddLayerVersionPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddLayerVersionPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddLayerVersionPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddLayerVersionPermission",
	}
}
