package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Folders(children ...g.Node) g.Node {
	svg := `<path d="M20 17a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3.9a2 2 0 0 1-1.69-.9l-.81-1.2a2 2 0 0 0-1.67-.9H8a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2Z" /><path d="M2 8v11a2 2 0 0 0 2 2h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
