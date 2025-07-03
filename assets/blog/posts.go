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
		Date:           "AUG 26 2024",
		Header:         "Rewriting R Functions in C",
		PreviewImage:   "assets/blog/rewriting/thumbnail.png",
		PreviewContent: "A guide to using the C API with R to improve performance by rewriting vector operations.",
		HTMLContent:    "assets/blog/rewriting/content.html",
		Id:             "rewriting",
	},
	{
		Date:           "JUN 15 2024",
		Header:         "Introduction to partial residual plots",
		PreviewImage:   "assets/blog/partial/thumbnail.png",
		PreviewContent: "Partial residual plots can give you valuble information about your model and relationships in the data like interactions and nonlinear trends. In this post we look at use cases and different ways PRP can be created in ggplot.",
		HTMLContent:    "assets/blog/partial/content.html",
		Id:             "partial",
	},
	{
		Date:           "JUN 12 2024",
		Header:         "Correlation Networks in R",
		PreviewImage:   "assets/blog/networks/thumbnail.png",
		PreviewContent: "Plotting correlations as networks can give you a good first impression of the interconnectivity of your variables. This is a quick tutorial on how to create correlation networks in R.",
		HTMLContent:    "assets/blog/networks/content.html",
		Id:             "networks",
	},
	{
		Date:           "JUN 04 2024",
		Header:         "Exploring frequency weights in logistic regression",
		PreviewImage:   "assets/blog/weights/thumbnail.png",
		PreviewContent: "Frequency weights can be used in logistic regression to address class imbalance caused by sampling bias. In this post, we will use simulations to explore how to effectively choose weights and visually understand the benefits of this approach.",
		HTMLContent:    "assets/blog/weights/content.html",
		Id:             "weights",
	},
}
