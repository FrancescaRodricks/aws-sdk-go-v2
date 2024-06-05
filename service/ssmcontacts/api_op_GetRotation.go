// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about an on-call rotation.
func (c *Client) GetRotation(ctx context.Context, params *GetRotationInput, optFns ...func(*Options)) (*GetRotationOutput, error) {
	if params == nil {
		params = &GetRotationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRotation", params, optFns, c.addOperationGetRotationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRotationInput struct {

	// The Amazon Resource Name (ARN) of the on-call rotation to retrieve information
	// about.
	//
	// This member is required.
	RotationId *string

	noSmithyDocumentSerde
}

type GetRotationOutput struct {

	// The Amazon Resource Names (ARNs) of the contacts assigned to the on-call
	// rotation team.
	//
	// This member is required.
	ContactIds []string

	// The name of the on-call rotation.
	//
	// This member is required.
	Name *string

	// Specifies how long a rotation lasts before restarting at the beginning of the
	// shift order.
	//
	// This member is required.
	Recurrence *types.RecurrenceSettings

	// The Amazon Resource Name (ARN) of the on-call rotation.
	//
	// This member is required.
	RotationArn *string

	// The specified start time for the on-call rotation.
	//
	// This member is required.
	StartTime *time.Time

	// The time zone that the rotation’s activity is based on, in Internet Assigned
	// Numbers Authority (IANA) format.
	//
	// This member is required.
	TimeZoneId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRotationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRotation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRotation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRotation"); err != nil {
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
	if err = addOpGetRotationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRotation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRotation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRotation",
	}
}
