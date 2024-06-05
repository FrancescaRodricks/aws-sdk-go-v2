// Code generated by smithy-go-codegen DO NOT EDIT.

package kendraranking

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a rescore execution plan. A rescore execution plan is an
// Amazon Kendra Intelligent Ranking resource used for provisioning the Rescore
// API.
func (c *Client) DescribeRescoreExecutionPlan(ctx context.Context, params *DescribeRescoreExecutionPlanInput, optFns ...func(*Options)) (*DescribeRescoreExecutionPlanOutput, error) {
	if params == nil {
		params = &DescribeRescoreExecutionPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRescoreExecutionPlan", params, optFns, c.addOperationDescribeRescoreExecutionPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRescoreExecutionPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRescoreExecutionPlanInput struct {

	// The identifier of the rescore execution plan that you want to get information
	// on.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribeRescoreExecutionPlanOutput struct {

	// The Amazon Resource Name (ARN) of the rescore execution plan.
	Arn *string

	// The capacity units set for the rescore execution plan. A capacity of zero
	// indicates that the rescore execution plan is using the default capacity. For
	// more information on the default capacity and additional capacity units, see [Adjusting capacity].
	//
	// [Adjusting capacity]: https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html
	CapacityUnits *types.CapacityUnitsConfiguration

	// The Unix timestamp of when the rescore execution plan was created.
	CreatedAt *time.Time

	// The description for the rescore execution plan.
	Description *string

	// When the Status field value is FAILED , the ErrorMessage field contains a
	// message that explains why.
	ErrorMessage *string

	// The identifier of the rescore execution plan.
	Id *string

	// The name for the rescore execution plan.
	Name *string

	// The current status of the rescore execution plan. When the value is ACTIVE , the
	// rescore execution plan is ready for use. If the Status field value is FAILED ,
	// the ErrorMessage field contains a message that explains why.
	Status types.RescoreExecutionPlanStatus

	// The Unix timestamp of when the rescore execution plan was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRescoreExecutionPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRescoreExecutionPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRescoreExecutionPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRescoreExecutionPlan"); err != nil {
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
	if err = addOpDescribeRescoreExecutionPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRescoreExecutionPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRescoreExecutionPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRescoreExecutionPlan",
	}
}
