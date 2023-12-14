package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Banknote(children ...g.Node) g.Node {
	svg := `<rect width="20" height="12" x="2" y="6" rx="2" /><circle cx="12" cy="12" r="2" /><path d="M6 12h.01M18 12h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
