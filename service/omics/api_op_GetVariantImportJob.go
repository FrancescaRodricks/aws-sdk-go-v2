// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about a variant import job.
func (c *Client) GetVariantImportJob(ctx context.Context, params *GetVariantImportJobInput, optFns ...func(*Options)) (*GetVariantImportJobOutput, error) {
	if params == nil {
		params = &GetVariantImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVariantImportJob", params, optFns, c.addOperationGetVariantImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVariantImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVariantImportJobInput struct {

	// The job's ID.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetVariantImportJobOutput struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's destination variant store.
	//
	// This member is required.
	DestinationName *string

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's items.
	//
	// This member is required.
	Items []types.VariantImportItemDetail

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's left normalization setting.
	//
	// This member is required.
	RunLeftNormalization bool

	// The job's status.
	//
	// This member is required.
	Status types.JobStatus

	// The job's status message.
	//
	// This member is required.
	StatusMessage *string

	// When the job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The annotation schema generated by the parsed annotation data.
	AnnotationFields map[string]string

	// When the job completed.
	CompletionTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVariantImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVariantImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVariantImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetVariantImportJob"); err != nil {
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
	if err = addEndpointPrefix_opGetVariantImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetVariantImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVariantImportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetVariantImportJobMiddleware struct {
}

func (*endpointPrefix_opGetVariantImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetVariantImportJobMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetVariantImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetVariantImportJobMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetVariantImportJobAPIClient is a client that implements the
// GetVariantImportJob operation.
type GetVariantImportJobAPIClient interface {
	GetVariantImportJob(context.Context, *GetVariantImportJobInput, ...func(*Options)) (*GetVariantImportJobOutput, error)
}

var _ GetVariantImportJobAPIClient = (*Client)(nil)

// VariantImportJobCreatedWaiterOptions are waiter options for
// VariantImportJobCreatedWaiter
type VariantImportJobCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VariantImportJobCreatedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, VariantImportJobCreatedWaiter will use default max delay of 600
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetVariantImportJobInput, *GetVariantImportJobOutput, error) (bool, error)
}

// VariantImportJobCreatedWaiter defines the waiters for VariantImportJobCreated
type VariantImportJobCreatedWaiter struct {
	client GetVariantImportJobAPIClient

	options VariantImportJobCreatedWaiterOptions
}

// NewVariantImportJobCreatedWaiter constructs a VariantImportJobCreatedWaiter.
func NewVariantImportJobCreatedWaiter(client GetVariantImportJobAPIClient, optFns ...func(*VariantImportJobCreatedWaiterOptions)) *VariantImportJobCreatedWaiter {
	options := VariantImportJobCreatedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = variantImportJobCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VariantImportJobCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VariantImportJobCreated waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *VariantImportJobCreatedWaiter) Wait(ctx context.Context, params *GetVariantImportJobInput, maxWaitDur time.Duration, optFns ...func(*VariantImportJobCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for VariantImportJobCreated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *VariantImportJobCreatedWaiter) WaitForOutput(ctx context.Context, params *GetVariantImportJobInput, maxWaitDur time.Duration, optFns ...func(*VariantImportJobCreatedWaiterOptions)) (*GetVariantImportJobOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetVariantImportJob(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for VariantImportJobCreated waiter")
}

func variantImportJobCreatedStateRetryable(ctx context.Context, input *GetVariantImportJobInput, output *GetVariantImportJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUBMITTED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "IN_PROGRESS"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCELLED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetVariantImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetVariantImportJob",
	}
}
