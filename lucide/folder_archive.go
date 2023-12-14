package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FolderArchive(children ...g.Node) g.Node {
	svg := `<circle cx="15" cy="19" r="2" /><path d="M20.9 19.8A2 2 0 0 0 22 18V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h5.1" /><path d="M15 11v-1" /><path d="M15 17v-2" />`
	return Icon(g.Group(children), g.Raw(svg))
}