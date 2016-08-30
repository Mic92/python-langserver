package python

import (
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/lsp"
)

func offsetForPosition(contents []byte, p lsp.Position) (offset uint64, valid bool) {
	line := 0
	col := 0
	// TODO(sqs): count chars, not bytes, per LSP. does that mean we
	// need to maintain 2 separate counters since we still need to
	// return the offset as bytes?
	for _, b := range contents {
		if line == p.Line && col == p.Character {
			return offset, true
		}
		if line > p.Line || (line == p.Line && col > p.Character) {
			return 0, false
		}

		offset++
		if b == '\n' {
			line++
			col = 0
		} else {
			col++
		}
	}
	return 0, false
}
