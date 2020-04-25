package utilities

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Font2Base64(t *testing.T) {
	// ----------------------------------------------------------
	// A simple utility to convert a binary ttf file into a base64
	// Go variable. This allows you to embed the font in the code.
	// ----------------------------------------------------------

	var bytes []byte

	fontFile := "./neuropol x rg.ttf" // input
	fontVariableName := "FontNeuropol"

	fmt.Printf("Reading binary font: %s\n", fontFile)

	fontGoFile := "../ranger/fonts/neuropol_font.go" // output

	bytes, err := ioutil.ReadFile(fontFile)

	if err != nil {
		panic("Font file not found")
	}

	fontEnc := base64.StdEncoding.EncodeToString(bytes)

	fontDoc := "// " + fontVariableName + " is an base64 encoded font.\n"

	goCode := "package fonts\n\n" + fontDoc + "const " + fontVariableName + " = `" + fontEnc + "`"

	_, err = os.Stat(fontGoFile)

	if os.IsExist(err) {
		fmt.Printf("Removing : %s\n", fontGoFile)
		err = os.Remove(fontGoFile)

		if err != nil {
			panic("Unable to remove current font go file")
		}
	}

	fmt.Printf("Writing encoded font: %s\n", fontGoFile)
	err = ioutil.WriteFile(fontGoFile, []byte(goCode), 0644)

	if err != nil {
		panic("Failed to write font Go file")
	}
}
