// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a member from the resource's set of delegates.
func (c *Client) DisassociateDelegateFromResource(ctx context.Context, params *DisassociateDelegateFromResourceInput, optFns ...func(*Options)) (*DisassociateDelegateFromResourceOutput, error) {
	if params == nil {
		params = &DisassociateDelegateFromResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDelegateFromResource", params, optFns, c.addOperationDisassociateDelegateFromResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDelegateFromResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDelegateFromResourceInput struct {

	// The identifier for the member (user, group) to be removed from the resource's
	// delegates.
	//
	// The entity ID can accept UserId or GroupID, Username or Groupname, or email.
	//
	//   - Entity: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: entity@domain.tld
	//
	//   - Entity: entity
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the resource exists.
	//
	// This member is required.
	OrganizationId *string

	// The identifier of the resource from which delegates' set members are removed.
	//
	// The identifier can accept ResourceId, Resourcename, or email. The following
	// identity formats are available:
	//
	//   - Resource ID: r-0123456789a0123456789b0123456789
	//
	//   - Email address: resource@domain.tld
	//
	//   - Resource name: resource
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type DisassociateDelegateFromResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateDelegateFromResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDelegateFromResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDelegateFromResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateDelegateFromResource"); err != nil {
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
	if err = addOpDisassociateDelegateFromResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDelegateFromResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateDelegateFromResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateDelegateFromResource",
	}
}
