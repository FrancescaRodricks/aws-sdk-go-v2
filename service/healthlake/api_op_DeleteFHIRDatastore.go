// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a data store.
func (c *Client) DeleteFHIRDatastore(ctx context.Context, params *DeleteFHIRDatastoreInput, optFns ...func(*Options)) (*DeleteFHIRDatastoreOutput, error) {
	if params == nil {
		params = &DeleteFHIRDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFHIRDatastore", params, optFns, c.addOperationDeleteFHIRDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFHIRDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFHIRDatastoreInput struct {

	//  The AWS-generated ID for the data store to be deleted.
	//
	// This member is required.
	DatastoreId *string

	noSmithyDocumentSerde
}

type DeleteFHIRDatastoreOutput struct {

	// The Amazon Resource Name (ARN) that gives AWS HealthLake access permission.
	//
	// This member is required.
	DatastoreArn *string

	// The AWS endpoint for the data store the user has requested to be deleted.
	//
	// This member is required.
	DatastoreEndpoint *string

	// The AWS-generated ID for the data store to be deleted.
	//
	// This member is required.
	DatastoreId *string

	// The status of the data store that the user has requested to be deleted.
	//
	// This member is required.
	DatastoreStatus types.DatastoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFHIRDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteFHIRDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteFHIRDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteFHIRDatastore"); err != nil {
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
	if err = addOpDeleteFHIRDatastoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFHIRDatastore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFHIRDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteFHIRDatastore",
	}
}
