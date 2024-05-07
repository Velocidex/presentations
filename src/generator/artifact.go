package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

const (
	artifact_template = `
name: Notebooks.Presentation
type: NOTEBOOK

sources:
- notebook:
   {{- range $val := .Cells }}
   - type: {{ $val.Type }}
     template: |
{{ $val.Data }}
   {{- end }}
`

	RootURL = "https://github.com/Velocidex/presentations/blob/master/%v?raw=true"
)

type Cell struct {
	Type string
	Data string
}

type Parameters struct {
	Cells []Cell
}

func (self Module) buildArtifact() string {
	tmpl, err := template.New("").Parse(artifact_template)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}

	parameters := &Parameters{}

	for _, topic := range self.Topics {
		base_path := filepath.Dir(topic.Path)

		for _, s := range topic.Slides {
			parameters.Cells = append(parameters.Cells,
				convertSlide(s, base_path))
		}
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, parameters)
	return string(b.Bytes())
}

func indent(in string, indent string) string {
	lines := strings.Split(in, "\n")
	result := make([]string, 0, len(lines))
	for _, l := range lines {
		result = append(result, indent+l)
	}

	return strings.Join(result, "\n")
}

var (
	img_regex  = regexp.MustCompile(`src="([^"]+)"`)
	img_regex2 = regexp.MustCompile(`!\[(.*)\]\(([^)]+)\)`)
)

func convertPathToUrl(path string, base_path string) string {
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf(RootURL, path)
	}

	if strings.HasPrefix(path, "http") {
		return path
	}

	return fmt.Sprintf(RootURL, filepath.Join(base_path, path))
}

func convertSlide(slide *Slide, base_path string) Cell {
	data := strings.TrimSpace(slide.raw_data)

	data = img_regex.ReplaceAllStringFunc(data,
		func(hit string) string {
			match := img_regex.FindStringSubmatch(hit)
			if len(match) > 1 {
				return fmt.Sprintf("src=\"%v\"",
					convertPathToUrl(match[1], base_path))
			}
			return hit
		})

	data = img_regex2.ReplaceAllStringFunc(data,
		func(hit string) string {
			match := img_regex2.FindStringSubmatch(hit)
			if len(match) > 2 {
				return fmt.Sprintf("![%s](%v)",
					match[1], convertPathToUrl(match[2], base_path))
			}
			return hit
		})

	return Cell{
		Type: "markdown",
		Data: indent(data, "       "),
	}
}
