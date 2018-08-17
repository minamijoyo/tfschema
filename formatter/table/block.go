package table

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/olekukonko/tablewriter"
)

// Block is wrapper for tfschema.Block.
// This is a layer for customization for table format.
type Block struct {
	// Simply embedding the structure.
	tfschema.Block
}

// NewBlock creates a new Block instance.
func NewBlock(b *tfschema.Block) *Block {
	return &Block{
		Block: *b,
	}
}

// Format returns a formatted string in table format.
func (b *Block) Format() (string, error) {
	return renderBlock(b)
}

// renderBlock returns a formatted string in table format for Block.
func renderBlock(b *Block) (string, error) {
	buf := new(bytes.Buffer)
	attributes, err := renderAttributes(b)
	if err != nil {
		return "", err
	}
	buf.WriteString(attributes)

	blockTypes, err := renderBlockTypes(b)
	if err != nil {
		return "", err
	}
	buf.WriteString(blockTypes)

	return buf.String(), nil
}

// renderAttributes returns a formatted string in table format for Attributes.
func renderAttributes(b *Block) (string, error) {
	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader([]string{"attribute", "type", "required", "optional", "computed", "sensitive"})

	// sort map keys
	keys := []string{}
	for k := range b.Attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := b.Attributes[k]
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

// renderBlockTypes returns a formatted string in table format for BlockTypes.
func renderBlockTypes(b *Block) (string, error) {
	if len(b.BlockTypes) == 0 {
		return "", nil
	}

	buf := new(bytes.Buffer)

	// sort map keys
	keys := []string{}
	for k := range b.BlockTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := b.BlockTypes[k]
		blockType := fmt.Sprintf("\nblock_type: %s, nesting: %s, min_items: %d, max_items: %d\n", k, v.Nesting, v.MinItems, v.MaxItems)
		buf.WriteString(blockType)

		// We need *Block, not *tfschema.NestedBlock. Convert it here.
		nested := NewBlock(&v.Block)
		block, err := renderBlock(nested)
		if err != nil {
			return "", err
		}

		buf.WriteString(block)
	}

	return buf.String(), nil
}
