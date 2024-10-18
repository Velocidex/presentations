package generator

import (
	"bytes"
	"strings"
	"text/template"
)

const (
	index_html = `
<!doctype html>
<html lang="en">

    <head>
        <meta charset="utf-8">

        <title>Velociraptor Deployment</title>
        <link rel="stylesheet" href="../../css/fontawesome-all.min.css?1688344844">
        <link rel="stylesheet" href="{Base}/dist/reveal.css">
        <link rel="stylesheet" href="{Base}/dist/theme/serif.css" id="theme">
        <link rel="stylesheet" href="{Base}/css/velo.css">
        <link rel="stylesheet" href="{Base}/plugin/highlight/vs.css">
        {{ .Module.Head }}
    </head>
    <body>
        <div class="reveal">
            <div class="slides">
{{ range .Sections }}
<section data-markdown
  data-transition="fade"
  data-separator="^---+\n\n"
  data-separator-vertical="^>+\n\n">
<textarea data-template>
{{ . }}
</textarea>
</section>
{{ end }}
            </div>
        </div>
        <div class="footer"><a href="../../"><i class="fa fa-home"></i></a></div>
        <script src="{Base}/dist/reveal.js"></script>
        <script src="{Base}/plugin/markdown/markdown.js"></script>
        <script src="{Base}/plugin/highlight/highlight.js"></script>
        <script src="{Base}/plugin/notes/notes.js"></script>
        <script src="{Base}/plugin/zoom/zoom.js"></script>
        <script src="{Base}/js/jquery-3.3.1.min.js?1688344844"></script>
        <script src="{Base}/js/slides.js"></script>
        <script>
            Reveal.initialize({
                controls: true,
                progress: true,
                history: false,
                hash: true,
                center: false,
                slideNumber: true,

                plugins: [ RevealMarkdown, RevealHighlight, RevealNotes, RevealZoom ]
            }).then(initializeSlides);

        </script>

    </body>
</html>
`
	section_template = `
`
)

type Args struct {
	Sections []string
	Module   *Module
}

func buildIndexHtml(module *Module, presentation *Presentation) string {
	args := &Args{
		Module: module,
	}
	for _, topic := range module.Topics {
		data, err := readFile("./" + topic.AbsPath)
		if err != nil {
			continue
		}
		args.Sections = append(args.Sections, string(data))
	}

	tmpl, err := template.New("").Parse(index_html)
	if err != nil {
		panic("index_html invalid")
	}

	buf := bytes.Buffer{}

	err = tmpl.Execute(&buf, args)
	if err != nil {
		panic(err)
	}

	return adjustSectionText(string(buf.Bytes()))
}

var (
	expansions = map[string]string{
		"<!-- title optional -->":                          `<!-- .slide: class="title optional" data-background-color="antiquewhite" -->`,
		"<!-- content optional -->":                        `<!-- .slide: class="content optional" data-background-color="antiquewhite" -->`,
		"<!-- content small-font optional -->":             `<!-- .slide: class="content optional small-font" data-background-color="antiquewhite" -->`,
		"<!-- content -->":                                 `<!-- .slide: class="content" -->`,
		"<!-- content small-font -->":                      `<!-- .slide: class="content small-font" -->`,
		"<!-- full_screen_diagram small-font -->":          `<!-- .slide: class="full_screen_diagram small-font" -->`,
		"<!-- full_screen_diagram small-font optional -->": `<!-- .slide: class="full_screen_diagram small-font" data-background-color="antiquewhite" -->`,
		"<!-- hidden -->":                                  `<!-- .slide: class="content" data-visibility="hidden" -->`,
	}
)

func adjustSectionText(in string) string {
	for k, v := range expansions {
		in = strings.ReplaceAll(in, k, v)
	}

	in = strings.ReplaceAll(in, "/modules/", "../../modules/")
	return strings.ReplaceAll(in, "{Base}", "../..")
}
