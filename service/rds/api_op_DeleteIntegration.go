// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes a zero-ETL integration with Amazon Redshift.
func (c *Client) DeleteIntegration(ctx context.Context, params *DeleteIntegrationInput, optFns ...func(*Options)) (*DeleteIntegrationOutput, error) {
	if params == nil {
		params = &DeleteIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIntegration", params, optFns, c.addOperationDeleteIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIntegrationInput struct {

	// The unique identifier of the integration.
	//
	// This member is required.
	IntegrationIdentifier *string

	noSmithyDocumentSerde
}

// A zero-ETL integration with Amazon Redshift.
type DeleteIntegrationOutput struct {

	// The encryption context for the integration. For more information, see [Encryption context] in the
	// Amazon Web Services Key Management Service Developer Guide.
	//
	// [Encryption context]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context
	AdditionalEncryptionContext map[string]string

	// The time when the integration was created, in Universal Coordinated Time (UTC).
	CreateTime *time.Time

	// Data filters for the integration. These filters determine which tables from the
	// source database are sent to the target Amazon Redshift data warehouse.
	DataFilter *string

	// A description of the integration.
	Description *string

	// Any errors associated with the integration.
	Errors []types.IntegrationError

	// The ARN of the integration.
	IntegrationArn *string

	// The name of the integration.
	IntegrationName *string

	// The Amazon Web Services Key Management System (Amazon Web Services KMS) key
	// identifier for the key used to to encrypt the integration.
	KMSKeyId *string

	// The Amazon Resource Name (ARN) of the database used as the source for
	// replication.
	SourceArn *string

	// The current status of the integration.
	Status types.IntegrationStatus

	// A list of tags. For more information, see [Tagging Amazon RDS Resources] in the Amazon RDS User Guide.
	//
	// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	Tags []types.Tag

	// The ARN of the Redshift data warehouse used as the target for replication.
	TargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteIntegration"); err != nil {
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
	if err = addOpDeleteIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIntegration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteIntegration",
	}
}
