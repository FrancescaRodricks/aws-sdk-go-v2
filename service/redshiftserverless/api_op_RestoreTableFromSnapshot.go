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

// Restores a table from a snapshot to your Amazon Redshift Serverless instance.
// You can't use this operation to restore tables with [interleaved sort keys].
//
// [interleaved sort keys]: https://docs.aws.amazon.com/redshift/latest/dg/t_Sorting_data.html#t_Sorting_data-interleaved
func (c *Client) RestoreTableFromSnapshot(ctx context.Context, params *RestoreTableFromSnapshotInput, optFns ...func(*Options)) (*RestoreTableFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreTableFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreTableFromSnapshot", params, optFns, c.addOperationRestoreTableFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreTableFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreTableFromSnapshotInput struct {

	// The namespace of the snapshot to restore from.
	//
	// This member is required.
	NamespaceName *string

	// The name of the table to create from the restore operation.
	//
	// This member is required.
	NewTableName *string

	// The name of the snapshot to restore the table from.
	//
	// This member is required.
	SnapshotName *string

	// The name of the source database that contains the table being restored.
	//
	// This member is required.
	SourceDatabaseName *string

	// The name of the source table being restored.
	//
	// This member is required.
	SourceTableName *string

	// The workgroup to restore the table to.
	//
	// This member is required.
	WorkgroupName *string

	// Indicates whether name identifiers for database, schema, and table are case
	// sensitive. If true, the names are case sensitive. If false, the names are not
	// case sensitive. The default is false.
	ActivateCaseSensitiveIdentifier *bool

	// The name of the source schema that contains the table being restored.
	SourceSchemaName *string

	// The name of the database to restore the table to.
	TargetDatabaseName *string

	// The name of the schema to restore the table to.
	TargetSchemaName *string

	noSmithyDocumentSerde
}

type RestoreTableFromSnapshotOutput struct {

	// The TableRestoreStatus object that contains the status of the restore operation.
	TableRestoreStatus *types.TableRestoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreTableFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreTableFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreTableFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreTableFromSnapshot"); err != nil {
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
	if err = addOpRestoreTableFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTableFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreTableFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreTableFromSnapshot",
	}
}
