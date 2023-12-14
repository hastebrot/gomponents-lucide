package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CloudSunRain(children ...g.Node) g.Node {
	svg := `<path d="M12 2v2" /><path d="m4.93 4.93 1.41 1.41" /><path d="M20 12h2" /><path d="m19.07 4.93-1.41 1.41" /><path d="M15.947 12.65a4 4 0 0 0-5.925-4.128" /><path d="M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24" /><path d="M11 20v2" /><path d="M7 19v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
