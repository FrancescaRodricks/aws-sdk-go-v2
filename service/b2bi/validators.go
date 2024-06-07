// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateCapability struct {
}

func (*validateOpCreateCapability) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCapability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCapabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCapabilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePartnership struct {
}

func (*validateOpCreatePartnership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePartnership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePartnershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePartnershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProfile struct {
}

func (*validateOpCreateProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTransformer struct {
}

func (*validateOpCreateTransformer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTransformer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTransformerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTransformerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCapability struct {
}

func (*validateOpDeleteCapability) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCapability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCapabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCapabilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePartnership struct {
}

func (*validateOpDeletePartnership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePartnership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePartnershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePartnershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProfile struct {
}

func (*validateOpDeleteProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTransformer struct {
}

func (*validateOpDeleteTransformer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTransformer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTransformerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTransformerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCapability struct {
}

func (*validateOpGetCapability) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCapability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCapabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCapabilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPartnership struct {
}

func (*validateOpGetPartnership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPartnership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPartnershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPartnershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProfile struct {
}

func (*validateOpGetProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTransformer struct {
}

func (*validateOpGetTransformer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTransformer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTransformerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTransformerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTransformerJob struct {
}

func (*validateOpGetTransformerJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTransformerJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTransformerJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTransformerJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTransformerJob struct {
}

func (*validateOpStartTransformerJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTransformerJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTransformerJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTransformerJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTestMapping struct {
}

func (*validateOpTestMapping) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTestMapping) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TestMappingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTestMappingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTestParsing struct {
}

func (*validateOpTestParsing) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTestParsing) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TestParsingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTestParsingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCapability struct {
}

func (*validateOpUpdateCapability) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCapability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCapabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCapabilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdatePartnership struct {
}

func (*validateOpUpdatePartnership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePartnership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePartnershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePartnershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateProfile struct {
}

func (*validateOpUpdateProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateTransformer struct {
}

func (*validateOpUpdateTransformer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateTransformer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateTransformerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateTransformerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateCapabilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCapability{}, middleware.After)
}

func addOpCreatePartnershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePartnership{}, middleware.After)
}

func addOpCreateProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProfile{}, middleware.After)
}

func addOpCreateTransformerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTransformer{}, middleware.After)
}

func addOpDeleteCapabilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCapability{}, middleware.After)
}

func addOpDeletePartnershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePartnership{}, middleware.After)
}

func addOpDeleteProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProfile{}, middleware.After)
}

func addOpDeleteTransformerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTransformer{}, middleware.After)
}

func addOpGetCapabilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCapability{}, middleware.After)
}

func addOpGetPartnershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPartnership{}, middleware.After)
}

func addOpGetProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProfile{}, middleware.After)
}

func addOpGetTransformerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTransformer{}, middleware.After)
}

func addOpGetTransformerJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTransformerJob{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartTransformerJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTransformerJob{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpTestMappingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTestMapping{}, middleware.After)
}

func addOpTestParsingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTestParsing{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateCapabilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCapability{}, middleware.After)
}

func addOpUpdatePartnershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePartnership{}, middleware.After)
}

func addOpUpdateProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateProfile{}, middleware.After)
}

func addOpUpdateTransformerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateTransformer{}, middleware.After)
}

func validateCapabilityConfiguration(v types.CapabilityConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CapabilityConfiguration"}
	switch uv := v.(type) {
	case *types.CapabilityConfigurationMemberEdi:
		if err := validateEdiConfiguration(&uv.Value); err != nil {
			invalidParams.AddNested("[edi]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEdiConfiguration(v *types.EdiConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EdiConfiguration"}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.InputLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputLocation"))
	}
	if v.OutputLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputLocation"))
	}
	if v.TransformerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCapabilityInput(v *CreateCapabilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCapabilityInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configuration"))
	} else if v.Configuration != nil {
		if err := validateCapabilityConfiguration(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePartnershipInput(v *CreatePartnershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePartnershipInput"}
	if v.ProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfileId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Email == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Email"))
	}
	if v.Capabilities == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Capabilities"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProfileInput(v *CreateProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProfileInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Phone == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Phone"))
	}
	if v.BusinessName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BusinessName"))
	}
	if len(v.Logging) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Logging"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTransformerInput(v *CreateTransformerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTransformerInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.FileFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FileFormat"))
	}
	if v.MappingTemplate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MappingTemplate"))
	}
	if v.EdiType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EdiType"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCapabilityInput(v *DeleteCapabilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCapabilityInput"}
	if v.CapabilityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CapabilityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePartnershipInput(v *DeletePartnershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePartnershipInput"}
	if v.PartnershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartnershipId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProfileInput(v *DeleteProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProfileInput"}
	if v.ProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfileId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTransformerInput(v *DeleteTransformerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTransformerInput"}
	if v.TransformerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCapabilityInput(v *GetCapabilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCapabilityInput"}
	if v.CapabilityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CapabilityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPartnershipInput(v *GetPartnershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPartnershipInput"}
	if v.PartnershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartnershipId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProfileInput(v *GetProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProfileInput"}
	if v.ProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfileId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTransformerInput(v *GetTransformerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTransformerInput"}
	if v.TransformerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTransformerJobInput(v *GetTransformerJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTransformerJobInput"}
	if v.TransformerJobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerJobId"))
	}
	if v.TransformerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartTransformerJobInput(v *StartTransformerJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTransformerJobInput"}
	if v.InputFile == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputFile"))
	}
	if v.OutputLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputLocation"))
	}
	if v.TransformerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTestMappingInput(v *TestMappingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TestMappingInput"}
	if v.InputFileContent == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputFileContent"))
	}
	if v.MappingTemplate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MappingTemplate"))
	}
	if len(v.FileFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FileFormat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTestParsingInput(v *TestParsingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TestParsingInput"}
	if v.InputFile == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputFile"))
	}
	if len(v.FileFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FileFormat"))
	}
	if v.EdiType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EdiType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCapabilityInput(v *UpdateCapabilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCapabilityInput"}
	if v.CapabilityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CapabilityId"))
	}
	if v.Configuration != nil {
		if err := validateCapabilityConfiguration(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdatePartnershipInput(v *UpdatePartnershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePartnershipInput"}
	if v.PartnershipId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartnershipId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateProfileInput(v *UpdateProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateProfileInput"}
	if v.ProfileId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfileId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateTransformerInput(v *UpdateTransformerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTransformerInput"}
	if v.TransformerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TransformerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
