package service

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	"strconv"
	// "github.com/a-h/templ"
	"gotth/view/pages"
	"gotth/view/layout"

	"gotth/model"
)

// AddProductHandler handles the creation of a new product.
func AddProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
        name := r.FormValue("name")
        category := r.FormValue("category")
		quantityStr := r.FormValue("quantity")
		unitPriceStr := r.FormValue("unit_price")
		
		// Convert quantity and unitPrice to their respective types
		quantity, errQuantity := strconv.Atoi(quantityStr)
		unitPrice, errUnitPrice := strconv.ParseFloat(unitPriceStr, 64)
		
		// Check if any required field is missing or invalid
		if name == "" || category == "" || errQuantity != nil || errUnitPrice != nil {
			http.Error(w, "Invalid request parameters.", http.StatusBadRequest)
			return
		}
		

		model.AddProduct(name, category, quantity, unitPrice)


        fmt.Fprintf(w, "Product %s added successfully!")
		// http.Redirect(w, r, "/viewproduct", http.StatusSeeOther)
    } else {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
    }

}

// GetAllProductsHandler retrieves all products.
func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := model.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByIDHandler retrieves a product by its ID.
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	product := model.GetProductByID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// UpdateProductHandler updates an existing product.
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
        id := r.FormValue("id")
        name := r.FormValue("name")
        category := r.FormValue("category")
		quantityStr := r.FormValue("quantity")
		unitPriceStr := r.FormValue("unit_price")
		
		// Convert quantity and unitPrice to their respective types
		quantity, errQuantity := strconv.Atoi(quantityStr)
		unitPrice, errUnitPrice := strconv.ParseFloat(unitPriceStr, 64)
		
		// Check if any required field is missing or invalid
		if id == "" || name == "" || category == "" || errQuantity != nil || errUnitPrice != nil {
			http.Error(w, "Invalid request parameters.", http.StatusBadRequest)
			return
		}
		

		model.UpdateProduct(id, name, category, quantity, unitPrice)

        fmt.Fprintf(w, "Product %s updated successfully!", id)
		// http.Redirect(w, r, "/viewproduct", http.StatusSeeOther)
    } else {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
    }
}

func ViewProductsHandler(w http.ResponseWriter, r *http.Request) {
    products := model.GetAllProducts() // Assume this returns a slice of products
    
    // Render rows as HTML directly for HTMX
    for _, product := range products {
        fmt.Fprintf(w, `
            <tr>
                <td class="border px-4 py-2">%d</td>
                <td class="border px-4 py-2">%s</td>
                <td class="border px-4 py-2">%s</td>
                <td class="border px-4 py-2">%d</td>
                <td class="border px-4 py-2">$%.2f</td>
                <td class="border px-4 py-2">
                    <a href="/update?id=%d" class="text-blue-500">Edit</a>
                </td>
            </tr>`, product.ID, product.Name, product.Category, product.Quantity, product.UnitPrice, product.ID)
    }
}

func RenderUpdate(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")

	product := model.GetProductByID(id)

	layout.Base(pages.UpdateProduct(id,product)).Render(r.Context(), w)

}