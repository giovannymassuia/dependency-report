package testutils

import (
	"bytes"
	"os/exec"
)

func GetModulePath() string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
		return ""
	}
	path := out.String()
	// remove \n
	path = path[:len(path)-1]
	return path
}
