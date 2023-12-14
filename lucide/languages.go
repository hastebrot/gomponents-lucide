package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Languages(children ...g.Node) g.Node {
	svg := `<path d="m5 8 6 6" /><path d="m4 14 6-6 2-3" /><path d="M2 5h12" /><path d="M7 2h1" /><path d="m22 22-5-10-5 10" /><path d="M14 18h6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
