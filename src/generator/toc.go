package generator

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	opts struct {
		Verbose  bool `short:"v" long:"verbose" description:"Show verbose debug information"`
		Generate *struct {
			Positional struct {
				Output string
			} `positional-args:"true" required:"true"`
		} `command:"generate"`

		Serve *struct {
			Port      int    `default:"1313"`
			Directory string `default:"."`
		} `command:"serve"`
	}

	copy_regex   = regexp.MustCompile("\\.(png|pdf|md|css|js|svg|woff2|ttf|woff|ttf|gif)$")
	asset_regex  = regexp.MustCompile(`src="([^"]+)"`)
	asset_regex2 = regexp.MustCompile(`!\[.*?\]\(([^\)]+)\)`)
)

func getHeading(part string) string {
	for _, line := range strings.Split(part, "\n") {
		if strings.HasPrefix(line, "#") {
			return strings.Trim(line, "# ")
		}
	}
	return ""
}

func GenerateSite(
	output_directory string,
	verbose bool, filterRegex string) error {

	now := time.Now()

	presentation, err := ParsePresentation()
	if err != nil {
		return err
	}

	if filterRegex != "" {
		err = presentation.Filter(filterRegex)
		if err != nil {
			return err
		}
	}

	defer func() {
		fmt.Printf("Loading presentations with %v to directory %v in %v\n",
			Stats(presentation), output_directory, time.Now().Sub(now))
	}()

	// Prepare the skeleton
	output_manager := OutputManager{output_directory, verbose}
	output_manager.CopyDirectory("./dist/", "dist")
	output_manager.CopyDirectory("./plugin/", "plugin")
	output_manager.CopyDirectory("./plugin/highlight", "plugin/highlight")
	output_manager.CopyDirectory("./plugin/markdown", "plugin/markdown")
	output_manager.CopyDirectory("./plugin/notes", "plugin/notes")
	output_manager.CopyDirectory("./plugin/zoom", "plugin/zoom")
	output_manager.CopyDirectory("./themes/workshop/", "themes/workshop")
	output_manager.CopyDirectory("./dist/theme", "dist/theme")
	output_manager.CopyDirectory("./resources", "resources")
	output_manager.CopyDirectory("./webfonts/", "webfonts")
	output_manager.CopyDirectory("./css", "css")
	output_manager.CopyDirectory("./js", "js")
	output_manager.CopyFile("./src/generator/README.md", "README.md")
	output_manager.CopyFile("./CNAME", "CNAME")
	output_manager.CopyFile("./.nojekyll", ".nojekyll")

	if verbose {
		Dump(presentation)
	}

	err = output_manager.WriteFile("index.html", buildPresentationTOC(presentation))
	if err != nil {
		return err
	}

	// Copy module directories into the output
	for _, module := range presentation.Modules {
		// The artifact is used to create the notebook for the module.
		output_manager.WriteFile(
			filepath.Join(module.Path, "artifact.yaml"),
			module.buildArtifact())

		// Create a HTML for the whole module
		output_manager.WriteFile(
			filepath.Join(module.Path, "index.html"),
			buildIndexHtml(module, presentation))

		// Copy all the files over
		output_manager.CopyDirectory(module.Path, module.Path)

		// Check all the topics and merge them with this module.
		for _, topic := range module.Topics {
			module_md_path := filepath.Join(module.Path, filepath.Base(topic.Path))

			// If a topic has an absolute path then it is an external
			// reference, copy the md file to the module.
			if filepath.IsAbs(topic.Path) {
				output_manager.CopyFile(topic.Path, module_md_path)
			}

			// Also create a html for each topic.
			if topic.Path != "index.md" {
				output_manager.WriteFile(
					topic.Link,
					buildIndexHtml(&Module{
						Topics: []*Topic{topic},
						Head:   module.Head,
					}, presentation))
			}

			// Copy all the assets into the module directory.
			for _, slide := range topic.Slides {
				for _, asset := range slide.Assets {
					// If the assets is absolute we need to copy the
					// asset into the same directory structure in the
					// output file so it can be found.
					if filepath.IsAbs(asset) {
						output_manager.CopyFile(asset, asset)
						continue
					}

					// If the asset is relative we need to copy the
					// asset into the module directory.
					topic_directory := filepath.Dir(topic.AbsPath)
					if !filepath.IsAbs(asset) {
						output_manager.CopyFile(
							filepath.Join(topic_directory, asset),
							filepath.Join(module.Path, asset))
					}
				}
			}
		}
	}

	return nil
}
