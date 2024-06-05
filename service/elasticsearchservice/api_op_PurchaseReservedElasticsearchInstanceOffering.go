// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to purchase reserved Elasticsearch instances.
func (c *Client) PurchaseReservedElasticsearchInstanceOffering(ctx context.Context, params *PurchaseReservedElasticsearchInstanceOfferingInput, optFns ...func(*Options)) (*PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
	if params == nil {
		params = &PurchaseReservedElasticsearchInstanceOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseReservedElasticsearchInstanceOffering", params, optFns, c.addOperationPurchaseReservedElasticsearchInstanceOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseReservedElasticsearchInstanceOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for parameters to PurchaseReservedElasticsearchInstanceOffering
type PurchaseReservedElasticsearchInstanceOfferingInput struct {

	// A customer-specified identifier to track this reservation.
	//
	// This member is required.
	ReservationName *string

	// The ID of the reserved Elasticsearch instance offering to purchase.
	//
	// This member is required.
	ReservedElasticsearchInstanceOfferingId *string

	// The number of Elasticsearch instances to reserve.
	InstanceCount *int32

	noSmithyDocumentSerde
}

// Represents the output of a PurchaseReservedElasticsearchInstanceOffering
// operation.
type PurchaseReservedElasticsearchInstanceOfferingOutput struct {

	// The customer-specified identifier used to track this reservation.
	ReservationName *string

	// Details of the reserved Elasticsearch instance which was purchased.
	ReservedElasticsearchInstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPurchaseReservedElasticsearchInstanceOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPurchaseReservedElasticsearchInstanceOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPurchaseReservedElasticsearchInstanceOffering{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PurchaseReservedElasticsearchInstanceOffering"); err != nil {
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
	if err = addOpPurchaseReservedElasticsearchInstanceOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedElasticsearchInstanceOffering(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseReservedElasticsearchInstanceOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PurchaseReservedElasticsearchInstanceOffering",
	}
}
