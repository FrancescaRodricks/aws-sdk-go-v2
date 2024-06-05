// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the number of days to retain snapshots in the destination Amazon Web
// Services Region after they are copied from the source Amazon Web Services
// Region. By default, this operation only changes the retention period of copied
// automated snapshots. The retention periods for both new and existing copied
// automated snapshots are updated with the new retention period. You can set the
// manual option to change only the retention periods of copied manual snapshots.
// If you set this option, only newly copied manual snapshots have the new
// retention period.
func (c *Client) ModifySnapshotCopyRetentionPeriod(ctx context.Context, params *ModifySnapshotCopyRetentionPeriodInput, optFns ...func(*Options)) (*ModifySnapshotCopyRetentionPeriodOutput, error) {
	if params == nil {
		params = &ModifySnapshotCopyRetentionPeriodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifySnapshotCopyRetentionPeriod", params, optFns, c.addOperationModifySnapshotCopyRetentionPeriodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifySnapshotCopyRetentionPeriodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifySnapshotCopyRetentionPeriodInput struct {

	// The unique identifier of the cluster for which you want to change the retention
	// period for either automated or manual snapshots that are copied to a destination
	// Amazon Web Services Region.
	//
	// Constraints: Must be the valid name of an existing cluster that has
	// cross-region snapshot copy enabled.
	//
	// This member is required.
	ClusterIdentifier *string

	// The number of days to retain automated snapshots in the destination Amazon Web
	// Services Region after they are copied from the source Amazon Web Services
	// Region.
	//
	// By default, this only changes the retention period of copied automated
	// snapshots.
	//
	// If you decrease the retention period for automated snapshots that are copied to
	// a destination Amazon Web Services Region, Amazon Redshift deletes any existing
	// automated snapshots that were copied to the destination Amazon Web Services
	// Region and that fall outside of the new retention period.
	//
	// Constraints: Must be at least 1 and no more than 35 for automated snapshots.
	//
	// If you specify the manual option, only newly copied manual snapshots will have
	// the new retention period.
	//
	// If you specify the value of -1 newly copied manual snapshots are retained
	// indefinitely.
	//
	// Constraints: The number of days must be either -1 or an integer between 1 and
	// 3,653 for manual snapshots.
	//
	// This member is required.
	RetentionPeriod *int32

	// Indicates whether to apply the snapshot retention period to newly copied manual
	// snapshots instead of automated snapshots.
	Manual *bool

	noSmithyDocumentSerde
}

type ModifySnapshotCopyRetentionPeriodOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifySnapshotCopyRetentionPeriodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifySnapshotCopyRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifySnapshotCopyRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifySnapshotCopyRetentionPeriod"); err != nil {
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
	if err = addOpModifySnapshotCopyRetentionPeriodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifySnapshotCopyRetentionPeriod(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifySnapshotCopyRetentionPeriod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifySnapshotCopyRetentionPeriod",
	}
}
