// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an alert for an anomaly detector.
func (c *Client) CreateAlert(ctx context.Context, params *CreateAlertInput, optFns ...func(*Options)) (*CreateAlertOutput, error) {
	if params == nil {
		params = &CreateAlertInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAlert", params, optFns, c.addOperationCreateAlertMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAlertOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAlertInput struct {

	// Action that will be triggered when there is an alert.
	//
	// This member is required.
	Action *types.Action

	// The name of the alert.
	//
	// This member is required.
	AlertName *string

	// An integer from 0 to 100 specifying the alert sensitivity threshold.
	//
	// This member is required.
	AlertSensitivityThreshold int32

	// The ARN of the detector to which the alert is attached.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// A description of the alert.
	AlertDescription *string

	// A list of tags
	// (https://docs.aws.amazon.com/lookoutmetrics/latest/dev/detectors-tags.html) to
	// apply to the alert.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAlertOutput struct {

	// The ARN of the alert.
	AlertArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAlertMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAlert{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAlert{}, middleware.After)
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
	if err = addOpCreateAlertValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlert(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAlert(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "CreateAlert",
	}
}