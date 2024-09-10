package main

import (
	"fmt"
	"net/http"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "pages/index.html")
	case "/contact-us":
		http.ServeFile(w, r, "pages/contact-us.html")
	case "/faqs":
		http.ServeFile(w, r, "pages/faqs.html")
	case "/about":
		http.ServeFile(w, r, "pages/about.html")
	case "/privacy-policy-2":
		http.ServeFile(w, r, "pages/privacy-policy.html")
	case "/rummy-slots-max":
		http.ServeFile(w, r, "pages/rummy-slots-max.html")
	case "/teen-patti-epic":
		http.ServeFile(w, r, "pages/teen-patti-epic.html")
	case "/teen-patti-gold":
		http.ServeFile(w, r, "pages/teen-patti-gold.html")
	case "/terms-conditions":
		http.ServeFile(w, r, "pages/terms-conditions.html")
	default:
		http.NotFound(w, r)
	}
}

func main() {
	// Serving static files like CSS, JS, etc.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Serve HTML pages
	http.HandleFunc("/", servePage)
	http.HandleFunc("/contact-us", servePage)
	http.HandleFunc("/faqs", servePage)
	http.HandleFunc("/about", servePage)
	http.HandleFunc("/privacy-policy-2", servePage)
	http.HandleFunc("/rummy-slots-max", servePage)
	http.HandleFunc("/teen-patti-epic", servePage)
	http.HandleFunc("/teen-patti-gold", servePage)
	http.HandleFunc("/terms-conditions", servePage)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
