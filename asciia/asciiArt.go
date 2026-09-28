package asciiArt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AsciiArt(text string, style string) (string, error) {
	input := text

	bannerFile, err := getBannerFile(style)
	if err != nil {
		return "", err
	}

	fontMap, err := loadFont(bannerFile)
	if err != nil {
		return "", fmt.Errorf("load banner %q: %w", style, err)
	}

	finalArt := renderText(input, fontMap)

	return finalArt, nil
}

func getBannerFile(style string) (string, error) {
	switch style {
	case "shadow":
		return filepath.Join("style", "shadow.txt"), nil
	case "thinkertoy":
		return filepath.Join("style", "thinkertoy.txt"), nil
	case "standard":
		return filepath.Join("style", "standard.txt"), nil
	default:
		return "", fmt.Errorf("unsupported banner %q", style)
	}
}

func loadFont(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r", "")
	lines := strings.Split(content, "\n")

	font := make(map[rune][]string)
	ascii := 32

	for i := 0; i+8 < len(lines); i += 9 {
		font[rune(ascii)] = lines[i+1 : i+9]
		ascii++
	}

	return font, nil
}

func renderText(input string, font map[rune][]string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")

	var finalArt strings.Builder
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			finalArt.WriteByte('\n')
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range line {
				if art, ok := font[char]; ok && len(art) == 8 {
					finalArt.WriteString(art[row])
				} else {
					finalArt.WriteString("        ")
				}
			}
			finalArt.WriteByte('\n')
		}
	}
	return finalArt.String()
}
