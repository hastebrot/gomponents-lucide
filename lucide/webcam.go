package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Webcam(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="10" r="8" /><circle cx="12" cy="10" r="3" /><path d="M7 22h10" /><path d="M12 22v-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
