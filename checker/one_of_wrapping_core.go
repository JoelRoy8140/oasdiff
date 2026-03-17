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
// Type mismatch messages
RequestBodyOneOfWrappingTypeMismatchId     = "request-body-one-of-wrapping-type-mismatch"
RequestPropertyOneOfWrappingTypeMismatchId = "request-property-one-of-wrapping-type-mismatch"
// Property added to all alternatives
RequestBodyOneOfWrappingPropertyAddedId     = "request-body-one-of-wrapping-property-added"
RequestPropertyOneOfWrappingPropertyAddedId = "request-property-one-of-wrapping-property-added"
// Property removed from all alternatives
RequestBodyOneOfWrappingPropertyRemovedId     = "request-body-one-of-wrapping-property-removed"
RequestPropertyOneOfWrappingPropertyRemovedId = "request-property-one-of-wrapping-property-removed"
)

// Message IDs for response body oneOf wrapping changes
const (
ResponseBodyOneOfWrappingPropertyBecameOptionalId     = "response-body-one-of-wrapping-property-became-optional"
ResponseBodyOneOfWrappingPropertyBecameRequiredId     = "response-body-one-of-wrapping-property-became-required"
ResponsePropertyOneOfWrappingPropertyBecameOptionalId = "response-property-one-of-wrapping-property-became-optional"
ResponsePropertyOneOfWrappingPropertyBecameRequiredId = "response-property-one-of-wrapping-property-became-required"
// Type mismatch messages
ResponseBodyOneOfWrappingTypeMismatchId     = "response-body-one-of-wrapping-type-mismatch"
ResponsePropertyOneOfWrappingTypeMismatchId = "response-property-one-of-wrapping-type-mismatch"
// Property added to all alternatives
ResponseBodyOneOfWrappingPropertyAddedId     = "response-body-one-of-wrapping-property-added"
ResponsePropertyOneOfWrappingPropertyAddedId = "response-property-one-of-wrapping-property-added"
// Property removed from all alternatives
ResponseBodyOneOfWrappingPropertyRemovedId     = "response-body-one-of-wrapping-property-removed"
ResponsePropertyOneOfWrappingPropertyRemovedId = "response-property-one-of-wrapping-property-removed"
)

// TODO: Implement helper functions for checking oneOf wrapping changes
// and suppressing legacy checker messages when wrapping is detected.

// checkBodyOneOfWrappingChange checks if a body schema has a oneOf wrapping diff
// and emits appropriate changes. Returns empty Changes if no wrapping detected.
func checkBodyOneOfWrappingChange(schemaDiff *diff.SchemaDiff, mediaType string, responseStatus string,
config *Config, operationsSources *diff.OperationsSourcesMap, operationItem *diff.MethodDiff,
operation string, path string, isRequest bool) Changes {
// TODO: Implement - iterate over OneOfWrappingDiff fields and emit ApiChanges
return make(Changes, 0)
}

// checkPropertyOneOfWrappingChange checks if a property has a oneOf wrapping diff
// and emits appropriate changes for nested property oneOf wrapping.
func checkPropertyOneOfWrappingChange(propertyPath string, propertyName string, propertyDiff *diff.SchemaDiff,
mediaType string, responseStatus string, config *Config, operationsSources *diff.OperationsSourcesMap,
operationItem *diff.MethodDiff, operation string, path string, isRequest bool) Changes {
// TODO: Implement - same as checkBodyOneOfWrappingChange but for nested properties
return make(Changes, 0)
}

// Suppression functions - called by existing checkers to avoid false positives
// TODO: These stubs must be implemented to return true when oneOf wrapping is detected

func shouldSuppressTypeChangedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
return false // TODO: Implement
}

func shouldSuppressPropertyTypeChangedForOneOfWrapping(propertyDiff *diff.SchemaDiff) bool {
return false // TODO: Implement
}

func shouldSuppressOneOfSchemaChangedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
return false // TODO: Implement
}

func shouldSuppressPropertyOneOfSchemaChangedForOneOfWrapping(propertyDiff *diff.SchemaDiff) bool {
return false // TODO: Implement
}

func shouldSuppressPropertyRemovedForOneOfWrapping(schemaDiff *diff.SchemaDiff) bool {
return false // TODO: Implement
}