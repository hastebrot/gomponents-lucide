package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func KeyboardMusic(children ...g.Node) g.Node {
	svg := `<rect width="20" height="16" x="2" y="4" rx="2" /><path d="M6 8h4" /><path d="M14 8h.01" /><path d="M18 8h.01" /><path d="M2 12h20" /><path d="M6 12v4" /><path d="M10 12v4" /><path d="M14 12v4" /><path d="M18 12v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
