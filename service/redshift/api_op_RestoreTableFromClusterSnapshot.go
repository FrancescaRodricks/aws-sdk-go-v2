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

// Creates a new table from a table in an Amazon Redshift cluster snapshot. You
// must create the new table within the Amazon Redshift cluster that the snapshot
// was taken from.
//
// You cannot use RestoreTableFromClusterSnapshot to restore a table with the same
// name as an existing table in an Amazon Redshift cluster. That is, you cannot
// overwrite an existing table in a cluster with a restored table. If you want to
// replace your original table with a new, restored table, then rename or drop your
// original table before you call RestoreTableFromClusterSnapshot . When you have
// renamed your original table, then you can pass the original name of the table as
// the NewTableName parameter value in the call to RestoreTableFromClusterSnapshot
// . This way, you can replace the original table with the table created from the
// snapshot.
//
// You can't use this operation to restore tables with [interleaved sort keys].
//
// [interleaved sort keys]: https://docs.aws.amazon.com/redshift/latest/dg/t_Sorting_data.html#t_Sorting_data-interleaved
func (c *Client) RestoreTableFromClusterSnapshot(ctx context.Context, params *RestoreTableFromClusterSnapshotInput, optFns ...func(*Options)) (*RestoreTableFromClusterSnapshotOutput, error) {
	if params == nil {
		params = &RestoreTableFromClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreTableFromClusterSnapshot", params, optFns, c.addOperationRestoreTableFromClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreTableFromClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreTableFromClusterSnapshotInput struct {

	// The identifier of the Amazon Redshift cluster to restore the table to.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the table to create as a result of the current request.
	//
	// This member is required.
	NewTableName *string

	// The identifier of the snapshot to restore the table from. This snapshot must
	// have been created from the Amazon Redshift cluster specified by the
	// ClusterIdentifier parameter.
	//
	// This member is required.
	SnapshotIdentifier *string

	// The name of the source database that contains the table to restore from.
	//
	// This member is required.
	SourceDatabaseName *string

	// The name of the source table to restore from.
	//
	// This member is required.
	SourceTableName *string

	// Indicates whether name identifiers for database, schema, and table are case
	// sensitive. If true , the names are case sensitive. If false (default), the
	// names are not case sensitive.
	EnableCaseSensitiveIdentifier *bool

	// The name of the source schema that contains the table to restore from. If you
	// do not specify a SourceSchemaName value, the default is public .
	SourceSchemaName *string

	// The name of the database to restore the table to.
	TargetDatabaseName *string

	// The name of the schema to restore the table to.
	TargetSchemaName *string

	noSmithyDocumentSerde
}

type RestoreTableFromClusterSnapshotOutput struct {

	// Describes the status of a RestoreTableFromClusterSnapshot operation.
	TableRestoreStatus *types.TableRestoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreTableFromClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreTableFromClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreTableFromClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreTableFromClusterSnapshot"); err != nil {
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
	if err = addOpRestoreTableFromClusterSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTableFromClusterSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreTableFromClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreTableFromClusterSnapshot",
	}
}
