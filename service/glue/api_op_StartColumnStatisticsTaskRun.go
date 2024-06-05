// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a column statistics task run, for a specified table and columns.
func (c *Client) StartColumnStatisticsTaskRun(ctx context.Context, params *StartColumnStatisticsTaskRunInput, optFns ...func(*Options)) (*StartColumnStatisticsTaskRunOutput, error) {
	if params == nil {
		params = &StartColumnStatisticsTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartColumnStatisticsTaskRun", params, optFns, c.addOperationStartColumnStatisticsTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartColumnStatisticsTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartColumnStatisticsTaskRunInput struct {

	// The name of the database where the table resides.
	//
	// This member is required.
	DatabaseName *string

	// The IAM role that the service assumes to generate statistics.
	//
	// This member is required.
	Role *string

	// The name of the table to generate statistics.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the table reside. If none is supplied, the
	// Amazon Web Services account ID is used by default.
	CatalogID *string

	// A list of the column names to generate statistics. If none is supplied, all
	// column names for the table will be used by default.
	ColumnNameList []string

	// The percentage of rows used to generate statistics. If none is supplied, the
	// entire table will be used to generate stats.
	SampleSize float64

	// Name of the security configuration that is used to encrypt CloudWatch logs for
	// the column stats task run.
	SecurityConfiguration *string

	noSmithyDocumentSerde
}

type StartColumnStatisticsTaskRunOutput struct {

	// The identifier for the column statistics task run.
	ColumnStatisticsTaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartColumnStatisticsTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartColumnStatisticsTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartColumnStatisticsTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartColumnStatisticsTaskRun"); err != nil {
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
	if err = addOpStartColumnStatisticsTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartColumnStatisticsTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartColumnStatisticsTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartColumnStatisticsTaskRun",
	}
}
