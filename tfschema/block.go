package tfschema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/config/configschema"
	"github.com/olekukonko/tablewriter"
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

func (b *Block) FormatTable() (string, error) {
	return b.renderBlock()
}

func (b *Block) renderBlock() (string, error) {
	buf := new(bytes.Buffer)
	attributes, err := b.renderAttributes()
	if err != nil {
		return "", err
	}
	buf.WriteString(attributes)

	blockTypes, err := b.renderBlockTypes()
	if err != nil {
		return "", err
	}
	buf.WriteString(blockTypes)

	return buf.String(), nil
}

func (b *Block) renderAttributes() (string, error) {
	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader([]string{"attribute", "type", "required", "optional", "computed", "sensitive"})
	for k, v := range b.Attributes {
		typeName, err := v.Type.Name()
		if err != nil {
			return "", err
		}

		row := []string{
			k,
			typeName,
			strconv.FormatBool(v.Required),
			strconv.FormatBool(v.Optional),
			strconv.FormatBool(v.Computed),
			strconv.FormatBool(v.Sensitive),
		}
		table.Append(row)
	}

	table.Render()

	return buf.String(), nil
}

func (b *Block) renderBlockTypes() (string, error) {
	if len(b.BlockTypes) == 0 {
		return "", nil
	}

	buf := new(bytes.Buffer)
	for k, v := range b.BlockTypes {
		blockType := fmt.Sprintf("\nblock_type: %s, nesting: %s, min_items: %d, max_items: %d\n", k, v.Nesting, v.MinItems, v.MaxItems)
		buf.WriteString(blockType)

		block, err := v.renderBlock()
		if err != nil {
			return "", err
		}

		buf.WriteString(block)
	}

	return buf.String(), nil
}
