package checker

import (
	"github.com/oasdiff/oasdiff/diff"
)

// Message IDs for oneOf wrapping changes
const (
	RequestBodyOneOfWrappingPropertyBecameOptionalId     = "request-body-one-of-wrapping-property-became-optional"
	RequestBodyOneOfWrappingPropertyBecameRequiredId     = "request-body-one-of-wrapping-property-became-required"
	RequestPropertyOneOfWrappingPropertyBecameOptionalId = "request-property-one-of-wrapping-property-became-optional"
	RequestPropertyOneOfWrappingPropertyBecameRequiredId = "request-property-one-of-wrapping-property-became-required"
	RequestBodyOneOfWrappingTypeMismatchId               = "request-body-one-of-wrapping-type-mismatch"
	RequestPropertyOneOfWrappingTypeMismatchId           = "request-property-one-of-wrapping-type-mismatch"
	RequestBodyOneOfWrappingPropertyAddedId              = "request-body-one-of-wrapping-property-added"
	RequestPropertyOneOfWrappingPropertyAddedId          = "request-property-one-of-wrapping-property-added"
	RequestBodyOneOfWrappingPropertyRemovedId            = "request-body-one-of-wrapping-property-removed"
	RequestPropertyOneOfWrappingPropertyRemovedId        = "request-property-one-of-wrapping-property-removed"

	ResponseBodyOneOfWrappingPropertyBecameOptionalId     = "response-body-one-of-wrapping-property-became-optional"
	ResponseBodyOneOfWrappingPropertyBecameRequiredId     = "response-body-one-of-wrapping-property-became-required"
	ResponsePropertyOneOfWrappingPropertyBecameOptionalId = "response-property-one-of-wrapping-property-became-optional"
	ResponsePropertyOneOfWrappingPropertyBecameRequiredId = "response-property-one-of-wrapping-property-became-required"
	ResponseBodyOneOfWrappingTypeMismatchId               = "response-body-one-of-wrapping-type-mismatch"
	ResponsePropertyOneOfWrappingTypeMismatchId           = "response-property-one-of-wrapping-type-mismatch"
	ResponseBodyOneOfWrappingPropertyAddedId              = "response-body-one-of-wrapping-property-added"
	ResponsePropertyOneOfWrappingPropertyAddedId          = "response-property-one-of-wrapping-property-added"
	ResponseBodyOneOfWrappingPropertyRemovedId            = "response-body-one-of-wrapping-property-removed"
	ResponsePropertyOneOfWrappingPropertyRemovedId        = "response-property-one-of-wrapping-property-removed"
)

func checkBodyOneOfWrappingChange(schemaDiff *diff.SchemaDiff, mediaType string, responseStatus string,
	config *Config, operationsSources *diff.OperationsSourcesMap, operationItem *diff.MethodDiff,
	operation string, path string, isRequest bool) Changes {
	return make(Changes, 0)
}

func checkPropertyOneOfWrappingChange(propertyPath string, propertyName string, propertyDiff *diff.SchemaDiff,
	mediaType string, responseStatus string, config *Config, operationsSources *diff.OperationsSourcesMap,
	operationItem *diff.MethodDiff, operation string, path string, isRequest bool) Changes {
	return make(Changes, 0)
}

func shouldSuppressTypeChangedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
	return false
}

func shouldSuppressPropertyTypeChangedForOneOfWrapping(propertyDiff *diff.SchemaDiff) bool {
	return false
}

func shouldSuppressOneOfSchemaChangedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
	return false
}

func shouldSuppressPropertyOneOfSchemaChangedForOneOfWrapping(propertyDiff *diff.SchemaDiff) bool {
	return false
}

func shouldSuppressPropertyRemovedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
	return false
}
