// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Retrieves information about the specified IAM user, including the user's
// creation date, path, unique ID, and ARN.
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID used to sign the request to this
// operation.
func (c *Client) GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) {
	if params == nil {
		params = &GetUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUser", params, optFns, c.addOperationGetUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserInput struct {

	// The name of the user to get information about.
	//
	// This parameter is optional. If it is not included, it defaults to the user
	// making the request. This parameter allows (through its [regex pattern]) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	UserName *string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetUser request.
type GetUserOutput struct {

	// A structure containing details about the IAM user.
	//
	// Due to a service issue, password last used data does not include password use
	// from May 3, 2018 22:50 PDT to May 23, 2018 14:08 PDT. This affects [last sign-in]dates shown
	// in the IAM console and password last used dates in the [IAM credential report], and returned by this
	// operation. If users signed in during the affected time, the password last used
	// date that is returned is the date the user last signed in before May 3, 2018.
	// For users that signed in after May 23, 2018 14:08 PDT, the returned password
	// last used date is accurate.
	//
	// You can use password last used information to identify unused credentials for
	// deletion. For example, you might delete users who did not sign in to Amazon Web
	// Services in the last 90 days. In cases like this, we recommend that you adjust
	// your evaluation window to include dates after May 23, 2018. Alternatively, if
	// your users use access keys to access Amazon Web Services programmatically you
	// can refer to access key last used information because it is accurate for all
	// dates.
	//
	// [IAM credential report]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html
	// [last sign-in]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_finding-unused.html
	//
	// This member is required.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUser"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUser(options.Region), middleware.Before); err != nil {
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

// GetUserAPIClient is a client that implements the GetUser operation.
type GetUserAPIClient interface {
	GetUser(context.Context, *GetUserInput, ...func(*Options)) (*GetUserOutput, error)
}

var _ GetUserAPIClient = (*Client)(nil)

// UserExistsWaiterOptions are waiter options for UserExistsWaiter
type UserExistsWaiterOptions struct {

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
	// UserExistsWaiter will use default minimum delay of 1 seconds. Note that MinDelay
	// must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, UserExistsWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *GetUserInput, *GetUserOutput, error) (bool, error)
}

// UserExistsWaiter defines the waiters for UserExists
type UserExistsWaiter struct {
	client GetUserAPIClient

	options UserExistsWaiterOptions
}

// NewUserExistsWaiter constructs a UserExistsWaiter.
func NewUserExistsWaiter(client GetUserAPIClient, optFns ...func(*UserExistsWaiterOptions)) *UserExistsWaiter {
	options := UserExistsWaiterOptions{}
	options.MinDelay = 1 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = userExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &UserExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for UserExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *UserExistsWaiter) Wait(ctx context.Context, params *GetUserInput, maxWaitDur time.Duration, optFns ...func(*UserExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for UserExists waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *UserExistsWaiter) WaitForOutput(ctx context.Context, params *GetUserInput, maxWaitDur time.Duration, optFns ...func(*UserExistsWaiterOptions)) (*GetUserOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
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

		out, err := w.client.GetUser(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for UserExists waiter")
}

func userExistsStateRetryable(ctx context.Context, input *GetUserInput, output *GetUserOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "NoSuchEntity" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUser",
	}
}
