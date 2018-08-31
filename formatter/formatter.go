package formatter

import (
	"fmt"

	"github.com/minamijoyo/tfschema/formatter/json"
	"github.com/minamijoyo/tfschema/formatter/table"
	"github.com/minamijoyo/tfschema/tfschema"
)

// BlockFormatter is an interface of formatting Block.
// This is an abstraction layer for separating data structure and output format.
type BlockFormatter interface {
	Format() (string, error)
}

// NewBlockFormatter is a factory method which returns a BlockFormatter interface.
func NewBlockFormatter(b *tfschema.Block, format string) (BlockFormatter, error) {
	var f BlockFormatter

	switch format {
	case "table":
		f = table.NewBlock(b)
	case "json":
		f = json.NewBlock(b)
	default:
		return nil, fmt.Errorf("Failed to new BlockFormatter. Unknown output format: %s", format)
	}

	return f, nil
}
