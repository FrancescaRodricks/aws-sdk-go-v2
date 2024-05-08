// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes multiple tables at once.
//
// After completing this operation, you no longer have access to the table
// versions and partitions that belong to the deleted table. Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion of the
// service.
//
// To ensure the immediate deletion of all related resources, before calling
// BatchDeleteTable , use DeleteTableVersion or BatchDeleteTableVersion , and
// DeletePartition or BatchDeletePartition , to delete any resources that belong to
// the table.
func (c *Client) BatchDeleteTable(ctx context.Context, params *BatchDeleteTableInput, optFns ...func(*Options)) (*BatchDeleteTableOutput, error) {
	if params == nil {
		params = &BatchDeleteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteTable", params, optFns, c.addOperationBatchDeleteTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableInput struct {

	// The name of the catalog database in which the tables to delete reside. For Hive
	// compatibility, this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// A list of the table to delete.
	//
	// This member is required.
	TablesToDelete []string

	// The ID of the Data Catalog where the table resides. If none is provided, the
	// Amazon Web Services account ID is used by default.
	CatalogId *string

	// The transaction ID at which to delete the table contents.
	TransactionId *string

	noSmithyDocumentSerde
}

type BatchDeleteTableOutput struct {

	// A list of errors encountered in attempting to delete the specified tables.
	Errors []types.TableError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchDeleteTable"); err != nil {
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
	if err = addOpBatchDeleteTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchDeleteTable",
	}
}
