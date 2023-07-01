package main

import (
	"embed"
	_ "embed"
	"log"
	"os"
	"text/template"

	"github.com/rwxrob/learnvim/internal"
)

//go:embed all:files
var f embed.FS

func main() {

	// create custom vimrc for this session
	vimrcfile, err := os.CreateTemp(``, `learnvim-`)
	defer vimrcfile.Close()
	defer os.Remove(vimrcfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	vimrc, err := f.ReadFile(`files/vimrc`)
	if err != nil {
		log.Fatal(err)
	}
	vimrcfile.Write(vimrc)

	// create the temporary file and clean up
	tutfile, err := os.CreateTemp(``, `learnvim-`)
	defer tutfile.Close()
	defer os.Remove(tutfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	custom := template.FuncMap{
		`has`:    internal.Has,
		`goos`:   internal.GOOS,
		`goarch`: internal.GOARCH,
	}

	data := map[string]any{`this`: `that`}

	t, err := template.New(`tutorial.txt`).Funcs(custom).ParseFS(f, `files/*.txt`)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(tutfile, data); err != nil {
		log.Fatal(err)
	}

	// TODO compile the template
	//tutorial, err := f.ReadFile(`files/tutorial.txt`)

	// hand off execution to vim itself using our custom vimrc
	internal.Exec(`vim`, `-u`, vimrcfile.Name(), tutfile.Name())

}
