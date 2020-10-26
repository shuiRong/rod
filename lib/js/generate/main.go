package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	list := getList()
	out := "package js\n\n"

	for _, fn := range list.Arr() {
		name := fn.Get("name").Str()
		def := fn.Get("definition").Str()
		out += utils.S(`
			var {{.Name}} = &Function{
				Name: "{{.name}}",
				Definition:   {{.definition}},
				Dependencies: {{.dependencies}},
			}
		`,
			"Name", fnName(name),
			"name", name,
			"definition", utils.EscapeGoString(def),
			"dependencies", getDeps(def),
		)
	}

	utils.E(utils.OutputFile("lib/js/helper.go", out))

	utils.Exec("gofmt", "-s", "-w", "lib/js/helper.go")
}

var regDeps = regexp.MustCompile(`\Wfunctions.(\w+)`)

func getDeps(fn string) string {
	ms := regDeps.FindAllStringSubmatch(fn, -1)

	list := []string{}

	for _, m := range ms {
		list = append(list, fnName(m[1]))
	}

	return "[]*Function{" + strings.Join(list, ",") + "}"
}

func fnName(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func getList() gson.JSON {
	code, err := exec.Command("npx", "-q", "uglify-es", "-c", "-m", "--", "lib/js/helper.js").CombinedOutput()
	if err != nil {
		panic(string(code))
	}

	script := fmt.Sprintf(`
		%s

		const list = []

		for (const name in functions) {
			const reg = new RegExp('^(async )?' + name)
			const definition = functions[name].toString().replace(reg, '$1function')
			list.push({name, definition})
		}

		console.log(JSON.stringify(list))
	`, string(code))

	b, err := exec.Command("node", "-e", script).CombinedOutput()
	if err != nil {
		panic(string(b))
	}

	return gson.New(b)
}
