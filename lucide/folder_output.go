package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FolderOutput(children ...g.Node) g.Node {
	svg := `<path d="M2 7.5V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2" /><path d="M2 13h10" /><path d="m5 10-3 3 3 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
