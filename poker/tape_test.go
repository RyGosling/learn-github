package poker

import (
	"testing"
	"io/ioutil"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTmpFile(t, "12345")
	defer clean()

	tape := &tape{file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}