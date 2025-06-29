package main

import (
	"Site/assets/blog"
	"fmt"
	"github.com/a-h/templ"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	comp := About()
	templ.Handler(comp).ServeHTTP(w, r)
}

func HandlerIndexRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/about/", http.StatusFound)

}

func HandlerBlog(w http.ResponseWriter, r *http.Request) {
	comp := Blog()
	templ.Handler(comp).ServeHTTP(w, r)
}

func HandlerBlogPost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	post := getBlogPost(slug)
	post.Content = getHTMLPostContent(post.HTMLContent)
	comp := BlogPost(post)
	templ.Handler(comp).ServeHTTP(w, r)
}

func getBlogPost(slug string) blog.Post {
	for _, post := range blog.Posts {
		if post.Id == slug {
			return post
		}
	}
	return blog.Post{}
}

func getHTMLPostContent(path string) string {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return "could not load post"
	}
	return string(f)
}

func HandlerPublications(w http.ResponseWriter, r *http.Request) {
	comp := Publications()
	templ.Handler(comp).ServeHTTP(w, r)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.Get("/", HandlerIndexRedirect)
	r.Get("/about/", HandlerAbout)
	r.Get("/blog/", HandlerBlog)
	r.Get("/blog/{slug}", HandlerBlogPost)
	r.Get("/publications/", HandlerPublications)

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", r)

}
