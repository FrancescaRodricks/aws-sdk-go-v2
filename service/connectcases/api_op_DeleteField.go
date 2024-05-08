// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a field from a cases template. You can delete up to 100 fields per
// domain.
//
// After a field is deleted:
//
//   - You can still retrieve the field by calling BatchGetField .
//
//   - You cannot update a deleted field by calling UpdateField ; it throws a
//     ValidationException .
//
//   - Deleted fields are not included in the ListFields response.
//
//   - Calling CreateCase with a deleted field throws a ValidationException
//     denoting which field IDs in the request have been deleted.
//
//   - Calling GetCase with a deleted field ID returns the deleted field's value if
//     one exists.
//
//   - Calling UpdateCase with a deleted field ID throws a ValidationException if
//     the case does not already contain a value for the deleted field. Otherwise it
//     succeeds, allowing you to update or remove (using emptyValue: {} ) the field's
//     value from the case.
//
//   - GetTemplate does not return field IDs for deleted fields.
//
//   - GetLayout does not return field IDs for deleted fields.
//
//   - Calling SearchCases with the deleted field ID as a filter returns any cases
//     that have a value for the deleted field that matches the filter criteria.
//
//   - Calling SearchCases with a searchTerm value that matches a deleted field's
//     value on a case returns the case in the response.
//
//   - Calling BatchPutFieldOptions with a deleted field ID throw a
//     ValidationException .
//
//   - Calling GetCaseEventConfiguration does not return field IDs for deleted
//     fields.
func (c *Client) DeleteField(ctx context.Context, params *DeleteFieldInput, optFns ...func(*Options)) (*DeleteFieldOutput, error) {
	if params == nil {
		params = &DeleteFieldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteField", params, optFns, c.addOperationDeleteFieldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFieldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFieldInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// Unique identifier of the field.
	//
	// This member is required.
	FieldId *string

	noSmithyDocumentSerde
}

type DeleteFieldOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFieldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteField{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteField{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteField"); err != nil {
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
	if err = addOpDeleteFieldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteField(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteField(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteField",
	}
}
