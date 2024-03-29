package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CandlestickChart(children ...g.Node) g.Node {
	svg := `<path d="M9 5v4" /><rect width="4" height="6" x="7" y="9" rx="1" /><path d="M9 15v2" /><path d="M17 3v2" /><rect width="4" height="8" x="15" y="5" rx="1" /><path d="M17 13v3" /><path d="M3 3v18h18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
