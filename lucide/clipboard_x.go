package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ClipboardX(children ...g.Node) g.Node {
	svg := `<rect width="8" height="4" x="8" y="2" rx="1" ry="1" /><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2" /><path d="m15 11-6 6" /><path d="m9 11 6 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
