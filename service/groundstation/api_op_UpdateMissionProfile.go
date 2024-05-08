// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a mission profile.
//
// Updating a mission profile will not update the execution parameters for
// existing future contacts.
func (c *Client) UpdateMissionProfile(ctx context.Context, params *UpdateMissionProfileInput, optFns ...func(*Options)) (*UpdateMissionProfileOutput, error) {
	if params == nil {
		params = &UpdateMissionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMissionProfile", params, optFns, c.addOperationUpdateMissionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMissionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMissionProfileInput struct {

	// UUID of a mission profile.
	//
	// This member is required.
	MissionProfileId *string

	// Amount of time after a contact ends that you’d like to receive a Ground Station
	// Contact State Change event indicating the pass has finished.
	ContactPostPassDurationSeconds *int32

	// Amount of time after a contact ends that you’d like to receive a Ground Station
	// Contact State Change event indicating the pass has finished.
	ContactPrePassDurationSeconds *int32

	// A list of lists of ARNs. Each list of ARNs is an edge, with a from Config and a
	// to Config .
	DataflowEdges [][]string

	// Smallest amount of time in seconds that you’d like to see for an available
	// contact. AWS Ground Station will not present you with contacts shorter than this
	// duration.
	MinimumViableContactDurationSeconds *int32

	// Name of a mission profile.
	Name *string

	// KMS key to use for encrypting streams.
	StreamsKmsKey types.KmsKey

	// Role to use for encrypting streams with KMS key.
	StreamsKmsRole *string

	// ARN of a tracking Config .
	TrackingConfigArn *string

	noSmithyDocumentSerde
}

type UpdateMissionProfileOutput struct {

	// UUID of a mission profile.
	MissionProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMissionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMissionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMissionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMissionProfile"); err != nil {
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
	if err = addOpUpdateMissionProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMissionProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMissionProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMissionProfile",
	}
}
