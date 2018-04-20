package styles

import (
	"github.com/alecthomas/chroma"
)

// Hinshun style.
var Hinshun = Register(chroma.MustNewStyle("hinshun", chroma.StyleEntries{
	chroma.LineNumbers:        "#3a3f4b",
	chroma.Comment:            "#526270",
	chroma.Keyword:            "#c679dd",
	chroma.KeywordConstant:    "#d19966",
	chroma.KeywordDeclaration: "#4096dd",
	chroma.KeywordReserved:    "bold #abb2c0",
	chroma.KeywordNamespace:   "#d19966",
	chroma.KeywordType:        "#4096dd",
	chroma.NameBuiltin:        "#d19966",
	chroma.Name:               "#abb2c0",
	chroma.LiteralNumber:      "#d19966",
	chroma.LiteralString:      "#98c373",
	chroma.Operator:           "#8abeb7",
	chroma.Background:         "bg:#282c34",
}))
