package main

import (
	"fmt"
	"os"
    "os/exec"
	"path/filepath"
	"text/template"
)

func Build(username string, password string) (err error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return fmt.Errorf("$GOPATH not set")
	}
    dir := filepath.Join(gopath, "src", "github.com", "cgcgbcbc", "gotunet")
    tmpl := template.Must(template.New("constant.go.tmpl").ParseFiles(filepath.Join(dir, "constant.go.tmpl")))
	data := struct {
		Username string
		Password string
	}{username, password}
    output, err := os.Create(filepath.Join(dir, "constant.go"))
    if err != nil {
        return
    }
    err = render(tmpl, data, output)
    if err != nil {
        return
    }
    build := exec.Command("go", "build", "-o", fmt.Sprintf("gotunet-%s", username))
    build.Dir = dir
    err = build.Run()
    if err != nil {
        return
    }
    fmt.Printf("successfully build gotunet-%s\n", username)
    clear := struct {
        Username string
        Password string
    }{"", ""}
    err = output.Close()
    if err != nil {
        return
    }
    output, err = os.Create(filepath.Join(dir, "constant.go"))
    if err != nil {
        return
    }
    err = render(tmpl, clear, output)
	return
}

func render(t *template.Template, data interface{}, output *os.File) (err error) {
    err = t.Execute(output, data)
    return
}
