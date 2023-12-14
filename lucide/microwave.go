package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Microwave(children ...g.Node) g.Node {
	svg := `<rect width="20" height="15" x="2" y="4" rx="2" /><rect width="8" height="7" x="6" y="8" rx="1" /><path d="M18 8v7" /><path d="M6 19v2" /><path d="M18 19v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
