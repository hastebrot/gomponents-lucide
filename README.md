# gomponents-lucide

[Lucide](https://lucide.dev/) for [gomponents](https://www.gomponents.com/).

**build instructions:**

```
❯ curl -sSL https://github.com/lucide-icons/lucide/releases/download/0.295.0/lucide-icons-0.295.0.zip | tar -xf - -C .

❯ ls -1 icons/*.svg | wc -l
1346

❯ deno run --allow-read --allow-write build.deno.ts

❯ ls -1 lucide/*.go | wc -l
1346
```
