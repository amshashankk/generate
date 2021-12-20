package converter

import "testing"

func TestConvert(t *testing.T) {
	inputFiles := []string{"test.json"}
	outputDir := "output"
	err := Convert(inputFiles, outputDir)
	if err != nil {
		t.Error(err)
	}
}
