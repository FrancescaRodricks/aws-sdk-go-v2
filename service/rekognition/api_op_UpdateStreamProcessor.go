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

//	Allows you to update a stream processor. You can change some settings and
//
// regions of interest and delete certain parameters.
func (c *Client) UpdateStreamProcessor(ctx context.Context, params *UpdateStreamProcessorInput, optFns ...func(*Options)) (*UpdateStreamProcessorOutput, error) {
	if params == nil {
		params = &UpdateStreamProcessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStreamProcessor", params, optFns, c.addOperationUpdateStreamProcessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStreamProcessorInput struct {

	//  Name of the stream processor that you want to update.
	//
	// This member is required.
	Name *string

	//  Shows whether you are sharing data with Rekognition to improve model
	// performance. You can choose this option at the account level or on a per-stream
	// basis. Note that if you opt out at the account level this setting is ignored on
	// individual streams.
	DataSharingPreferenceForUpdate *types.StreamProcessorDataSharingPreference

	//  A list of parameters you want to delete from the stream processor.
	ParametersToDelete []types.StreamProcessorParameterToDelete

	//  Specifies locations in the frames where Amazon Rekognition checks for objects
	// or people. This is an optional parameter for label detection stream processors.
	RegionsOfInterestForUpdate []types.RegionOfInterest

	//  The stream processor settings that you want to update. Label detection
	// settings can be updated to detect different labels with a different minimum
	// confidence.
	SettingsForUpdate *types.StreamProcessorSettingsForUpdate

	noSmithyDocumentSerde
}

type UpdateStreamProcessorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStreamProcessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateStreamProcessor"); err != nil {
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
	if err = addOpUpdateStreamProcessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStreamProcessor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStreamProcessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateStreamProcessor",
	}
}
