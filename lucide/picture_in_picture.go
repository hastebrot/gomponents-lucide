package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func PictureInPicture(children ...g.Node) g.Node {
	svg := `<path d="M8 4.5v5H3m-1-6 6 6m13 0v-3c0-1.16-.84-2-2-2h-7m-9 9v2c0 1.05.95 2 2 2h3" /><rect width="10" height="7" x="12" y="13.5" ry="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
