// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a progress update stream, including all of its tasks, which was
// previously created as an AWS resource used for access control. This API has the
// following traits:
//
//   - The only parameter needed for DeleteProgressUpdateStream is the stream name
//     (same as a CreateProgressUpdateStream call).
//
//   - The call will return, and a background process will asynchronously delete
//     the stream and all of its resources (tasks, associated resources, resource
//     attributes, created artifacts).
//
//   - If the stream takes time to be deleted, it might still show up on a
//     ListProgressUpdateStreams call.
//
//   - CreateProgressUpdateStream , ImportMigrationTask , NotifyMigrationTaskState
//     , and all Associate[*] APIs related to the tasks belonging to the stream will
//     throw "InvalidInputException" if the stream of the same name is in the process
//     of being deleted.
//
//   - Once the stream and all of its resources are deleted,
//     CreateProgressUpdateStream for a stream of the same name will succeed, and
//     that stream will be an entirely new logical resource (without any resources
//     associated with the old stream).
func (c *Client) DeleteProgressUpdateStream(ctx context.Context, params *DeleteProgressUpdateStreamInput, optFns ...func(*Options)) (*DeleteProgressUpdateStreamOutput, error) {
	if params == nil {
		params = &DeleteProgressUpdateStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProgressUpdateStream", params, optFns, c.addOperationDeleteProgressUpdateStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProgressUpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProgressUpdateStreamInput struct {

	// The name of the ProgressUpdateStream. Do not store personal data in this field.
	//
	// This member is required.
	ProgressUpdateStreamName *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun bool

	noSmithyDocumentSerde
}

type DeleteProgressUpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteProgressUpdateStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProgressUpdateStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProgressUpdateStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteProgressUpdateStream"); err != nil {
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
	if err = addOpDeleteProgressUpdateStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProgressUpdateStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteProgressUpdateStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteProgressUpdateStream",
	}
}
