package fonts

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_Load_Base64_Font(t *testing.T) {

	_, err := base64.StdEncoding.DecodeString(FontNeuropol)

	if err != nil {
		fmt.Printf("Panic: %s\n", err.Error())
		panic("Failed to write font Go file")
	}
	// err = ioutil.WriteFile(fontGoFile, []byte(goCode), 0644)

}
