// file: main.go
package main

import (
    "log"
    "net/http"

    "gotth/view/layout"
	"gotth/view/pages"
	"gotth/model"

    "github.com/a-h/templ"
	"gotth/service"
)

func main() {

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // http.Handle("/", templ.Handler(c))

  
	http.Handle("/viewproduct", templ.Handler(layout.Base(pages.Viewproduct(model.TestValue()))))
	http.Handle("/add", templ.Handler(layout.Base(pages.AddProduct())))
	// http.Handle("/update", templ.Handler(layout.Base(pages.UpdateProduct())))
	// http.Handle("/update", service.RenderUpdate)
	http.HandleFunc("/update", service.RenderUpdate)

	


	http.HandleFunc("/products", service.ViewProductsHandler)
	http.HandleFunc("/product", service.GetProductByIDHandler)
	http.HandleFunc("/product/add", service.AddProductHandler)
	http.HandleFunc("/product/update", service.UpdateProductHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
