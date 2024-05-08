// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides a description of the retraining scheduler, including information such
// as the model name and retraining parameters.
func (c *Client) DescribeRetrainingScheduler(ctx context.Context, params *DescribeRetrainingSchedulerInput, optFns ...func(*Options)) (*DescribeRetrainingSchedulerOutput, error) {
	if params == nil {
		params = &DescribeRetrainingSchedulerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRetrainingScheduler", params, optFns, c.addOperationDescribeRetrainingSchedulerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRetrainingSchedulerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRetrainingSchedulerInput struct {

	// The name of the model that the retraining scheduler is attached to.
	//
	// This member is required.
	ModelName *string

	noSmithyDocumentSerde
}

type DescribeRetrainingSchedulerOutput struct {

	// Indicates the time and date at which the retraining scheduler was created.
	CreatedAt *time.Time

	// The number of past days of data used for retraining.
	LookbackWindow *string

	// The ARN of the model that the retraining scheduler is attached to.
	ModelArn *string

	// The name of the model that the retraining scheduler is attached to.
	ModelName *string

	// Indicates how the service uses new models. In MANAGED mode, new models are used
	// for inference if they have better performance than the current model. In MANUAL
	// mode, the new models are not used until they are [manually activated].
	//
	// [manually activated]: https://docs.aws.amazon.com/lookout-for-equipment/latest/ug/versioning-model.html#model-activation
	PromoteMode types.ModelPromoteMode

	// The frequency at which the model retraining is set. This follows the [ISO 8601]
	// guidelines.
	//
	// [ISO 8601]: https://en.wikipedia.org/wiki/ISO_8601#Durations
	RetrainingFrequency *string

	// The start date for the retraining scheduler. Lookout for Equipment truncates
	// the time you provide to the nearest UTC day.
	RetrainingStartDate *time.Time

	// The status of the retraining scheduler.
	Status types.RetrainingSchedulerStatus

	// Indicates the time and date at which the retraining scheduler was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRetrainingSchedulerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRetrainingScheduler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRetrainingScheduler{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRetrainingScheduler"); err != nil {
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
	if err = addOpDescribeRetrainingSchedulerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRetrainingScheduler(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRetrainingScheduler(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRetrainingScheduler",
	}
}
