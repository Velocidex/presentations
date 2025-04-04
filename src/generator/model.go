package generator

import (
	"fmt"
	"regexp"
)

// A Presentation contains modules. Each module represents a
// directory.
type Presentation struct {
	Modules []*Module `yaml:"toc" json:"modules,omitempty"`
}

func (self *Presentation) Filter(regex string) error {
	re, err := regexp.Compile(regex)
	if err != nil {
		return err
	}

	modules := make([]*Module, 0, len(self.Modules))
	for _, m := range self.Modules {
		if re.MatchString(m.Path) {
			modules = append(modules, m)
		}
	}

	self.Modules = modules
	return nil
}

// A Module represents a collection of topics.
type Module struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"` // The name of the markdown file

	Topics []*Topic `yaml:"toc" json:"topics,omitempty"`

	Head string `yaml:"head" json:"head,omitempty"`
}

// A Topic is collection of slides in a single markdown file.
type Topic struct {
	Name    string `yaml:"name"`
	Path    string `yaml:"path"` // The name of the markdown file
	AbsPath string
	Link    string `yaml:"link"`

	Slides []*Slide `yaml:"-" json:"slides,omitempty"`
}

// A Slide may contain assets (like images)
type Slide struct {
	Title string
	Index int

	// Any assets in slides.
	Assets []string `yaml:"assets"`

	raw_data string
}

func Stats(presentation *Presentation) string {
	assets := 0
	slides := 0
	topics := 0
	modules := 0

	for _, module := range presentation.Modules {
		modules++
		for _, topic := range module.Topics {
			topics++
			for _, slide := range topic.Slides {
				slides++

				assets += len(slide.Assets)
			}
		}
	}

	return fmt.Sprintf("Stats: %v modules, %v topics, %v slides, %v assets",
		modules, topics, slides, assets)
}
