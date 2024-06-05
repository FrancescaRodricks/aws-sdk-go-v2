// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the migration project using the specified parameters.
//
// You can run this action only after you create an instance profile and data
// providers using [CreateInstanceProfile]and [CreateDataProvider].
//
// [CreateDataProvider]: https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateDataProvider.html
// [CreateInstanceProfile]: https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateInstanceProfile.html
func (c *Client) CreateMigrationProject(ctx context.Context, params *CreateMigrationProjectInput, optFns ...func(*Options)) (*CreateMigrationProjectOutput, error) {
	if params == nil {
		params = &CreateMigrationProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMigrationProject", params, optFns, c.addOperationCreateMigrationProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMigrationProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMigrationProjectInput struct {

	// The identifier of the associated instance profile. Identifiers must begin with
	// a letter and must contain only ASCII letters, digits, and hyphens. They can't
	// end with a hyphen, or contain two consecutive hyphens.
	//
	// This member is required.
	InstanceProfileIdentifier *string

	// Information about the source data provider, including the name, ARN, and
	// Secrets Manager parameters.
	//
	// This member is required.
	SourceDataProviderDescriptors []types.DataProviderDescriptorDefinition

	// Information about the target data provider, including the name, ARN, and Amazon
	// Web Services Secrets Manager parameters.
	//
	// This member is required.
	TargetDataProviderDescriptors []types.DataProviderDescriptorDefinition

	// A user-friendly description of the migration project.
	Description *string

	// A user-friendly name for the migration project.
	MigrationProjectName *string

	// The schema conversion application attributes, including the Amazon S3 bucket
	// name and Amazon S3 role ARN.
	SchemaConversionApplicationAttributes *types.SCApplicationAttributes

	// One or more tags to be assigned to the migration project.
	Tags []types.Tag

	// The settings in JSON format for migration rules. Migration rules make it
	// possible for you to change the object names according to the rules that you
	// specify. For example, you can change an object name to lowercase or uppercase,
	// add or remove a prefix or suffix, or rename objects.
	TransformationRules *string

	noSmithyDocumentSerde
}

type CreateMigrationProjectOutput struct {

	// The migration project that was created.
	MigrationProject *types.MigrationProject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMigrationProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMigrationProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMigrationProject{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMigrationProject"); err != nil {
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
	if err = addOpCreateMigrationProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMigrationProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMigrationProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMigrationProject",
	}
}
