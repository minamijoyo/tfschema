package table

import (
	"github.com/minamijoyo/tfschema/tfschema"
)

// NestedBlock is wrapper for tfschema.NestedBlock
type NestedBlock struct {
	// Simply embedding the structure.
	tfschema.NestedBlock
}

// NewNestedBlock creates a new NestedBlock instance.
func NewNestedBlock(b *tfschema.NestedBlock) *NestedBlock {
	return &NestedBlock{
		NestedBlock: *b,
	}
}
