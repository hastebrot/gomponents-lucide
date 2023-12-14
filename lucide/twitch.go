package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Twitch(children ...g.Node) g.Node {
	svg := `<path d="M21 2H3v16h5v4l4-4h5l4-4V2zm-10 9V7m5 4V7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
