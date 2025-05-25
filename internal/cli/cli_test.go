package cli

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/kmattix/cccondense/internal/condenser"
	"github.com/kmattix/cccondense/internal/utils"
)

func TestRun(t *testing.T) {
	/*
		Usage:
			cccondense [OPTION]... FILE [DEST]
			cccondense [OPTION]... DIRECTORY [DEST]
		Condenses closed-caption file(s).
		When DEST is not specified, the program saves a new file in the source directory with the name
		format FILE-condensed.EXTENSION

		OPTION
			-r, --recursive		Condense all closed-caption files in the directory and its subdirectories.
			-v, --verbose		Display progress information.
			-w, --write			Overwrite the original file with the condensed version.
	*/
}

func TestValidateExtensionNone(t *testing.T) {
	data := ""
	actual := validateExtension(data)
	expected := false
	if actual != expected {
		t.Errorf("validateExtension(%s) = %v; want %v", data, actual, expected)
	}
}

func TestValidateExtensionTrue(t *testing.T) {
	data := "/path/to/my/.secret/cc.srt"
	actual := validateExtension(data)
	expected := true
	if actual != expected {
		t.Errorf("validateExtension(%s) = %v; want %v", data, actual, expected)
	}
}

func TestValidateExtensionFalse(t *testing.T) {
	data := "/path/to/my/.secret/audio.wav"
	actual := validateExtension(data)
	expected := false
	if actual != expected {
		t.Errorf("validateExtension(%s) = %v; want %v", data, actual, expected)
	}
}

func TestHandleFile(t *testing.T) {
	path := "testdata/test.srt"
	expected := "testdata/test" + fileSuffix + ".srt"
	handleFile(path, "", false)
	err := os.Remove(expected)
	if err != nil {
		t.Errorf("handleFile(%s) = %v; want %s", path, err, expected)
	}
}

func TestHandleFileDestFile(t *testing.T) {
	path := "testdata/test.srt"
	expected := "testdata/somedir/anotherfilename.srt"
	handleFile(path, expected, false)
	err := os.Remove(expected)
	if err != nil {
		t.Errorf("handleFile(%s) = %v; want %s", path, err, expected)
	}
}

func TestHandleFileDestDir(t *testing.T) {
	path := "testdata/test.srt"
	outPath := "testdata/somedir/"
	expected := "testdata/somedir/test.srt"
	handleFile(path, outPath, false)
	err := os.Remove(expected)
	if err != nil {
		t.Errorf("handleFile(%s) = %v; want %s", path, err, expected)
	}
}

func TestHandleFileWrite(t *testing.T) {
	tempFile, err := os.Create("testdata/temp.srt")
	utils.Check(err)
	defer tempFile.Close()

	file, err := os.Open("testdata/test.srt")
	utils.Check(err)
	defer file.Close()

	io.Copy(file, tempFile)

	handleFile("testdata/test.srt", "", true)

	expected := condenser.CondenseSrt(condenser.ParseSrt(tempFile))
	actual := condenser.ParseSrt(file)

	if reflect.DeepEqual(expected, actual) {
		t.Errorf("handleFile(%v) = %v; want %v", file, actual, expected)
	}

	io.Copy(tempFile, file)
	os.Remove("testdata/temp.srt")
}
func TestHandleFileWriteDest(t *testing.T) {

}

func TestHandleDirectory(t *testing.T) {

}

func TestHandleDirectoryWrite(t *testing.T) {

}

func TestHandleDirectoryDest(t *testing.T) {

}

func TestHandleDirectoryWriteDest(t *testing.T) {

}
