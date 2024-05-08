// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes permissions granted to a member (user or group).
func (c *Client) DeleteMailboxPermissions(ctx context.Context, params *DeleteMailboxPermissionsInput, optFns ...func(*Options)) (*DeleteMailboxPermissionsOutput, error) {
	if params == nil {
		params = &DeleteMailboxPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMailboxPermissions", params, optFns, c.addOperationDeleteMailboxPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMailboxPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMailboxPermissionsInput struct {

	// The identifier of the entity that owns the mailbox.
	//
	// The identifier can be UserId or Group Id, Username or Groupname, or email.
	//
	//   - Entity ID: 12345678-1234-1234-1234-123456789012,
	//   r-0123456789a0123456789b0123456789, or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: entity@domain.tld
	//
	//   - Entity name: entity
	//
	// This member is required.
	EntityId *string

	// The identifier of the entity for which to delete granted permissions.
	//
	// The identifier can be UserId, ResourceID, or Group Id, Username or Groupname,
	// or email.
	//
	//   - Grantee ID:
	//   12345678-1234-1234-1234-123456789012,r-0123456789a0123456789b0123456789, or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: grantee@domain.tld
	//
	//   - Grantee name: grantee
	//
	// This member is required.
	GranteeId *string

	// The identifier of the organization under which the member (user or group)
	// exists.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type DeleteMailboxPermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMailboxPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMailboxPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMailboxPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteMailboxPermissions"); err != nil {
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
	if err = addOpDeleteMailboxPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMailboxPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMailboxPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteMailboxPermissions",
	}
}
