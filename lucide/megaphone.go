package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Megaphone(children ...g.Node) g.Node {
	svg := `<path d="m3 11 18-5v12L3 14v-3z" /><path d="M11.6 16.8a3 3 0 1 1-5.8-1.6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
