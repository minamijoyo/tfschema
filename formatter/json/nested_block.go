package json

import (
	"github.com/hashicorp/terraform/config/configschema"
	"github.com/minamijoyo/tfschema/tfschema"
)

// NestedBlock is wrapper for tfschema.NestedBlock
type NestedBlock struct {
	// TypeName is a name of block type.
	// Note that this key does not exist in the original data structure.
	// In order to reduce the possibility of future conflicts,
	// naming borrowed from the schema definition of the grpc provider's proto file.
	// https://github.com/hashicorp/terraform/pull/18550
	TypeName string `json:"type_name"`

	// Block is a nested child block.
	Block
	// Nesting is a nesting mode.
	Nesting configschema.NestingMode `json:"nesting"`
	// MinItems is a lower limit on number of nested child blocks.
	MinItems int `json:"min_items"`
	// MaxItems is a upper limit on number of nested child blocks.
	MaxItems int `json:"max_items"`
}

// NewNestedBlock creates a new NestedBlock instance.
func NewNestedBlock(b *tfschema.NestedBlock, typeName string) *NestedBlock {
	block := NewBlock(&b.Block)
	return &NestedBlock{
		TypeName: typeName,
		Block:    *block,
		Nesting:  b.Nesting,
		MinItems: b.MinItems,
		MaxItems: b.MaxItems,
	}
}
