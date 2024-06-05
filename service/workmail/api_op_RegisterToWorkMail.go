// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an existing and disabled user, group, or resource for WorkMail use by
// associating a mailbox and calendaring capabilities. It performs no change if the
// user, group, or resource is enabled and fails if the user, group, or resource is
// deleted. This operation results in the accumulation of costs. For more
// information, see [Pricing]. The equivalent console functionality for this operation is
// Enable.
//
// Users can either be created by calling the CreateUser API operation or they can be
// synchronized from your directory. For more information, see DeregisterFromWorkMail.
//
// [Pricing]: https://aws.amazon.com/workmail/pricing
func (c *Client) RegisterToWorkMail(ctx context.Context, params *RegisterToWorkMailInput, optFns ...func(*Options)) (*RegisterToWorkMailOutput, error) {
	if params == nil {
		params = &RegisterToWorkMailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterToWorkMail", params, optFns, c.addOperationRegisterToWorkMailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterToWorkMailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterToWorkMailInput struct {

	// The email for the user, group, or resource to be updated.
	//
	// This member is required.
	Email *string

	// The identifier for the user, group, or resource to be updated.
	//
	// The identifier can accept UserId, ResourceId, or GroupId, or Username,
	// Resourcename, or Groupname. The following identity formats are available:
	//
	//   - Entity ID: 12345678-1234-1234-1234-123456789012,
	//   r-0123456789a0123456789b0123456789, or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Entity name: entity
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the user, group, or resource
	// exists.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type RegisterToWorkMailOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterToWorkMailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterToWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterToWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterToWorkMail"); err != nil {
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
	if err = addOpRegisterToWorkMailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterToWorkMail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterToWorkMail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterToWorkMail",
	}
}
