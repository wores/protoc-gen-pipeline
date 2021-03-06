package shared

import (
	"text/template"

	"github.com/lyft/protoc-gen-star"
)

func RegisterFunctions(tpl *template.Template, params pgs.Parameters) {
	tpl.Funcs(map[string]interface{}{
		"context":  rulesContext,
		"render":   Render(tpl),
		"has":      Has,
		//"needs":    Needs,
	})
}
