package routes

import (
	"ecommerce/api/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func adminPanel(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		panic(err)
	}
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/admin.tmpl",
		"./web/templates/pages/admin/home.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	
	data := session.Values["user"]
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err)
	}
}

func adminProducts(w http.ResponseWriter, r *http.Request) {
	// handling actions
	params := mux.Vars(r)
	// delete action
	if params["action"] == "delete" {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Redirect(w, r, "/admin/products", http.StatusFound)
			return
		}
		if ok, err := d.DeleteProduct(uint(id)); err != nil || !ok{
			http.Redirect(w, r, "/admin/products", http.StatusFound)
			return
		}
	}

	//modify/add actions
	if params["action"] == "modify" {
		product := &models.Product{}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Redirect(w, r, "/admin/products", http.StatusFound)
			return
		}


		if id > 0 {
			product, err = d.GetProduct(id)
			if err != nil {
				http.Redirect(w, r, "/admin/products", http.StatusFound)
				return
			}
		} 

		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Redirect(w, r, "/admin/products", http.StatusFound)
				return
			}
			fmt.Println(r.PostForm)
			prod := models.Product{}
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				fmt.Printf("err: %v\n", err)
				http.Redirect(w, r, "/admin/products", http.StatusFound)
				return
			}
			price, err := strconv.ParseFloat(r.FormValue("price"), 32)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				http.Redirect(w, r, "/admin/products", http.StatusFound)
				return
			}
			prod.ID = uint(id)
			prod.Name = r.FormValue("name")
			prod.Description = r.FormValue("description")
			prod.Price = float32(price)
			if ok, err := d.ModifyProduct(prod); err != nil || !ok {
				fmt.Printf("err: %v\n", err)
				http.Redirect(w, r, "/admin/products", http.StatusFound)
				return
			}
			http.Redirect(w, r, "/admin/products", http.StatusFound)
			return
		}


		files := []string{
			"./web/templates/base.tmpl",
			"./web/templates/partials/nav.tmpl",
			"./web/templates/pages/admin.tmpl",
			"./web/templates/pages/admin/modifyProduct.tmpl",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}

		err = t.ExecuteTemplate(w, "base", product)
		if err != nil {
			log.Println(err)
		}

		return
	}


	// normal flow
	products, err := d.GetProducts()
	if err != nil {
		panic(err)
	}
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/admin.tmpl",
		"./web/templates/pages/admin/products.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "base", products)
	if err != nil {
		log.Println(err)
	}
}

