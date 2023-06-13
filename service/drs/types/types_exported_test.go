// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
)

func ExampleEventResourceData_outputUsage() {
	var union types.EventResourceData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EventResourceDataMemberSourceNetworkData:
		_ = v.Value // Value is types.SourceNetworkData

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SourceNetworkData

func ExampleParticipatingResourceID_outputUsage() {
	var union types.ParticipatingResourceID
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ParticipatingResourceIDMemberSourceNetworkID:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
