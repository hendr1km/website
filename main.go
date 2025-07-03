package main

import (
	"Site/assets/blog"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RenderAbout(f io.Writer) {
	err := About().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func RenderProjects(f io.Writer) {
	err := Projects().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func RenderBlog(f io.Writer) {
	err := Blog().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func RenderBlogPost(post blog.Post, f io.Writer) {
	post.Content = getHTMLPostContent(post.HTMLContent)
	err := BlogPost(post).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func getHTMLPostContent(path string) string {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return "could not load post"
	}
	return string(f)
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

		outputDir := filepath.Join("docs", page.Path)
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
		outputDir := filepath.Join("docs/blog", post.Id)
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
	cmd := exec.Command("cp", "-r", "assets", "docs/")
	if err := cmd.Run(); err != nil {
		log.Fatalf("copy failed: %v", err)
	}
}
