package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func DraftingCompass(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="5" r="2" /><path d="m3 21 8.02-14.26" /><path d="m12.99 6.74 1.93 3.44" /><path d="M19 12c-3.87 4-10.13 4-14 0" /><path d="m21 21-2.16-3.84" />`
	return Icon(g.Group(children), g.Raw(svg))
}
