package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MonitorPlay(children ...g.Node) g.Node {
	svg := `<path d="m10 7 5 3-5 3Z" /><rect width="20" height="14" x="2" y="3" rx="2" /><path d="M12 17v4" /><path d="M8 21h8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
