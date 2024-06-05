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

// Associates phone numbers with the specified Amazon Chime SDK Voice Connector
// group.
func (c *Client) AssociatePhoneNumbersWithVoiceConnectorGroup(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorGroupInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) {
	if params == nil {
		params = &AssociatePhoneNumbersWithVoiceConnectorGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociatePhoneNumbersWithVoiceConnectorGroup", params, optFns, c.addOperationAssociatePhoneNumbersWithVoiceConnectorGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociatePhoneNumbersWithVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePhoneNumbersWithVoiceConnectorGroupInput struct {

	// List of phone numbers, in E.164 format.
	//
	// This member is required.
	E164PhoneNumbers []string

	// The Amazon Chime SDK Voice Connector group ID.
	//
	// This member is required.
	VoiceConnectorGroupId *string

	// If true, associates the provided phone numbers with the provided Amazon Chime
	// SDK Voice Connector Group and removes any previously existing associations. If
	// false, does not associate any phone numbers that have previously existing
	// associations.
	ForceAssociate *bool

	noSmithyDocumentSerde
}

type AssociatePhoneNumbersWithVoiceConnectorGroupOutput struct {

	// If the action fails for one or more of the phone numbers in the request, a list
	// of the phone numbers is returned, along with error codes and error messages.
	PhoneNumberErrors []types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociatePhoneNumbersWithVoiceConnectorGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociatePhoneNumbersWithVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociatePhoneNumbersWithVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociatePhoneNumbersWithVoiceConnectorGroup"); err != nil {
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
	if err = addOpAssociatePhoneNumbersWithVoiceConnectorGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePhoneNumbersWithVoiceConnectorGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociatePhoneNumbersWithVoiceConnectorGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociatePhoneNumbersWithVoiceConnectorGroup",
	}
}
