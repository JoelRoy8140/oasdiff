package diff

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// OneOfWrappingDiff captures detected oneOf wrapping transitions
type OneOfWrappingDiff struct {
	NumAlternatives          int      `json:"numAlternatives" yaml:"numAlternatives"`
	BestAlternativeIndex     int      `json:"bestAlternativeIndex" yaml:"bestAlternativeIndex"`
	MatchScore               float64  `json:"matchScore" yaml:"matchScore"`
	RequiredBecameOptional   []string `json:"requiredBecameOptional,omitempty" yaml:"requiredBecameOptional,omitempty"`
	RequiredBecameRequired   []string `json:"requiredBecameRequired,omitempty" yaml:"requiredBecameRequired,omitempty"`
	TypeMismatchedProperties []string `json:"typeMismatchedProperties,omitempty" yaml:"typeMismatchedProperties,omitempty"`
	PropertiesAddedToAll     []string `json:"propertiesAddedToAll,omitempty" yaml:"propertiesAddedToAll,omitempty"`
	PropertiesRemovedFromAll []string `json:"propertiesRemovedFromAll,omitempty" yaml:"propertiesRemovedFromAll,omitempty"`
	DiscriminatorProperty    string   `json:"discriminatorProperty,omitempty" yaml:"discriminatorProperty,omitempty"`
}

// Empty returns true if no wrapping was detected
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

// getOneOfWrappingDiff detects oneOf wrapping pattern transitions
func getOneOfWrappingDiff(config *Config, state *state, base, revision *openapi3.Schema) (*OneOfWrappingDiff, error) {
	// TODO: implement
	return nil, nil
}
