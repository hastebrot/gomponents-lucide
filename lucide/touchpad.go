package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Touchpad(children ...g.Node) g.Node {
	svg := `<rect width="20" height="16" x="2" y="4" rx="2" /><path d="M2 14h20" /><path d="M12 20v-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
