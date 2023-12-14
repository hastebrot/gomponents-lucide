package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tv2(children ...g.Node) g.Node {
	svg := `<path d="M7 21h10" /><rect width="20" height="14" x="2" y="3" rx="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
