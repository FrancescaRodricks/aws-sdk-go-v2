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

// Gets metadata information about a slot.
func (c *Client) DescribeSlot(ctx context.Context, params *DescribeSlotInput, optFns ...func(*Options)) (*DescribeSlotOutput, error) {
	if params == nil {
		params = &DescribeSlotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSlot", params, optFns, c.addOperationDescribeSlotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSlotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSlotInput struct {

	// The identifier of the bot associated with the slot.
	//
	// This member is required.
	BotId *string

	// The version of the bot associated with the slot.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the intent that contains the slot.
	//
	// This member is required.
	IntentId *string

	// The identifier of the language and locale of the slot to describe. The string
	// must match one of the supported locales. For more information, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	// The unique identifier for the slot.
	//
	// This member is required.
	SlotId *string

	noSmithyDocumentSerde
}

type DescribeSlotOutput struct {

	// The identifier of the bot associated with the slot.
	BotId *string

	// The version of the bot associated with the slot.
	BotVersion *string

	// A timestamp of the date and time that the slot was created.
	CreationDateTime *time.Time

	// The description specified for the slot.
	Description *string

	// The identifier of the intent associated with the slot.
	IntentId *string

	// A timestamp of the date and time that the slot was last updated.
	LastUpdatedDateTime *time.Time

	// The language and locale specified for the slot.
	LocaleId *string

	// Indicates whether the slot accepts multiple values in a single utterance.
	//
	// If the multipleValuesSetting is not set, the default value is false .
	MultipleValuesSetting *types.MultipleValuesSetting

	// Whether slot values are shown in Amazon CloudWatch logs. If the value is None ,
	// the actual value of the slot is shown in logs.
	ObfuscationSetting *types.ObfuscationSetting

	// The unique identifier generated for the slot.
	SlotId *string

	// The name specified for the slot.
	SlotName *string

	// The identifier of the slot type that determines the values entered into the
	// slot.
	SlotTypeId *string

	// Specifications for the constituent sub slots and the expression for the
	// composite slot.
	SubSlotSetting *types.SubSlotSetting

	// Prompts that Amazon Lex uses to elicit a value for the slot.
	ValueElicitationSetting *types.SlotValueElicitationSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSlotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSlot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSlot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSlot"); err != nil {
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
	if err = addOpDescribeSlotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSlot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSlot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSlot",
	}
}
