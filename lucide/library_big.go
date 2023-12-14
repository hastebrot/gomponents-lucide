package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func LibraryBig(children ...g.Node) g.Node {
	svg := `<rect width="8" height="18" x="3" y="3" rx="1" /><path d="M7 3v18" /><path d="M20.4 18.9c.2.5-.1 1.1-.6 1.3l-1.9.7c-.5.2-1.1-.1-1.3-.6L11.1 5.1c-.2-.5.1-1.1.6-1.3l1.9-.7c.5-.2 1.1.1 1.3.6Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
