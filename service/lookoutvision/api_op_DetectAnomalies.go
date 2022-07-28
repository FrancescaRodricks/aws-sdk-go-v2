// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Detects anomalies in an image that you supply. The response from DetectAnomalies
// includes a boolean prediction that the image contains one or more anomalies and
// a confidence value for the prediction. If the model is an image segmentation
// model, the response also includes segmentation information for each type of
// anomaly found in the image. Before calling DetectAnomalies, you must first start
// your model with the StartModel operation. You are charged for the amount of
// time, in minutes, that a model runs and for the number of anomaly detection
// units that your model uses. If you are not using a model, use the StopModel
// operation to stop your model. For more information, see Detecting anomalies in
// an image in the Amazon Lookout for Vision developer guide. This operation
// requires permissions to perform the lookoutvision:DetectAnomalies operation.
func (c *Client) DetectAnomalies(ctx context.Context, params *DetectAnomaliesInput, optFns ...func(*Options)) (*DetectAnomaliesOutput, error) {
	if params == nil {
		params = &DetectAnomaliesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectAnomalies", params, optFns, c.addOperationDetectAnomaliesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectAnomaliesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectAnomaliesInput struct {

	// The unencrypted image bytes that you want to analyze.
	//
	// This member is required.
	Body io.Reader

	// The type of the image passed in Body. Valid values are image/png (PNG format
	// images) and image/jpeg (JPG format images).
	//
	// This member is required.
	ContentType *string

	// The version of the model that you want to use.
	//
	// This member is required.
	ModelVersion *string

	// The name of the project that contains the model version that you want to use.
	//
	// This member is required.
	ProjectName *string

	noSmithyDocumentSerde
}

type DetectAnomaliesOutput struct {

	// The results of the DetectAnomalies operation.
	DetectAnomalyResult *types.DetectAnomalyResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectAnomaliesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDetectAnomalies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDetectAnomalies{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDetectAnomaliesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectAnomalies(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDetectAnomalies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "DetectAnomalies",
	}
}
