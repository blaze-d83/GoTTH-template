// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func BaseTemplate() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"GoTTH Stack - A full-stack web development template using Go, TailwindCSS, Templ, and HTMX\"><meta name=\"keywords\" content=\"Go, TailwindCSS, Templ, HTMX, Fullstack\"><title>GoTTH Stack</title><link href=\"/static/dist/styles.css\" rel=\"stylesheet\"><script src=\"https://unpkg.com/htmx.org@1.9.2\"></script></head><body class=\"bg-gray-100 text-gray-900 antialiased\"><header class=\"p-6 bg-blue-600 text-white\"><div class=\"container mx-auto\"><h1 class=\"text-5xl font-extrabold\">Welcome to the GoTTH Stack!</h1><p class=\"text-lg mt-2\">Go + TailwindCSS + Templ + HTMX</p></div></header><main class=\"container mx-auto my-10 prose prose-lg\"><p>This template serves as the starting point for your GoTTH Stack application. Use Go for the backend, TailwindCSS for utility-first styling, Templ for templating, and HTMX for modern web interactions without JavaScript.</p><button class=\"px-4 py-2 mt-4 text-white bg-blue-600 rounded hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400\" hx-get=\"/some-endpoint\" hx-target=\"#response\">HTMX Example Button</button><div id=\"response\" class=\"mt-4\"></div></main><footer class=\"p-6 bg-blue-600 text-white\"><div class=\"container mx-auto text-center\"><p>&copy; 2024 GoTTH Stack. All rights reserved.</p></div></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
