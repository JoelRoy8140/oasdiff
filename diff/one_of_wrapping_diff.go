package diff

import (
"sort"

"github.com/getkin/kin-openapi/openapi3"
)

// OneOfWrappingDiff represents the detection of a "oneOf wrapping" pattern.
// This occurs when a concrete object schema (type: object with properties)
// is changed to a oneOf composition containing multiple sub-schemas that
// closely resemble the original object. This is a common API evolution pattern.
type OneOfWrappingDiff struct {
// NumAlternatives is the total number of alternatives in the oneOf
NumAlternatives int `json:"numAlternatives" yaml:"numAlternatives"`
// BestAlternativeIndex is the 0-based index of the oneOf alternative that best matches the base schema
BestAlternativeIndex int `json:"bestAlternativeIndex" yaml:"bestAlternativeIndex"`
// MatchScore is the ratio of type-compatible overlapping properties to base properties (0.0 to 1.0)
MatchScore float64 `json:"matchScore" yaml:"matchScore"`
// RequiredBecameOptional lists properties that were required in the base
// but are not required in ALL alternatives (became effectively optional).
// Sorted alphabetically.
RequiredBecameOptional []string `json:"requiredBecameOptional,omitempty" yaml:"requiredBecameOptional,omitempty"`
// RequiredBecameRequired lists properties that were NOT required in the base
// but ARE required in ALL alternatives (became effectively required).
// Sorted alphabetically.
RequiredBecameRequired []string `json:"requiredBecameRequired,omitempty" yaml:"requiredBecameRequired,omitempty"`
// TypeMismatchedProperties lists properties that exist in both base and best-matching alternative
// but have incompatible types. Sorted alphabetically.
TypeMismatchedProperties []string `json:"typeMismatchedProperties,omitempty" yaml:"typeMismatchedProperties,omitempty"`
// PropertiesAddedToAll lists properties that exist in ALL alternatives but not in base.
// Sorted alphabetically.
PropertiesAddedToAll []string `json:"propertiesAddedToAll,omitempty" yaml:"propertiesAddedToAll,omitempty"`
// PropertiesRemovedFromAll lists properties that existed in base but do not exist
// in ANY of the alternatives. Sorted alphabetically.
PropertiesRemovedFromAll []string `json:"propertiesRemovedFromAll,omitempty" yaml:"propertiesRemovedFromAll,omitempty"`
// DiscriminatorProperty is the name of the discriminator property, if the revision
// oneOf schema has a discriminator. Empty string if no discriminator.
DiscriminatorProperty string `json:"discriminatorProperty,omitempty" yaml:"discriminatorProperty,omitempty"`
}

// Empty indicates whether a change was found in this element
func (diff *OneOfWrappingDiff) Empty() bool {
if diff == nil {
return true
}
return diff.NumAlternatives == 0 &&
len(diff.RequiredBecameOptional) == 0 &&
len(diff.RequiredBecameRequired) == 0 &&
len(diff.TypeMismatchedProperties) == 0 &&
len(diff.PropertiesAddedToAll) == 0 &&
len(diff.PropertiesRemovedFromAll) == 0 &&
diff.DiscriminatorProperty == ""
}

// getOneOfWrappingDiff detects the "oneOf wrapping" pattern and returns
// a populated OneOfWrappingDiff, or nil if no wrapping is detected.
// TODO: Implement the detection algorithm
func getOneOfWrappingDiff(config *Config, state *state, base, revision *openapi3.Schema) (*OneOfWrappingDiff, error) {
_ = sort.Strings // ensure sort import is used
return nil, nil
}