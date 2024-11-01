// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AddProduct() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@1.8.4\"></script><title>Add Product</title></head><body class=\"bg-gray-100\"><div class=\"container mx-auto p-5\"><h1 class=\"text-2xl font-bold mb-4\">Add Product</h1><form hx-post=\"/product/add\" hx-target=\"#response-message\" hx-swap=\"innerHTML\"><input type=\"text\" name=\"name\" placeholder=\"Product Name\" class=\"p-2 border rounded mb-2 w-1/3\" required><br><input type=\"text\" name=\"category\" placeholder=\"Category\" class=\"p-2 border rounded mb-2 w-1/3\" required><br><input type=\"number\" name=\"quantity\" placeholder=\"Quantity\" class=\"p-2 border rounded mb-2 w-1/3\" required><br><input type=\"number\" name=\"unit_price\" placeholder=\"Unit Price\" class=\"p-2 border rounded mb-2 w-1/3\" required><br><button type=\"submit\" class=\"bg-blue-500 text-white p-2 rounded\">Add Product</button></form><div id=\"response-message\" class=\"mt-4\"></div></div><script>\n\t\tdocument.body.addEventListener(\"htmx:afterRequest\", function(event) {\n\t\t\tif (event.detail.xhr.status === 200) { \n\t\t\t\twindow.location.href = \"/viewproduct\";\n\t\t\t}\n\t\t});\n\n\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
