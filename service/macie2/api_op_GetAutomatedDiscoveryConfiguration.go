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

// Retrieves the configuration settings and status of automated sensitive data
// discovery for an account.
func (c *Client) GetAutomatedDiscoveryConfiguration(ctx context.Context, params *GetAutomatedDiscoveryConfigurationInput, optFns ...func(*Options)) (*GetAutomatedDiscoveryConfigurationOutput, error) {
	if params == nil {
		params = &GetAutomatedDiscoveryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAutomatedDiscoveryConfiguration", params, optFns, c.addOperationGetAutomatedDiscoveryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAutomatedDiscoveryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAutomatedDiscoveryConfigurationInput struct {
	noSmithyDocumentSerde
}

type GetAutomatedDiscoveryConfigurationOutput struct {

	// The unique identifier for the classification scope that's used when performing
	// automated sensitive data discovery for the account. The classification scope
	// specifies S3 buckets to exclude from automated sensitive data discovery.
	ClassificationScopeId *string

	// The date and time, in UTC and extended ISO 8601 format, when automated
	// sensitive data discovery was most recently disabled for the account. This value
	// is null if automated sensitive data discovery wasn't enabled and subsequently
	// disabled for the account.
	DisabledAt *time.Time

	// The date and time, in UTC and extended ISO 8601 format, when automated
	// sensitive data discovery was initially enabled for the account. This value is
	// null if automated sensitive data discovery has never been enabled for the
	// account.
	FirstEnabledAt *time.Time

	// The date and time, in UTC and extended ISO 8601 format, when automated
	// sensitive data discovery was most recently enabled or disabled for the account.
	LastUpdatedAt *time.Time

	// The unique identifier for the sensitivity inspection template that's used when
	// performing automated sensitive data discovery for the account. The template
	// specifies which allow lists, custom data identifiers, and managed data
	// identifiers to use when analyzing data.
	SensitivityInspectionTemplateId *string

	// The current status of the automated sensitive data discovery configuration for
	// the account. Possible values are: ENABLED, use the specified settings to perform
	// automated sensitive data discovery activities for the account; and, DISABLED,
	// don't perform automated sensitive data discovery activities for the account.
	Status types.AutomatedDiscoveryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAutomatedDiscoveryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAutomatedDiscoveryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAutomatedDiscoveryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAutomatedDiscoveryConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAutomatedDiscoveryConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAutomatedDiscoveryConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAutomatedDiscoveryConfiguration",
	}
}
