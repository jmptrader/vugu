package devutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestTinygoCompiler(t *testing.T) {

	tmpDir, err := ioutil.TempDir("", "TestTinygoCompiler")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	must(ioutil.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(`module example.org/testtgc`), 0644))
	must(ioutil.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(`package main 

import "fmt"
import "github.com/vugu/html"

func main() {
	fmt.Printf("testapp says hello! %v\n", html.NodeType(1))
}
`), 0644))

	tgc := MustNewTinygoCompiler()
	defer tgc.Close()
	tgc.AddGoGet("go get github.com/vugu/html")
	tgc.SetBuildDir(tmpDir)

	outpath, err := tgc.Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outpath)
	b, err := ioutil.ReadFile(outpath)
	if err != nil {
		t.Fatal(err)
	}
	if string(b[:4]) != "\x00asm" {
		t.Fatalf("bad asm magic num: %X", b[:4])
	}
	t.Logf("output file length: %d", len(b))

	r, err := tgc.WasmExecJS()
	if err != nil {
		t.Fatal(err)
	}
	b, err = ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	// t.Logf("wasm_exec.js contents:\n%s", b)
	if !bytes.Contains(b, []byte(`global.Go`)) {
		t.Fatalf("unable to find global.Go in wasm_exec.js")
	}
	if !bytes.Contains(b, []byte(`TinyGo`)) {
		t.Fatalf("unable to find TinyGo in wasm_exec.js")
	}

}
