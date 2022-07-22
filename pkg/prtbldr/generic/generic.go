// package generic is not meant to be used directly.
//
// It has generic logic for building SQL parts that can be reused by dialect
// specific builders.
package generic

import (
	"github.com/yo3jones/gosql/pkg/prtbldr"
)

type (
	// PartBuilderFactory implements prtbldr.SQLPartBuilderFactory for generic
	// dialect.
	PartBuilderFactory struct{}
)

//nolint:grouper // can't group single variable
var _ prtbldr.SQLPartBuilderFactory = (*PartBuilderFactory)(nil)

// NewPartBuilder returns the generic prtbldr.SQLPartBuilder for the given
// partType and part.
//
//nolint:ireturn//need to return an interface a a builder
func (*PartBuilderFactory) NewPartBuilder(
	partType prtbldr.SQLPartType,
	part prtbldr.SQLPart,
) prtbldr.SQLPartBuilder {
	factories := []prtbldr.SQLPartBuilderFactory{
		&CommonFactory{},
		&CreateTableFactory{},
	}

	for _, factory := range factories {
		if builder := factory.NewPartBuilder(partType, part); builder != nil {
			return builder
		}
	}

	return nil
}

// NewFactory returns a new generic prtbldr.SQLPartBuilderFactory.
func NewFactory() *PartBuilderFactory {
	return &PartBuilderFactory{}
}
