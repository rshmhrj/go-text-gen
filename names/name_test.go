package names

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	ng := NewGenerator()
	// Validate string representation
	output := ng.String()
	stringified := fmt.Sprintf("Nouns: {\n FullLength=1525, ShortLength=438, LongestWord=responsibility, ShortestWord=a \n Frequency: { 1:1, 2:4, 3:110, 4:323, 5:304, 6:245, 7:189, 8:135, 9:82, 10+:132 }\n}\nAdjectives: {\n FullLength=527, ShortLength=127, LongestWord=administrative, ShortestWord=ok \n Frequency: { 1:0, 2:2, 3:29, 4:96, 5:106, 6:88, 7:73, 8:54, 9:38, 10+:41 }\n}\nAdverbs: {\n FullLength=255, ShortLength=74, LongestWord=significantly, ShortestWord=as \n Frequency: { 1:0, 2:9, 3:16, 4:49, 5:44, 6:39, 7:28, 8:21, 9:23, 10+:26 }\n}")
	if len(output) != len(stringified) {
		t.Errorf("Incorrectly set string, got: %v, wanted: %v", len(output), len(stringified))
	}
	if strings.Compare(output, stringified) != 0 {
		t.Errorf("Incorrectly set string, \ngot:\n%v\nwanted:\n%v\n", output, stringified)
	}

	// Validate has dot delim by default
	if ng.delimiter != DotDelimiter {
		t.Errorf("Incorrectly set dot delimiter by default, got: %v, wanted %v", ng.delimiter, DotDelimiter)
	}

	var nameList []string

	nameList = append(nameList, ng.Generate())
	// Validate proper delim
	if !strings.Contains(nameList[0], DotDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[0], DotDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[0], DotDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[0], 2)
	}
	// Validate three words display
	count := len(wordCount(strings.Replace(nameList[0], DotDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	nameList = append(nameList, ng.GenerateShort())
	// Validate proper delim
	if !strings.Contains(nameList[1], DotDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[1], DotDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[1], DotDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[1], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[1], DotDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	ng.SetDelimiter(DashDelimiter)

	nameList = append(nameList, ng.Generate())
	// Validate proper delim
	if !strings.Contains(nameList[2], DashDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[2], DashDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[2], DashDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[2], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[2], DashDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	nameList = append(nameList, ng.GenerateShort())
	// Validate proper delim
	if !strings.Contains(nameList[3], DashDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[3], DashDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[3], DashDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[3], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[3], DashDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	ng.SetDelimiter(BarDelimiter)

	nameList = append(nameList, ng.Generate())
	// Validate proper delim
	if !strings.Contains(nameList[4], BarDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[4], BarDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[4], BarDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[4], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[4], BarDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	nameList = append(nameList, ng.GenerateShort())
	// Validate proper delim
	if !strings.Contains(nameList[5], BarDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[5], BarDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[5], BarDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[5], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[5], BarDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	ng.SetDelimiter(CommaDelimiter)

	nameList = append(nameList, ng.Generate())
	// Validate proper delim
	if !strings.Contains(nameList[6], CommaDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[6], CommaDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[6], CommaDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[6], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[6], CommaDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	nameList = append(nameList, ng.GenerateShort())
	// Validate proper delim
	if !strings.Contains(nameList[7], CommaDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[6], CommaDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[7], CommaDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[7], 2)
	}
	// Validate three words display
	count = len(wordCount(strings.Replace(nameList[7], CommaDelimiter, SpaceDelimiter, 2)))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	ng.SetDelimiter(SpaceDelimiter)

	nameList = append(nameList, ng.Generate())
	// Validate proper delim
	if !strings.Contains(nameList[8], SpaceDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[8], SpaceDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[8], SpaceDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[8], 2)
	}
	// Validate three words display
	count = len(wordCount(nameList[8]))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}

	nameList = append(nameList, ng.GenerateShort())
	// Validate proper delim
	if !strings.Contains(nameList[9], SpaceDelimiter) {
		t.Errorf("Incorrect delimiter, got: %v, wanted: %v", nameList[8], SpaceDelimiter)
	}
	// Validate no trailing delims
	if strings.Count(nameList[9], SpaceDelimiter) != 2 {
		t.Errorf("Incorrect number of delimiters, got: %v, wanted: %v", nameList[9], 2)
	}
	// Validate three words display
	count = len(wordCount(nameList[9]))
	if count != 3 {
		t.Errorf("Incorrect number of words, got: %v, wanted: %v", count, 3)
	}
}
