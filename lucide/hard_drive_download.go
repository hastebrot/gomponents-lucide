package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func HardDriveDownload(children ...g.Node) g.Node {
	svg := `<path d="M12 2v8" /><path d="m16 6-4 4-4-4" /><rect width="20" height="8" x="2" y="14" rx="2" /><path d="M6 18h.01" /><path d="M10 18h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
