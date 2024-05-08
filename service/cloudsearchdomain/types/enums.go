// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ContentType string

// Enum values for ContentType
const (
	ContentTypeApplicationJson ContentType = "application/json"
	ContentTypeApplicationXml  ContentType = "application/xml"
)

// Values returns all known values for ContentType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContentType) Values() []ContentType {
	return []ContentType{
		"application/json",
		"application/xml",
	}
}

type QueryParser string

// Enum values for QueryParser
const (
	QueryParserSimple     QueryParser = "simple"
	QueryParserStructured QueryParser = "structured"
	QueryParserLucene     QueryParser = "lucene"
	QueryParserDismax     QueryParser = "dismax"
)

// Values returns all known values for QueryParser. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryParser) Values() []QueryParser {
	return []QueryParser{
		"simple",
		"structured",
		"lucene",
		"dismax",
	}
}
