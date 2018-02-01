package tfschema

import (
	"github.com/hashicorp/terraform/config/configschema"
)

// Block is wrapper for configschema.Block
type Block struct {
	Attributes map[string]*Attribute   `json:"attributes"`
	BlockTypes map[string]*NestedBlock `json:"block_types"`
}

// Attribute is wrapper for configschema.Attribute
type Attribute struct {
	// Type is a type of the attribute's value.
	// Note that Type is not cty.Type
	// We cannot import github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty/cty
	// On the other hand, tfschema does not need a dynamic type.
	// So, we use a simple representation of type.
	Type      Type `json:"type"`
	Required  bool `json:"required"`
	Optional  bool `json:"optional"`
	Computed  bool `json:"computed"`
	Sensitive bool `json:"sensitive"`
}

// NestedBlock is wrapper for configschema.NestedBlock
type NestedBlock struct {
	Block
	Nesting  configschema.NestingMode `json:"nesting"`
	MinItems int                      `json:"min_items"`
	MaxItems int                      `json:"max_items"`
}

func NewBlock(b *configschema.Block) *Block {
	return &Block{
		Attributes: NewAttributes(b.Attributes),
		BlockTypes: NewBlockTypes(b.BlockTypes),
	}
}

func NewAttributes(as map[string]*configschema.Attribute) map[string]*Attribute {
	m := make(map[string]*Attribute)

	for k, v := range as {
		m[k] = NewAttribute(v)
	}

	return m
}

func NewAttribute(a *configschema.Attribute) *Attribute {
	return &Attribute{
		Type:      *NewType(a.Type),
		Required:  a.Required,
		Optional:  a.Optional,
		Computed:  a.Computed,
		Sensitive: a.Sensitive,
	}
}

func NewBlockTypes(bs map[string]*configschema.NestedBlock) map[string]*NestedBlock {
	m := make(map[string]*NestedBlock)

	for k, v := range bs {
		m[k] = NewNestedBlock(v)
	}

	return m
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
