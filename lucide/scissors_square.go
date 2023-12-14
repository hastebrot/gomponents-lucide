package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ScissorsSquare(children ...g.Node) g.Node {
	svg := `<rect width="20" height="20" x="2" y="2" rx="2" /><circle cx="8" cy="8" r="2" /><path d="M9.414 9.414 12 12" /><path d="M14.8 14.8 18 18" /><circle cx="8" cy="16" r="2" /><path d="m18 6-8.586 8.586" />`
	return Icon(g.Group(children), g.Raw(svg))
}
