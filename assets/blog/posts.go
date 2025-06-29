package blog

type Post struct {
	Date           string
	Header         string
	PreviewImage   string
	PreviewContent string
	HTMLContent    string
	Content        string
	Id             string
}

var Posts = []Post{
	{
		Date:           "JAN 6 2024",
		Header:         "test",
		PreviewImage:   "assets/blog/content/test_post/typewriter.png",
		PreviewContent: "test preview",
		HTMLContent:    "assets/blog/content/test_post/test.html",
		Id:             "test-post",
	},
	{
		Date:           "JAN 6 2024",
		Header:         "test 2",
		PreviewImage:   "/assets/blog/pipeline/preview.png",
		PreviewContent: "test preview",
		HTMLContent:    "test content",
		Id:             "test-post2",
	},
}
