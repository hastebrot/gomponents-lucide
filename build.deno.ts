#!/usr/bin/env -S deno run --allow-read --allow-write

const template = (name: string, svg: string) => `
package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ${name}(children ...g.Node) g.Node {
	svg := \`${svg}\`
	return Icon(g.Group(children), g.Raw(svg))
}
`;

const pathSourceSvg = "./icons/";
const pathTargetGo = "./lucide/";

await Deno.mkdir(pathTargetGo, { recursive: true });

for await (const entry of Deno.readDir(pathSourceSvg)) {
  if (!entry.name.endsWith(".svg")) {
    continue;
  }

  const svgElement = await Deno.readTextFile(pathSourceSvg + entry.name);
  const svgChildren = svgElement.slice(
    svgElement.indexOf(">") + 1,
    svgElement.lastIndexOf("<")
  );
  const svg = svgChildren
    .split("\n")
    .map((it) => it.trim())
    .join("");

  const name = entry.name
    .slice(0, entry.name.lastIndexOf("."))
    .replaceAll("-", "_");
  const goName = name
    .split("_")
    .map((it) => it.slice(0, 1).toUpperCase() + it.slice(1).toLowerCase())
    .join("");
  const goFilename = name + ".go";

  const goSource = template(goName, svg).trimStart();
  await Deno.writeTextFile(pathTargetGo + goFilename, goSource);
  await Deno.stdout.write(new TextEncoder().encode("."));
}
