package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FolderSearch2(children ...g.Node) g.Node {
	svg := `<circle cx="11.5" cy="12.5" r="2.5" /><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z" /><path d="M13.3 14.3 15 16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
