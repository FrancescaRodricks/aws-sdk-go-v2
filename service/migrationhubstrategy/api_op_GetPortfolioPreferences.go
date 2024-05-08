// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves your migration and modernization preferences.
func (c *Client) GetPortfolioPreferences(ctx context.Context, params *GetPortfolioPreferencesInput, optFns ...func(*Options)) (*GetPortfolioPreferencesOutput, error) {
	if params == nil {
		params = &GetPortfolioPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPortfolioPreferences", params, optFns, c.addOperationGetPortfolioPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPortfolioPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPortfolioPreferencesInput struct {
	noSmithyDocumentSerde
}

type GetPortfolioPreferencesOutput struct {

	// The classification for application component types.
	ApplicationMode types.ApplicationMode

	//  The transformation preferences for non-database applications.
	ApplicationPreferences *types.ApplicationPreferences

	//  The transformation preferences for database applications.
	DatabasePreferences *types.DatabasePreferences

	//  The rank of business goals based on priority.
	PrioritizeBusinessGoals *types.PrioritizeBusinessGoals

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPortfolioPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPortfolioPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPortfolioPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPortfolioPreferences"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPortfolioPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPortfolioPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPortfolioPreferences",
	}
}
