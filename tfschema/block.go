package tfschema

import (
	"encoding/json"

	"github.com/hashicorp/terraform/config/configschema"
)

// Block is wrapper for configschema.Block
type Block struct {
	Attributes map[string]*Attribute   `json:"attributes"`
	BlockTypes map[string]*NestedBlock `json:"block_types"`
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

func NewBlockTypes(bs map[string]*configschema.NestedBlock) map[string]*NestedBlock {
	m := make(map[string]*NestedBlock)

	for k, v := range bs {
		m[k] = NewNestedBlock(v)
	}

	return m
}

func (b *Block) FormatJSON() (string, error) {
	bytes, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
