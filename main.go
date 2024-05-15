package main

import (
	"net/http"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func homePage() gomponents.Node {
	return components.HTML5(components.HTML5Props{
		Title: "Mi página",
		Head: []gomponents.Node{
			html.Script(
				html.Src("//unpkg.com/alpinejs"),
				html.Defer(),
			),
			html.Script(
				html.Src("https://unpkg.com/htmx.org@1.9.12"),
			),
			html.Link(
				html.Rel("stylesheet"),
				html.Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"),
			),
		},
		Body: []gomponents.Node{
			html.P(
				gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Enim facilisis gravida neque convallis a. Ante metus dictum at tempor commodo ullamcorper a lacus."),
			),
			html.Button(
				gomponents.Attr("hx-get", "/more-info"),
				gomponents.Attr("hx-on", "click"),
				gomponents.Attr("hx-target", "#more-info-container"),
				gomponents.Attr("hx-swap", "innerHTML"),
				gomponents.Text("Más información"),
			),
			html.Section(
				html.ID("more-info-container"),
			),
		},
	})
}

func moreInfoPage() gomponents.Node {
	return html.Div(
		gomponents.Attr("x-data", `{show: false}`),
		gomponents.Text(
			"Más información",
		),
		html.Button(
			gomponents.Attr("x-on:click", "show = !show"),
			gomponents.Text("Mostrar aún más información"),
		),
		html.P(
			gomponents.Attr("x-show", "show == true"),
			gomponents.Text("Más información aquí"),
		),
	)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		homePage().Render(w)
	})

	mux.HandleFunc("GET /more-info", func(w http.ResponseWriter, r *http.Request) {
		moreInfoPage().Render(w)
	})

	http.ListenAndServe(":8080", mux)
}
