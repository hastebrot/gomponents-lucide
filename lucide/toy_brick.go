package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ToyBrick(children ...g.Node) g.Node {
	svg := `<rect width="18" height="12" x="3" y="8" rx="1" /><path d="M10 8V5c0-.6-.4-1-1-1H6a1 1 0 0 0-1 1v3" /><path d="M19 8V5c0-.6-.4-1-1-1h-3a1 1 0 0 0-1 1v3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
