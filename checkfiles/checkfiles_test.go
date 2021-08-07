package checkfiles

import "testing"

func TestRun(t *testing.T) {
	pathName, size, date := Run("./TestFile")

	if pathName != "__TestFile" {

		t.Error(pathName, "is Wrong")
	}

	if size == "0" {
		t.Error("could not get file size")
	}

	if date == "0" {
		t.Error("cloud not get file date")
	}

}
