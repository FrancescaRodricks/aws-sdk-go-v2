// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a business glossary in Amazon DataZone.
func (c *Client) GetGlossary(ctx context.Context, params *GetGlossaryInput, optFns ...func(*Options)) (*GetGlossaryOutput, error) {
	if params == nil {
		params = &GetGlossaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGlossary", params, optFns, c.addOperationGetGlossaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGlossaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGlossaryInput struct {

	// The ID of the Amazon DataZone domain in which this business glossary exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the business glossary.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetGlossaryOutput struct {

	// The ID of the Amazon DataZone domain in which this business glossary exists.
	//
	// This member is required.
	DomainId *string

	// The ID of the business glossary.
	//
	// This member is required.
	Id *string

	// The name of the business glossary.
	//
	// This member is required.
	Name *string

	// The ID of the project that owns this business glossary.
	//
	// This member is required.
	OwningProjectId *string

	// The status of the business glossary.
	//
	// This member is required.
	Status types.GlossaryStatus

	// The timestamp of when this business glossary was created.
	CreatedAt *time.Time

	// The Amazon DataZone user who created this business glossary.
	CreatedBy *string

	// The description of the business glossary.
	Description *string

	// The timestamp of when the business glossary was updated.
	UpdatedAt *time.Time

	// The Amazon DataZone user who updated the business glossary.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGlossaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGlossary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGlossary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGlossary"); err != nil {
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
	if err = addOpGetGlossaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGlossary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGlossary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGlossary",
	}
}
