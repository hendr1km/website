package thesis

type Chart struct {
	Date           string
	Header         string
	PreviewImage   string
	PreviewContent string
	HTMLContent    string
	Content        string
	Id             string
}

var Charts = []Chart{
	{
		Date:           "AUG 26 2024",
		Header:         "Test Header",
		PreviewImage:   "assets/thesis/test.svg",
		PreviewContent: "Test Preview",
		HTMLContent:    "assets/thesis/test.html",

		Id: "test",
	},
}
