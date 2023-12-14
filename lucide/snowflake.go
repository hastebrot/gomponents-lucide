package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Snowflake(children ...g.Node) g.Node {
	svg := `<line x1="2" x2="22" y1="12" y2="12" /><line x1="12" x2="12" y1="2" y2="22" /><path d="m20 16-4-4 4-4" /><path d="m4 8 4 4-4 4" /><path d="m16 4-4 4-4-4" /><path d="m8 20 4-4 4 4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
