// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	"sync"
)

// Invokes a model at the specified endpoint to return the inference response as a
// stream. The inference stream provides the response payload incrementally as a
// series of parts. Before you can get an inference stream, you must have access to
// a model that's deployed using Amazon SageMaker hosting services, and the
// container for that model must support inference streaming.
//
// For more information that can help you use this API, see the following sections
// in the Amazon SageMaker Developer Guide:
//
//   - For information about how to add streaming support to a model, see [How Containers Serve Requests].
//
//   - For information about how to process the streaming response, see [Invoke real-time endpoints].
//
// Before you can use this operation, your IAM permissions must allow the
// sagemaker:InvokeEndpoint action. For more information about Amazon SageMaker
// actions for IAM policies, see [Actions, resources, and condition keys for Amazon SageMaker]in the IAM Service Authorization Reference.
//
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpointWithResponseStream are authenticated by using Amazon Web
// Services Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API
// Reference.
//
// [How Containers Serve Requests]: https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-code-how-containe-serves-requests
// [Invoke real-time endpoints]: https://docs.aws.amazon.com/sagemaker/latest/dg/realtime-endpoints-test-endpoints.html
// [Actions, resources, and condition keys for Amazon SageMaker]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonsagemaker.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
func (c *Client) InvokeEndpointWithResponseStream(ctx context.Context, params *InvokeEndpointWithResponseStreamInput, optFns ...func(*Options)) (*InvokeEndpointWithResponseStreamOutput, error) {
	if params == nil {
		params = &InvokeEndpointWithResponseStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeEndpointWithResponseStream", params, optFns, c.addOperationInvokeEndpointWithResponseStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeEndpointWithResponseStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeEndpointWithResponseStreamInput struct {

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model.
	//
	// For information about the format of the request body, see [Common Data Formats-Inference].
	//
	// [Common Data Formats-Inference]: https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html
	//
	// This member is required.
	Body []byte

	// The name of the endpoint that you specified when you created the endpoint using
	// the [CreateEndpoint]API.
	//
	// [CreateEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html
	//
	// This member is required.
	EndpointName *string

	// The desired MIME type of the inference response from the model container.
	Accept *string

	// The MIME type of the input data in the request body.
	ContentType *string

	// Provides additional information about a request for an inference submitted to a
	// model hosted at an Amazon SageMaker endpoint. The information is an opaque value
	// that is forwarded verbatim. You could use this value, for example, to provide an
	// ID that you can use to track a request or to provide other metadata that a
	// service endpoint was programmed to process. The value must consist of no more
	// than 1024 visible US-ASCII characters as specified in [Section 3.3.6. Field Value Components]of the Hypertext Transfer
	// Protocol (HTTP/1.1).
	//
	// The code in your model is responsible for setting or updating any custom
	// attributes in the response. If your code does not set this value in the
	// response, an empty value is returned. For example, if a custom attribute
	// represents the trace ID, your model can prepend the custom attribute with Trace
	// ID: in your post-processing function.
	//
	// This feature is currently supported in the Amazon Web Services SDKs but not in
	// the Amazon SageMaker Python SDK.
	//
	// [Section 3.3.6. Field Value Components]: https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.6
	CustomAttributes *string

	// If the endpoint hosts one or more inference components, this parameter
	// specifies the name of inference component to invoke for a streaming response.
	InferenceComponentName *string

	// An identifier that you assign to your request.
	InferenceId *string

	// If the endpoint hosts multiple containers and is configured to use direct
	// invocation, this parameter specifies the host name of the container to invoke.
	TargetContainerHostname *string

	// Specify the production variant to send the inference request to when invoking
	// an endpoint that is running two or more variants. Note that this parameter
	// overrides the default behavior for the endpoint, which is to distribute the
	// invocation traffic based on the variant weights.
	//
	// For information about how to use variant targeting to perform a/b testing, see [Test models in production]
	//
	// [Test models in production]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-ab-testing.html
	TargetVariant *string

	noSmithyDocumentSerde
}

type InvokeEndpointWithResponseStreamOutput struct {

	// The MIME type of the inference returned from the model container.
	ContentType *string

	// Provides additional information in the response about the inference returned by
	// a model hosted at an Amazon SageMaker endpoint. The information is an opaque
	// value that is forwarded verbatim. You could use this value, for example, to
	// return an ID received in the CustomAttributes header of a request or other
	// metadata that a service endpoint was programmed to produce. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in [Section 3.3.6. Field Value Components]of the
	// Hypertext Transfer Protocol (HTTP/1.1). If the customer wants the custom
	// attribute returned, the model must set the custom attribute to be included on
	// the way back.
	//
	// The code in your model is responsible for setting or updating any custom
	// attributes in the response. If your code does not set this value in the
	// response, an empty value is returned. For example, if a custom attribute
	// represents the trace ID, your model can prepend the custom attribute with Trace
	// ID: in your post-processing function.
	//
	// This feature is currently supported in the Amazon Web Services SDKs but not in
	// the Amazon SageMaker Python SDK.
	//
	// [Section 3.3.6. Field Value Components]: https://tools.ietf.org/html/rfc7230#section-3.2.6
	CustomAttributes *string

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string

	eventStream *InvokeEndpointWithResponseStreamEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *InvokeEndpointWithResponseStreamOutput) GetStream() *InvokeEndpointWithResponseStreamEventStream {
	return o.eventStream
}

func (c *Client) addOperationInvokeEndpointWithResponseStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeEndpointWithResponseStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeEndpointWithResponseStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeEndpointWithResponseStream"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamInvokeEndpointWithResponseStreamMiddleware(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpInvokeEndpointWithResponseStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeEndpointWithResponseStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeEndpointWithResponseStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeEndpointWithResponseStream",
	}
}

// InvokeEndpointWithResponseStreamEventStream provides the event stream handling for the InvokeEndpointWithResponseStream operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewInvokeEndpointWithResponseStreamEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type InvokeEndpointWithResponseStreamEventStream struct {
	// ResponseStreamReader is the EventStream reader for the ResponseStream events.
	// This value is automatically set by the SDK when the API call is made Use this
	// member when unit testing your code with the SDK to mock out the EventStream
	// Reader.
	//
	// Must not be nil.
	Reader ResponseStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewInvokeEndpointWithResponseStreamEventStream initializes an InvokeEndpointWithResponseStreamEventStream.
// This function should only be used for testing and mocking the InvokeEndpointWithResponseStreamEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewInvokeEndpointWithResponseStreamEventStream(optFns ...func(*InvokeEndpointWithResponseStreamEventStream)) *InvokeEndpointWithResponseStreamEventStream {
	es := &InvokeEndpointWithResponseStreamEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *InvokeEndpointWithResponseStreamEventStream) Events() <-chan types.ResponseStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *InvokeEndpointWithResponseStreamEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *InvokeEndpointWithResponseStreamEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *InvokeEndpointWithResponseStreamEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *InvokeEndpointWithResponseStreamEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
