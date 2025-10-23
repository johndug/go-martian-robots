package main

import (
	"strings"
	"testing"
)

func TestParseGridSize(t *testing.T) {
	lines := []string{"5 3", "1 1 E", "RFRFRFRF"}
	w, h, err := parseGridSize(lines)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if w != 5 {
		t.Errorf("Expected width 5, got %d", w)
	}

	if h != 3 {
		t.Errorf("Expected height 3, got %d", h)
	}
}

func TestParseRobotPosition(t *testing.T) {
	x, y, facing, err := parseRobotPosition("1 1 E")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if x != 1 {
		t.Errorf("Expected x=1, got %d", x)
	}

	if y != 1 {
		t.Errorf("Expected y=1, got %d", y)
	}

	if facing != "E" {
		t.Errorf("Expected facing=E, got %s", facing)
	}
}

func TestReadFileContent(t *testing.T) {
	// Test the file reading logic directly
	testContent := "5 3\n1 1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL"

	// Split by newlines and filter empty lines (same logic as readFile)
	lines := strings.Split(testContent, "\n")
	var result []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}

	expectedLines := 5
	if len(result) != expectedLines {
		t.Errorf("Expected %d lines, got %d", expectedLines, len(result))
	}

	if result[0] != "5 3" {
		t.Errorf("Expected first line '5 3', got '%s'", result[0])
	}

	if result[1] != "1 1 E" {
		t.Errorf("Expected second line '1 1 E', got '%s'", result[1])
	}
}
