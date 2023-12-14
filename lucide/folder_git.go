package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FolderGit(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="13" r="2" /><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z" /><path d="M14 13h3" /><path d="M7 13h3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
