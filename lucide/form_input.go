package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FormInput(children ...g.Node) g.Node {
	svg := `<rect width="20" height="12" x="2" y="6" rx="2" /><path d="M12 12h.01" /><path d="M17 12h.01" /><path d="M7 12h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
