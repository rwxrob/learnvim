package main

import (
	_ "embed"
	"os"

	"github.com/rwxrob/smartvimtutorial/internal"
)

//go:embed tutorial.txt
var tutorial string

//go:embed vimrc
var vimrc string

func main() {

	tutfile, err := os.CreateTemp(``, `smartvimtutorial-`)
	defer tutfile.Close()
	defer os.Remove(tutfile.Name())
	if err != nil {
		os.Exit(1)
	}
	tutfile.WriteString(tutorial)

	vimrcfile, err := os.CreateTemp(``, `smartvimrc-`)
	defer vimrcfile.Close()
	defer os.Remove(vimrcfile.Name())
	if err != nil {
		os.Exit(1)
	}
	vimrcfile.WriteString(vimrc)

	internal.Exec(`vim`, `-u`, vimrcfile.Name(), tutfile.Name())

}
