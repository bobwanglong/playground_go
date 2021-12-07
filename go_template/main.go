package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	type PathBankend struct {
		PathPrefix string
		Proxy      string
	}

	pathBankend := PathBankend{
		PathPrefix: "/api",
		Proxy:      "apitest",
	}
	pathBankend2 := PathBankend{
		PathPrefix: "/api/bc",
		Proxy:      "apitest2",
	}
	// tpl, err := template.New("test").Parse("{{.Proxy}} is a bankend of {{.PathPrefix}}")
	// checkErr(err)
	// 写入文件
	file, err := os.OpenFile("nginx.conf", os.O_CREATE|os.O_WRONLY, 0755)
	checkErr(err)

	//
	var bankends = struct {
		Fields []PathBankend
	}{
		Fields: []PathBankend{
			pathBankend,
			pathBankend2,
		},
	}
	var Text = `
	{{range .Fields }}
	Proxy: {{.Proxy}} - PathPrefix:{{.PathPrefix}}
	{{ end }}
	`
	tpl, err := template.New("test").Parse(Text)
	checkErr(err)
	err = tpl.Execute(file, bankends)
	checkErr(err)

}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
