// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about when the specified access key was last used. The
// information includes the date and time of last use, along with the Amazon Web
// Services service and Region that were specified in the last request made with
// that key.
func (c *Client) GetAccessKeyLastUsed(ctx context.Context, params *GetAccessKeyLastUsedInput, optFns ...func(*Options)) (*GetAccessKeyLastUsedOutput, error) {
	if params == nil {
		params = &GetAccessKeyLastUsedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessKeyLastUsed", params, optFns, c.addOperationGetAccessKeyLastUsedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessKeyLastUsedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessKeyLastUsedInput struct {

	// The identifier of an access key.
	//
	// This parameter allows (through its [regex pattern]) a string of characters that can consist of
	// any upper or lowercased letter or digit.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	AccessKeyId *string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetAccessKeyLastUsed request. It is also returned as a member
// of the AccessKeyMetaDatastructure returned by the ListAccessKeys action.
type GetAccessKeyLastUsedOutput struct {

	// Contains information about the last time the access key was used.
	AccessKeyLastUsed *types.AccessKeyLastUsed

	// The name of the IAM user that owns this access key.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessKeyLastUsedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetAccessKeyLastUsed{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetAccessKeyLastUsed{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessKeyLastUsed"); err != nil {
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
	if err = addOpGetAccessKeyLastUsedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessKeyLastUsed(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAccessKeyLastUsed(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessKeyLastUsed",
	}
}
