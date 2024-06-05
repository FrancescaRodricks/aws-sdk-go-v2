// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Converts a recovery point to a snapshot. For more information about recovery
// points and snapshots, see [Working with snapshots and recovery points].
//
// [Working with snapshots and recovery points]: https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-snapshots-recovery.html
func (c *Client) ConvertRecoveryPointToSnapshot(ctx context.Context, params *ConvertRecoveryPointToSnapshotInput, optFns ...func(*Options)) (*ConvertRecoveryPointToSnapshotOutput, error) {
	if params == nil {
		params = &ConvertRecoveryPointToSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConvertRecoveryPointToSnapshot", params, optFns, c.addOperationConvertRecoveryPointToSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConvertRecoveryPointToSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConvertRecoveryPointToSnapshotInput struct {

	// The unique identifier of the recovery point.
	//
	// This member is required.
	RecoveryPointId *string

	// The name of the snapshot.
	//
	// This member is required.
	SnapshotName *string

	// How long to retain the snapshot.
	RetentionPeriod *int32

	// An array of [Tag objects] to associate with the created snapshot.
	//
	// [Tag objects]: https://docs.aws.amazon.com/redshift-serverless/latest/APIReference/API_Tag.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ConvertRecoveryPointToSnapshotOutput struct {

	// The snapshot converted from the recovery point.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConvertRecoveryPointToSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConvertRecoveryPointToSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConvertRecoveryPointToSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ConvertRecoveryPointToSnapshot"); err != nil {
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
	if err = addOpConvertRecoveryPointToSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConvertRecoveryPointToSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConvertRecoveryPointToSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConvertRecoveryPointToSnapshot",
	}
}
