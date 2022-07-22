package mysql

import (
	"github.com/yo3jones/gosql/pkg/prtbldr"
	"github.com/yo3jones/gosql/pkg/prtbldr/generic"
)

type (
	// PartBuilderFactory implements prtbldr.SQLPartBuilderFactory for mysql
	// dialect.
	PartBuilderFactory struct {
		genericFactory prtbldr.SQLPartBuilderFactory
	}
)

//nolint:grouper // can't group single variable
var _ prtbldr.SQLPartBuilderFactory = (*PartBuilderFactory)(nil)

// NewPartBuilder returns a new SQLPartBuilder for the given partType
// and part.
//nolint:ireturn//need to return an interface a a builder
func (factory *PartBuilderFactory) NewPartBuilder(
	partType prtbldr.SQLPartType,
	part prtbldr.SQLPart,
) prtbldr.SQLPartBuilder {
	return factory.genericFactory.NewPartBuilder(partType, part)
}

func NewFactory() *PartBuilderFactory {
	return &PartBuilderFactory{
		genericFactory: generic.NewFactory(),
	}
}
