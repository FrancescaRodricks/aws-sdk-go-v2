// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an existing opt-out list. All opted out phone numbers in the opt-out
// list are deleted.
//
// If the specified opt-out list name doesn't exist or is in-use by an origination
// phone number or pool, an error is returned.
func (c *Client) DeleteOptOutList(ctx context.Context, params *DeleteOptOutListInput, optFns ...func(*Options)) (*DeleteOptOutListOutput, error) {
	if params == nil {
		params = &DeleteOptOutListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOptOutList", params, optFns, c.addOperationDeleteOptOutListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOptOutListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOptOutListInput struct {

	// The OptOutListName or OptOutListArn of the OptOutList to delete. You can use DescribeOptOutLists
	// to find the values for OptOutListName and OptOutListArn.
	//
	// This member is required.
	OptOutListName *string

	noSmithyDocumentSerde
}

type DeleteOptOutListOutput struct {

	// The time when the OptOutList was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the OptOutList that was removed.
	OptOutListArn *string

	// The name of the OptOutList that was removed.
	OptOutListName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOptOutListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteOptOutList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteOptOutList{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteOptOutList"); err != nil {
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
	if err = addOpDeleteOptOutListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOptOutList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOptOutList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteOptOutList",
	}
}
