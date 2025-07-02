package main

import (
	"Site/assets/blog"
	"context"
	"fmt"
	"github.com/a-h/templ"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	comp := About()
	templ.Handler(comp).ServeHTTP(w, r)
}

func RenderAbout(f io.Writer) {
	err := About().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func HandlerProjects(w http.ResponseWriter, r *http.Request) {
	comp := Projects()
	templ.Handler(comp).ServeHTTP(w, r)
}

func RenderProjects(f io.Writer) {
	err := Projects().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func HandlerIndexRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/about/", http.StatusFound)

}

func HandlerBlog(w http.ResponseWriter, r *http.Request) {
	comp := Blog()
	templ.Handler(comp).ServeHTTP(w, r)
}

func RenderBlog(f io.Writer) {
	err := Blog().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func HandlerBlogPost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	post := getBlogPost(slug)
	post.Content = getHTMLPostContent(post.HTMLContent)
	comp := BlogPost(post)
	templ.Handler(comp).ServeHTTP(w, r)
}

func RenderBlogPost(post blog.Post, f io.Writer) {
	post.Content = getHTMLPostContent(post.HTMLContent)
	err := BlogPost(post).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
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

func RenderPublications(f io.Writer) {
	err := Publications().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func main() {
	GenerateStaticWebsite()
	CopyAssets()
	GenerateStaticBlogPosts()
}

func GenerateStaticWebsite() {

	pages := []struct {
		Path       string
		RenderFunc func(io.Writer)
	}{
		{"/about/", RenderAbout},
		{"/", RenderAbout},
		{"/publications/", RenderPublications},
		{"/blog/", RenderBlog},
		{"/projects/", RenderProjects},
	}

	for _, page := range pages {

		outputDir := filepath.Join("dist", page.Path)
		outputFile := filepath.Join(outputDir, "index.html")

		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		defer f.Close()

		page.RenderFunc(f)

	}
}

func GenerateStaticBlogPosts() {
	for _, post := range blog.Posts {
		outputDir := filepath.Join("dist/blog", post.Id)
		outputFile := filepath.Join(outputDir, "index.html")

		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		defer f.Close()

		RenderBlogPost(post, f)
	}
}

func CopyAssets() {
	cmd := exec.Command("cp", "-r", "assets", "dist/")
	if err := cmd.Run(); err != nil {
		log.Fatalf("copy failed: %v", err)
	}
}

//		{"/", HandlerIndexRedirect},
//		{"/projects/", HandlerProjects},
//		{"/blog/", HandlerBlog},
//		{"/blog/{slug}", HandlerBlogPost},
//		{"/publications/", HandlerPublications},

func ServeWebsite() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.Get("/", HandlerIndexRedirect)
	r.Get("/about/", HandlerAbout)
	r.Get("/projects/", HandlerProjects)
	r.Get("/blog/", HandlerBlog)
	r.Get("/blog/{slug}", HandlerBlogPost)
	r.Get("/publications/", HandlerPublications)

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", r)

}
