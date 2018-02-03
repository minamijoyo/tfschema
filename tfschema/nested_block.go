package tfschema

import (
	"github.com/hashicorp/terraform/config/configschema"
)

// NestedBlock is wrapper for configschema.NestedBlock
type NestedBlock struct {
	Block
	Nesting  configschema.NestingMode `json:"nesting"`
	MinItems int                      `json:"min_items"`
	MaxItems int                      `json:"max_items"`
}

func NewNestedBlock(b *configschema.NestedBlock) *NestedBlock {
	block := NewBlock(&b.Block)
	return &NestedBlock{
		Block:    *block,
		Nesting:  b.Nesting,
		MinItems: b.MinItems,
		MaxItems: b.MaxItems,
	}
}
