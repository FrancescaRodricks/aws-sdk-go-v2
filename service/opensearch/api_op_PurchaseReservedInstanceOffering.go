// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to purchase Amazon OpenSearch Service Reserved Instances.
func (c *Client) PurchaseReservedInstanceOffering(ctx context.Context, params *PurchaseReservedInstanceOfferingInput, optFns ...func(*Options)) (*PurchaseReservedInstanceOfferingOutput, error) {
	if params == nil {
		params = &PurchaseReservedInstanceOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseReservedInstanceOffering", params, optFns, c.addOperationPurchaseReservedInstanceOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseReservedInstanceOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to the PurchaseReservedInstanceOffering
// operation.
type PurchaseReservedInstanceOfferingInput struct {

	// A customer-specified identifier to track this reservation.
	//
	// This member is required.
	ReservationName *string

	// The ID of the Reserved Instance offering to purchase.
	//
	// This member is required.
	ReservedInstanceOfferingId *string

	// The number of OpenSearch instances to reserve.
	InstanceCount *int32

	noSmithyDocumentSerde
}

// Represents the output of a PurchaseReservedInstanceOffering operation.
type PurchaseReservedInstanceOfferingOutput struct {

	// The customer-specified identifier used to track this reservation.
	ReservationName *string

	// The ID of the Reserved Instance offering that was purchased.
	ReservedInstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPurchaseReservedInstanceOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPurchaseReservedInstanceOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPurchaseReservedInstanceOffering{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PurchaseReservedInstanceOffering"); err != nil {
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
	if err = addOpPurchaseReservedInstanceOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedInstanceOffering(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseReservedInstanceOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PurchaseReservedInstanceOffering",
	}
}
