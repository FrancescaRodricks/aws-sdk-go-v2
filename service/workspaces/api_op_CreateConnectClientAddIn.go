// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a client-add-in for Amazon Connect within a directory. You can create
// only one Amazon Connect client add-in within a directory.
//
// This client add-in allows WorkSpaces users to seamlessly connect to Amazon
// Connect.
func (c *Client) CreateConnectClientAddIn(ctx context.Context, params *CreateConnectClientAddInInput, optFns ...func(*Options)) (*CreateConnectClientAddInOutput, error) {
	if params == nil {
		params = &CreateConnectClientAddInInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnectClientAddIn", params, optFns, c.addOperationCreateConnectClientAddInMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectClientAddInOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectClientAddInInput struct {

	// The name of the client add-in.
	//
	// This member is required.
	Name *string

	// The directory identifier for which to configure the client add-in.
	//
	// This member is required.
	ResourceId *string

	// The endpoint URL of the Amazon Connect client add-in.
	//
	// This member is required.
	URL *string

	noSmithyDocumentSerde
}

type CreateConnectClientAddInOutput struct {

	// The client add-in identifier.
	AddInId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectClientAddInMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConnectClientAddIn{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConnectClientAddIn{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConnectClientAddIn"); err != nil {
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
	if err = addOpCreateConnectClientAddInValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnectClientAddIn(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnectClientAddIn(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConnectClientAddIn",
	}
}
