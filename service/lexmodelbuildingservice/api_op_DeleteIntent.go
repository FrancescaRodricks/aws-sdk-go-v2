// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes all versions of the intent, including the $LATEST version. To delete a
// specific version of the intent, use the DeleteIntentVersionoperation.
//
// You can delete a version of an intent only if it is not referenced. To delete
// an intent that is referred to in one or more bots (see how-it-works), you must remove those
// references first.
//
// If you get the ResourceInUseException exception, it provides an example
// reference that shows where the intent is referenced. To remove the reference to
// the intent, either update the bot or delete it. If you get the same exception
// when you attempt to delete the intent again, repeat until the intent has no
// references and the call to DeleteIntent is successful.
//
// This operation requires permission for the lex:DeleteIntent action.
func (c *Client) DeleteIntent(ctx context.Context, params *DeleteIntentInput, optFns ...func(*Options)) (*DeleteIntentOutput, error) {
	if params == nil {
		params = &DeleteIntentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIntent", params, optFns, c.addOperationDeleteIntentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIntentInput struct {

	// The name of the intent. The name is case sensitive.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteIntentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIntentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIntent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIntent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteIntent"); err != nil {
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
	if err = addOpDeleteIntentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIntent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIntent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteIntent",
	}
}
