// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Groups the aggregated values that match the query into percentile groupings.
// The default percentile groupings are: 1,5,25,50,75,95,99, although you can
// specify your own when you call GetPercentiles . This function returns a value
// for each percentile group specified (or the default percentile groupings). The
// percentile group "1" contains the aggregated field value that occurs in
// approximately one percent of the values that match the query. The percentile
// group "5" contains the aggregated field value that occurs in approximately five
// percent of the values that match the query, and so on. The result is an
// approximation, the more values that match the query, the more accurate the
// percentile values.
//
// Requires permission to access the [GetPercentiles] action.
//
// [GetPercentiles]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) GetPercentiles(ctx context.Context, params *GetPercentilesInput, optFns ...func(*Options)) (*GetPercentilesOutput, error) {
	if params == nil {
		params = &GetPercentilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPercentiles", params, optFns, c.addOperationGetPercentilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPercentilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPercentilesInput struct {

	// The search query string.
	//
	// This member is required.
	QueryString *string

	// The field to aggregate.
	AggregationField *string

	// The name of the index to search.
	IndexName *string

	// The percentile groups returned.
	Percents []float64

	// The query version.
	QueryVersion *string

	noSmithyDocumentSerde
}

type GetPercentilesOutput struct {

	// The percentile values of the aggregated fields.
	Percentiles []types.PercentPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPercentilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPercentiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPercentiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPercentiles"); err != nil {
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
	if err = addOpGetPercentilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPercentiles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPercentiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPercentiles",
	}
}
