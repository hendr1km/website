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
		Date:           "JUN 12 2024",
		Header:         "Correlation Networks in R",
		PreviewImage:   "/assets/blog/networks/thumbnail.png",
		PreviewContent: "Plotting correlations as networks can give you a good first impression of the interconnectivity of your variables. This is a quick tutorial on how to create correlation networks in R. ",
		HTMLContent:    "assets/blog/networks/content.html",
		Id:             "networks",
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
