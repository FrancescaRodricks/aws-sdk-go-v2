// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates details of the specified Amazon Chime Voice Connector group, such as
// the name and Amazon Chime Voice Connector priority ranking.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [UpdateVoiceConnectorGroup], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by UpdateVoiceConnectorGroup in the Amazon Chime SDK Voice
// Namespace
//
// [UpdateVoiceConnectorGroup]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_UpdateVoiceConnectorGroup.html
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
func (c *Client) UpdateVoiceConnectorGroup(ctx context.Context, params *UpdateVoiceConnectorGroupInput, optFns ...func(*Options)) (*UpdateVoiceConnectorGroupOutput, error) {
	if params == nil {
		params = &UpdateVoiceConnectorGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVoiceConnectorGroup", params, optFns, c.addOperationUpdateVoiceConnectorGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVoiceConnectorGroupInput struct {

	// The name of the Amazon Chime Voice Connector group.
	//
	// This member is required.
	Name *string

	// The Amazon Chime Voice Connector group ID.
	//
	// This member is required.
	VoiceConnectorGroupId *string

	// The VoiceConnectorItems to associate with the group.
	//
	// This member is required.
	VoiceConnectorItems []types.VoiceConnectorItem

	noSmithyDocumentSerde
}

type UpdateVoiceConnectorGroupOutput struct {

	// The updated Amazon Chime Voice Connector group details.
	VoiceConnectorGroup *types.VoiceConnectorGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVoiceConnectorGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateVoiceConnectorGroup"); err != nil {
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
	if err = addOpUpdateVoiceConnectorGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVoiceConnectorGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVoiceConnectorGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateVoiceConnectorGroup",
	}
}
