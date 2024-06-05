// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the results of a specific Face Liveness session. It requires the
// sessionId as input, which was created using CreateFaceLivenessSession . Returns
// the corresponding Face Liveness confidence score, a reference image that
// includes a face bounding box, and audit images that also contain face bounding
// boxes. The Face Liveness confidence score ranges from 0 to 100.
//
// The number of audit images returned by GetFaceLivenessSessionResults is defined
// by the AuditImagesLimit paramater when calling CreateFaceLivenessSession .
// Reference images are always returned when possible.
func (c *Client) GetFaceLivenessSessionResults(ctx context.Context, params *GetFaceLivenessSessionResultsInput, optFns ...func(*Options)) (*GetFaceLivenessSessionResultsOutput, error) {
	if params == nil {
		params = &GetFaceLivenessSessionResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFaceLivenessSessionResults", params, optFns, c.addOperationGetFaceLivenessSessionResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFaceLivenessSessionResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFaceLivenessSessionResultsInput struct {

	// A unique 128-bit UUID. This is used to uniquely identify the session and also
	// acts as an idempotency token for all operations associated with the session.
	//
	// This member is required.
	SessionId *string

	noSmithyDocumentSerde
}

type GetFaceLivenessSessionResultsOutput struct {

	// The sessionId for which this request was called.
	//
	// This member is required.
	SessionId *string

	// Represents a status corresponding to the state of the session. Possible
	// statuses are: CREATED, IN_PROGRESS, SUCCEEDED, FAILED, EXPIRED.
	//
	// This member is required.
	Status types.LivenessSessionStatus

	// A set of images from the Face Liveness video that can be used for audit
	// purposes. It includes a bounding box of the face and the Base64-encoded bytes
	// that return an image. If the CreateFaceLivenessSession request included an
	// OutputConfig argument, the image will be uploaded to an S3Object specified in
	// the output configuration. If no Amazon S3 bucket is defined, raw bytes are sent
	// instead.
	AuditImages []types.AuditImage

	// Probabalistic confidence score for if the person in the given video was live,
	// represented as a float value between 0 to 100.
	Confidence *float32

	// A high-quality image from the Face Liveness video that can be used for face
	// comparison or search. It includes a bounding box of the face and the
	// Base64-encoded bytes that return an image. If the CreateFaceLivenessSession
	// request included an OutputConfig argument, the image will be uploaded to an
	// S3Object specified in the output configuration. In case the reference image is
	// not returned, it's recommended to retry the Liveness check.
	ReferenceImage *types.AuditImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFaceLivenessSessionResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFaceLivenessSessionResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFaceLivenessSessionResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFaceLivenessSessionResults"); err != nil {
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
	if err = addOpGetFaceLivenessSessionResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFaceLivenessSessionResults(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFaceLivenessSessionResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFaceLivenessSessionResults",
	}
}
