package prtbldr

import (
	"fmt"
)

type (
	// NotImplementedBuilder is a part builder for handling when a part builder
	// is not implemented.
	NotImplementedBuilder struct {
		part     SQLPart
		partType SQLPartType
	}
)

//nolint:grouper//can't group single variable
var _ SQLPartBuilder = (*NotImplementedBuilder)(nil)

func (builder *NotImplementedBuilder) Build(
	res SQLPartBuilderResult,
	opts ...SQLPartBuilderOption,
) {
	res.Print(opts, "[UNIMPLEMENTED]")
	res.AppendError(
		fmt.Errorf(
			"%w SQLPartBuilder is not implemented. type: %s part: %T",
			ErrUnimplemented,
			builder.partType,
			builder.part,
		),
	)
}
