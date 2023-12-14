package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Heading4(children ...g.Node) g.Node {
	svg := `<path d="M4 12h8" /><path d="M4 18V6" /><path d="M12 18V6" /><path d="M17 10v4h4" /><path d="M21 10v8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
