package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Signpost(children ...g.Node) g.Node {
	svg := `<path d="M12 3v3" /><path d="M18.5 13h-13L2 9.5 5.5 6h13L22 9.5Z" /><path d="M12 13v8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
