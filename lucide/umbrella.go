package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Umbrella(children ...g.Node) g.Node {
	svg := `<path d="M22 12a10.06 10.06 1 0 0-20 0Z" /><path d="M12 12v8a2 2 0 0 0 4 0" /><path d="M12 2v1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
