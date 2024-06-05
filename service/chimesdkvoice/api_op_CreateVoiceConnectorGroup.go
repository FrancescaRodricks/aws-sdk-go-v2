// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Chime SDK Voice Connector group under the administrator's AWS
// account. You can associate Amazon Chime SDK Voice Connectors with the Voice
// Connector group by including VoiceConnectorItems in the request.
//
// You can include Voice Connectors from different AWS Regions in your group. This
// creates a fault tolerant mechanism for fallback in case of availability events.
func (c *Client) CreateVoiceConnectorGroup(ctx context.Context, params *CreateVoiceConnectorGroupInput, optFns ...func(*Options)) (*CreateVoiceConnectorGroupOutput, error) {
	if params == nil {
		params = &CreateVoiceConnectorGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVoiceConnectorGroup", params, optFns, c.addOperationCreateVoiceConnectorGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVoiceConnectorGroupInput struct {

	// The name of the Voice Connector group.
	//
	// This member is required.
	Name *string

	// Lists the Voice Connectors that inbound calls are routed to.
	VoiceConnectorItems []types.VoiceConnectorItem

	noSmithyDocumentSerde
}

type CreateVoiceConnectorGroupOutput struct {

	// The details of the Voice Connector group.
	VoiceConnectorGroup *types.VoiceConnectorGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVoiceConnectorGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVoiceConnectorGroup"); err != nil {
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
	if err = addOpCreateVoiceConnectorGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVoiceConnectorGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVoiceConnectorGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVoiceConnectorGroup",
	}
}
