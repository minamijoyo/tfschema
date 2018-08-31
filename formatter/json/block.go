package json

import (
	"encoding/json"
	"sort"

	"github.com/minamijoyo/tfschema/tfschema"
)

// Block is wrapper for tfschema.Block.
// This is a layer for customization for JSON format.
// Although the original data structure is a map,
// it is hard to parse because the names of the key are not predictable.
// We use a slice here to format easy-to-parse JSON.
type Block struct {
	// Attributes is a slice of any attributes.
	Attributes []*Attribute `json:"attributes"`
	// BlockTypes is a slice of any nested block types.
	BlockTypes []*NestedBlock `json:"block_types"`
}

// NewBlock creates a new Block instance.
func NewBlock(b *tfschema.Block) *Block {
	return &Block{
		Attributes: NewAttributes(b.Attributes),
		BlockTypes: NewBlockTypes(b.BlockTypes),
	}
}

// NewAttributes creates a new slice of Attributes.
func NewAttributes(as map[string]*tfschema.Attribute) []*Attribute {
	m := make([]*Attribute, 0)

	// sort map keys
	keys := []string{}
	for k := range as {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		attr := NewAttribute(as[k], k)
		m = append(m, attr)
	}

	return m
}

// NewBlockTypes creates a new slice of NestedBlocks.
func NewBlockTypes(bs map[string]*tfschema.NestedBlock) []*NestedBlock {
	m := make([]*NestedBlock, 0)

	// sort map keys
	keys := []string{}
	for k := range bs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		nestedBlock := NewNestedBlock(bs[k], k)
		m = append(m, nestedBlock)
	}

	return m
}

// Format returns a formatted string in JSON format.
func (b *Block) Format() (string, error) {
	bytes, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
