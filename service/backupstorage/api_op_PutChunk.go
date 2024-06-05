// Code generated by smithy-go-codegen DO NOT EDIT.

package backupstorage

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backupstorage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Upload chunk.
func (c *Client) PutChunk(ctx context.Context, params *PutChunkInput, optFns ...func(*Options)) (*PutChunkOutput, error) {
	if params == nil {
		params = &PutChunkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutChunk", params, optFns, c.addOperationPutChunkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutChunkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutChunkInput struct {

	// Backup job Id for the in-progress backup.
	//
	// This member is required.
	BackupJobId *string

	// Data checksum
	//
	// This member is required.
	Checksum *string

	// Checksum algorithm
	//
	// This member is required.
	ChecksumAlgorithm types.DataChecksumAlgorithm

	// Describes this chunk's position relative to the other chunks
	//
	// This member is required.
	ChunkIndex *int64

	// Data to be uploaded
	//
	// This member is required.
	Data io.Reader

	// Data length
	//
	// This member is required.
	Length int64

	// Upload Id for the in-progress upload.
	//
	// This member is required.
	UploadId *string

	noSmithyDocumentSerde
}

type PutChunkOutput struct {

	// Chunk checksum
	//
	// This member is required.
	ChunkChecksum *string

	// Checksum algorithm
	//
	// This member is required.
	ChunkChecksumAlgorithm types.DataChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutChunkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutChunk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutChunk{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutChunk"); err != nil {
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
	if err = addUnsignedPayload(stack); err != nil {
		return err
	}
	if err = addContentSHA256Header(stack); err != nil {
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
	if err = addOpPutChunkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutChunk(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutChunk(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutChunk",
	}
}
