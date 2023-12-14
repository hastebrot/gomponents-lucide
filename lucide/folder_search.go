package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FolderSearch(children ...g.Node) g.Node {
	svg := `<circle cx="17" cy="17" r="3" /><path d="M10.7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v4.1" /><path d="m21 21-1.5-1.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
