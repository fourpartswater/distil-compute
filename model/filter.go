package model

import (
	"sort"
)

const (
	// DefaultFilterSize represents the default filter search size.
	DefaultFilterSize = 100
	// FilterSizeLimit represents the largest filter size.
	FilterSizeLimit = 1000
	// CategoricalFilter represents a categorical filter type.
	CategoricalFilter = "categorical"
	// NumericalFilter represents a numerical filter type.
	NumericalFilter = "numerical"
	// FeatureFilter represents a feature filter type.
	FeatureFilter = "feature"
	// TextFilter represents a text filter type.
	TextFilter = "text"
	// RowFilter represents a numerical filter type.
	RowFilter = "row"
	// IncludeFilter represents an inclusive filter mode.
	IncludeFilter = "include"
	// ExcludeFilter represents an exclusive filter mode.
	ExcludeFilter = "exclude"
)

func stringSliceEqual(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Filter defines a variable filter.
type Filter struct {
	Key        string   `json:"key"`
	Type       string   `json:"type"`
	Mode       string   `json:"mode"`
	Min        *float64 `json:"min"`
	Max        *float64 `json:"max"`
	Categories []string `json:"categories"`
	D3mIndices []string `json:"d3mIndices"`
}

// NewNumericalFilter instantiates a numerical filter.
func NewNumericalFilter(key string, mode string, min float64, max float64) *Filter {
	return &Filter{
		Key:  key,
		Type: NumericalFilter,
		Mode: mode,
		Min:  &min,
		Max:  &max,
	}
}

// NewCategoricalFilter instantiates a categorical filter.
func NewCategoricalFilter(key string, mode string, categories []string) *Filter {
	sort.Strings(categories)
	return &Filter{
		Key:        key,
		Type:       CategoricalFilter,
		Mode:       mode,
		Categories: categories,
	}
}

// NewFeatureFilter instantiates a feature filter.
func NewFeatureFilter(key string, mode string, categories []string) *Filter {
	sort.Strings(categories)
	return &Filter{
		Key:        key,
		Type:       FeatureFilter,
		Mode:       mode,
		Categories: categories,
	}
}

// NewTextFilter instantiates a text filter.
func NewTextFilter(key string, mode string, categories []string) *Filter {
	sort.Strings(categories)
	return &Filter{
		Key:        key,
		Type:       TextFilter,
		Mode:       mode,
		Categories: categories,
	}
}

// NewRowFilter instantiates a row filter.
func NewRowFilter(mode string, d3mIndices []string) *Filter {
	return &Filter{
		Type:       RowFilter,
		Mode:       mode,
		D3mIndices: d3mIndices,
	}
}