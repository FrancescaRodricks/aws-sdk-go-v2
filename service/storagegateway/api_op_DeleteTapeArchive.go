// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified virtual tape from the virtual tape shelf (VTS). This
// operation is only supported in the tape gateway type.
func (c *Client) DeleteTapeArchive(ctx context.Context, params *DeleteTapeArchiveInput, optFns ...func(*Options)) (*DeleteTapeArchiveOutput, error) {
	if params == nil {
		params = &DeleteTapeArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTapeArchive", params, optFns, c.addOperationDeleteTapeArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTapeArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteTapeArchiveInput
type DeleteTapeArchiveInput struct {

	// The Amazon Resource Name (ARN) of the virtual tape to delete from the virtual
	// tape shelf (VTS).
	//
	// This member is required.
	TapeARN *string

	// Set to TRUE to delete an archived tape that belongs to a custom pool with tape
	// retention lock. Only archived tapes with tape retention lock set to governance
	// can be deleted. Archived tapes with tape retention lock set to compliance can't
	// be deleted.
	BypassGovernanceRetention bool

	noSmithyDocumentSerde
}

// DeleteTapeArchiveOutput
type DeleteTapeArchiveOutput struct {

	// The Amazon Resource Name (ARN) of the virtual tape that was deleted from the
	// virtual tape shelf (VTS).
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTapeArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTapeArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTapeArchive{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteTapeArchive"); err != nil {
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
	if err = addOpDeleteTapeArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTapeArchive(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTapeArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteTapeArchive",
	}
}
