// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a dataset.
//
// Amazon Lookout for Metrics API actions are eventually consistent. If you do a
// read operation on a resource immediately after creating or modifying it, use
// retries to allow time for the write operation to complete.
func (c *Client) DescribeMetricSet(ctx context.Context, params *DescribeMetricSetInput, optFns ...func(*Options)) (*DescribeMetricSetOutput, error) {
	if params == nil {
		params = &DescribeMetricSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMetricSet", params, optFns, c.addOperationDescribeMetricSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMetricSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMetricSetInput struct {

	// The ARN of the dataset.
	//
	// This member is required.
	MetricSetArn *string

	noSmithyDocumentSerde
}

type DescribeMetricSetOutput struct {

	// The ARN of the detector that contains the dataset.
	AnomalyDetectorArn *string

	// The time at which the dataset was created.
	CreationTime *time.Time

	// The dimensions and their values that were used to filter the dataset.
	DimensionFilterList []types.MetricSetDimensionFilter

	// A list of the dimensions chosen for analysis.
	DimensionList []string

	// The time at which the dataset was last modified.
	LastModificationTime *time.Time

	// A list of the metrics defined by the dataset.
	MetricList []types.Metric

	// The ARN of the dataset.
	MetricSetArn *string

	// The dataset's description.
	MetricSetDescription *string

	// The interval at which the data will be analyzed for anomalies.
	MetricSetFrequency types.Frequency

	// The name of the dataset.
	MetricSetName *string

	// Contains information about the dataset's source data.
	MetricSource *types.MetricSource

	// After an interval ends, the amount of seconds that the detector waits before
	// importing data. Offset is only supported for S3, Redshift, Athena and
	// datasources.
	Offset *int32

	// Contains information about the column used for tracking time in your source
	// data.
	TimestampColumn *types.TimestampColumn

	// The time zone in which the dataset's data was recorded.
	Timezone *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMetricSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMetricSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMetricSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMetricSet"); err != nil {
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
	if err = addOpDescribeMetricSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMetricSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMetricSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMetricSet",
	}
}
