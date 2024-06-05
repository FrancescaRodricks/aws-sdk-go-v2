// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the finding aggregation configuration. Used to update the Region
// linking mode and the list of included or excluded Regions. You cannot use
// UpdateFindingAggregator to change the aggregation Region.
//
// You must run UpdateFindingAggregator from the current aggregation Region.
func (c *Client) UpdateFindingAggregator(ctx context.Context, params *UpdateFindingAggregatorInput, optFns ...func(*Options)) (*UpdateFindingAggregatorOutput, error) {
	if params == nil {
		params = &UpdateFindingAggregatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFindingAggregator", params, optFns, c.addOperationUpdateFindingAggregatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFindingAggregatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFindingAggregatorInput struct {

	// The ARN of the finding aggregator. To obtain the ARN, use ListFindingAggregators
	// .
	//
	// This member is required.
	FindingAggregatorArn *string

	// Indicates whether to aggregate findings from all of the available Regions in
	// the current partition. Also determines whether to automatically aggregate
	// findings from new Regions as Security Hub supports them and you opt into them.
	//
	// The selected option also determines how to use the Regions provided in the
	// Regions list.
	//
	// The options are as follows:
	//
	//   - ALL_REGIONS - Indicates to aggregate findings from all of the Regions where
	//   Security Hub is enabled. When you choose this option, Security Hub also
	//   automatically aggregates findings from new Regions as Security Hub supports them
	//   and you opt into them.
	//
	//   - ALL_REGIONS_EXCEPT_SPECIFIED - Indicates to aggregate findings from all of
	//   the Regions where Security Hub is enabled, except for the Regions listed in the
	//   Regions parameter. When you choose this option, Security Hub also
	//   automatically aggregates findings from new Regions as Security Hub supports them
	//   and you opt into them.
	//
	//   - SPECIFIED_REGIONS - Indicates to aggregate findings only from the Regions
	//   listed in the Regions parameter. Security Hub does not automatically aggregate
	//   findings from new Regions.
	//
	// This member is required.
	RegionLinkingMode *string

	// If RegionLinkingMode is ALL_REGIONS_EXCEPT_SPECIFIED , then this is a
	// space-separated list of Regions that do not aggregate findings to the
	// aggregation Region.
	//
	// If RegionLinkingMode is SPECIFIED_REGIONS , then this is a space-separated list
	// of Regions that do aggregate findings to the aggregation Region.
	Regions []string

	noSmithyDocumentSerde
}

type UpdateFindingAggregatorOutput struct {

	// The aggregation Region.
	FindingAggregationRegion *string

	// The ARN of the finding aggregator.
	FindingAggregatorArn *string

	// Indicates whether to link all Regions, all Regions except for a list of
	// excluded Regions, or a list of included Regions.
	RegionLinkingMode *string

	// The list of excluded Regions or included Regions.
	Regions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFindingAggregatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFindingAggregator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFindingAggregator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFindingAggregator"); err != nil {
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
	if err = addOpUpdateFindingAggregatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFindingAggregator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFindingAggregator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFindingAggregator",
	}
}
