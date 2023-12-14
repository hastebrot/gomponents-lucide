package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func JapaneseYen(children ...g.Node) g.Node {
	svg := `<path d="M12 9.5V21m0-11.5L6 3m6 6.5L18 3" /><path d="M6 15h12" /><path d="M6 11h12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
