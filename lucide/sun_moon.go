package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SunMoon(children ...g.Node) g.Node {
	svg := `<path d="M12 8a2.83 2.83 0 0 0 4 4 4 4 0 1 1-4-4" /><path d="M12 2v2" /><path d="M12 20v2" /><path d="m4.9 4.9 1.4 1.4" /><path d="m17.7 17.7 1.4 1.4" /><path d="M2 12h2" /><path d="M20 12h2" /><path d="m6.3 17.7-1.4 1.4" /><path d="m19.1 4.9-1.4 1.4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
