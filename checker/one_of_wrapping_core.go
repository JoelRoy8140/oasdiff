package checker

import (
	"github.com/oasdiff/oasdiff/diff"
)

// Message IDs for request body oneOf wrapping changes
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
)

// Message IDs for response body oneOf wrapping changes
const (
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

// checkBodyOneOfWrappingChange checks if a body schema has a oneOf wrapping diff
// and emits appropriate changes
// TODO: Implement checker logic
func checkBodyOneOfWrappingChange(schemaDiff *diff.SchemaDiff, mediaType string, responseStatus string,
	config *Config, operationsSources *diff.OperationsSourcesMap, operationItem *diff.MethodDiff,
	operation string, path string, isRequest bool) Changes {
	// TODO: Implement
	return make(Changes, 0)
}

// checkPropertyOneOfWrappingChange checks if a property has a oneOf wrapping diff
// TODO: Implement checker logic
func checkPropertyOneOfWrappingChange(propertyPath string, propertyName string, propertyDiff *diff.SchemaDiff,
	mediaType string, responseStatus string, config *Config, operationsSources *diff.OperationsSourcesMap,
	operationItem *diff.MethodDiff, operation string, path string, isRequest bool) Changes {
	// TODO: Implement
	return make(Changes, 0)
}

// Suppression functions - called by existing checkers to avoid false positives

// shouldSuppressTypeChangedForOneOfWrapping checks if a type change should be suppressed
func shouldSuppressTypeChangedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
	// TODO: Implement
	return false
}

// shouldSuppressPropertyTypeChangedForOneOfWrapping checks if a property type change should be suppressed
func shouldSuppressPropertyTypeChangedForOneOfWrapping(propertyDiff *diff.SchemaDiff) bool {
	// TODO: Implement
	return false
}

// shouldSuppressOneOfSchemaChangedForOneOfWrapping checks if oneOf/anyOf schema changes should be suppressed
func shouldSuppressOneOfSchemaChangedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
	// TODO: Implement
	return false
}

// shouldSuppressPropertyOneOfSchemaChangedForOneOfWrapping checks if property oneOf/anyOf changes should be suppressed
func shouldSuppressPropertyOneOfSchemaChangedForOneOfWrapping(propertyDiff *diff.SchemaDiff) bool {
	// TODO: Implement
	return false
}

// shouldSuppressPropertyRemovedForOneOfWrapping checks if a property removal should be suppressed
func shouldSuppressPropertyRemovedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
	// TODO: Implement
	return false
}
