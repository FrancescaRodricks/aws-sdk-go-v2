// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a custom slot type
//
// To create a custom slot type, specify a name for the slot type and a set of
// enumeration values, the values that a slot of this type can assume.
func (c *Client) CreateSlotType(ctx context.Context, params *CreateSlotTypeInput, optFns ...func(*Options)) (*CreateSlotTypeOutput, error) {
	if params == nil {
		params = &CreateSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSlotType", params, optFns, c.addOperationCreateSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSlotTypeInput struct {

	// The identifier of the bot associated with this slot type.
	//
	// This member is required.
	BotId *string

	// The identifier of the bot version associated with this slot type.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale that the slot type will be used in.
	// The string must match one of the supported locales. All of the bots, intents,
	// and slots used by the slot type must have the same locale. For more information,
	// see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	// The name for the slot. A slot type name must be unique within the intent.
	//
	// This member is required.
	SlotTypeName *string

	// Specifications for a composite slot type.
	CompositeSlotTypeSetting *types.CompositeSlotTypeSetting

	// A description of the slot type. Use the description to help identify the slot
	// type in lists.
	Description *string

	// Sets the type of external information used to create the slot type.
	ExternalSourceSetting *types.ExternalSourceSetting

	// The built-in slot type used as a parent of this slot type. When you define a
	// parent slot type, the new slot type has the configuration of the parent slot
	// type.
	//
	// Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature *string

	// A list of SlotTypeValue objects that defines the values that the slot type can
	// take. Each value can have a list of synonyms, additional values that help train
	// the machine learning model about the values that it resolves for a slot.
	SlotTypeValues []types.SlotTypeValue

	// Determines the strategy that Amazon Lex uses to select a value from the list of
	// possible values. The field can be set to one of the following values:
	//
	//   - ORIGINAL_VALUE - Returns the value entered by the user, if the user value is
	//   similar to the slot value.
	//
	//   - TOP_RESOLUTION - If there is a resolution list for the slot, return the
	//   first value in the resolution list. If there is no resolution list, return null.
	//
	// If you don't specify the valueSelectionSetting parameter, the default is
	// ORIGINAL_VALUE .
	ValueSelectionSetting *types.SlotValueSelectionSetting

	noSmithyDocumentSerde
}

type CreateSlotTypeOutput struct {

	// The identifier for the bot associated with the slot type.
	BotId *string

	// The version of the bot associated with the slot type.
	BotVersion *string

	// Specifications for a composite slot type.
	CompositeSlotTypeSetting *types.CompositeSlotTypeSetting

	// A timestamp of the date and time that the slot type was created.
	CreationDateTime *time.Time

	// The description specified for the slot type.
	Description *string

	// The type of external information used to create the slot type.
	ExternalSourceSetting *types.ExternalSourceSetting

	// The specified language and local specified for the slot type.
	LocaleId *string

	// The signature of the base slot type specified for the slot type.
	ParentSlotTypeSignature *string

	// The unique identifier assigned to the slot type. Use this to identify the slot
	// type in the UpdateSlotType and DeleteSlotType operations.
	SlotTypeId *string

	// The name specified for the slot type.
	SlotTypeName *string

	// The list of values that the slot type can assume.
	SlotTypeValues []types.SlotTypeValue

	// The strategy that Amazon Lex uses to select a value from the list of possible
	// values.
	ValueSelectionSetting *types.SlotValueSelectionSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSlotType"); err != nil {
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
	if err = addOpCreateSlotTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSlotType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSlotType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSlotType",
	}
}
