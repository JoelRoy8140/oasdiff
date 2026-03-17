package checker

import (
	"github.com/oasdiff/oasdiff/diff"
)

// RequestOneOfWrappingChangedCheck detects oneOf wrapping patterns in request bodies and properties
func RequestOneOfWrappingChangedCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap, config *Config) Changes {
	result := make(Changes, 0)
	if diffReport.PathsDiff == nil {
		return result
	}
	for path, pathItem := range diffReport.PathsDiff.Modified {
		if pathItem.OperationsDiff == nil {
			continue
		}
		for operation, operationItem := range pathItem.OperationsDiff.Modified {
			if operationItem.RequestBodyDiff == nil ||
				operationItem.RequestBodyDiff.ContentDiff == nil ||
				operationItem.RequestBodyDiff.ContentDiff.MediaTypeModified == nil {
				continue
			}

			modifiedMediaTypes := operationItem.RequestBodyDiff.ContentDiff.MediaTypeModified
			for mediaType, mediaTypeDiff := range modifiedMediaTypes {
				if mediaTypeDiff.SchemaDiff == nil {
					continue
				}

				// Check request body schema for oneOf wrapping
				changes := checkBodyOneOfWrappingChange(
					mediaTypeDiff.SchemaDiff,
					mediaType,
					"", // responseStatus not applicable for requests
					config,
					operationsSources,
					operationItem,
					operation,
					path,
					true, // isRequest
				)
				result = append(result, changes...)

				// Check request body properties for oneOf wrapping
				CheckModifiedPropertiesDiff(
					mediaTypeDiff.SchemaDiff,
					func(propertyPath string, propertyName string, propertyDiff *diff.SchemaDiff, parent *diff.SchemaDiff) {
						if propertyDiff.Revision == nil {
							return
						}
						if propertyDiff.Revision.ReadOnly {
							return
						}

						changes := checkPropertyOneOfWrappingChange(
							propertyPath,
							propertyName,
							propertyDiff,
							mediaType,
							"", // responseStatus not applicable for requests
							config,
							operationsSources,
							operationItem,
							operation,
							path,
							true, // isRequest
						)
						result = append(result, changes...)
					})
			}
		}
	}
	return result
}
