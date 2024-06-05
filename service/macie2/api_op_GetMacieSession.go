// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the status and configuration settings for an Amazon Macie account.
func (c *Client) GetMacieSession(ctx context.Context, params *GetMacieSessionInput, optFns ...func(*Options)) (*GetMacieSessionOutput, error) {
	if params == nil {
		params = &GetMacieSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMacieSession", params, optFns, c.addOperationGetMacieSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMacieSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMacieSessionInput struct {
	noSmithyDocumentSerde
}

type GetMacieSessionOutput struct {

	// The date and time, in UTC and extended ISO 8601 format, when the Amazon Macie
	// account was created.
	CreatedAt *time.Time

	// The frequency with which Amazon Macie publishes updates to policy findings for
	// the account. This includes publishing updates to Security Hub and Amazon
	// EventBridge (formerly Amazon CloudWatch Events).
	FindingPublishingFrequency types.FindingPublishingFrequency

	// The Amazon Resource Name (ARN) of the service-linked role that allows Amazon
	// Macie to monitor and analyze data in Amazon Web Services resources for the
	// account.
	ServiceRole *string

	// The current status of the Amazon Macie account. Possible values are: PAUSED,
	// the account is enabled but all Macie activities are suspended (paused) for the
	// account; and, ENABLED, the account is enabled and all Macie activities are
	// enabled for the account.
	Status types.MacieStatus

	// The date and time, in UTC and extended ISO 8601 format, of the most recent
	// change to the status or configuration settings for the Amazon Macie account.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMacieSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMacieSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMacieSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMacieSession"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMacieSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMacieSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMacieSession",
	}
}
