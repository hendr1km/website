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
		Date:           "Dezember 2025",
		Header:         "μ-Komponente (Baseline)",
		PreviewImage:   "assets/thesis/response_surface_mu.svg",
		PreviewContent: "3D Darstellung der μ-Komponente (Baseline) des Modells ohne Lag-Effekte.",
		HTMLContent:    "assets/thesis/response_surface_mu.html",
		Id:             "response_surface_mu",
	},
	{
		Date:           "Dezember 2025",
		Header:         "λ-Komponente (Lag-Effekt)",
		PreviewImage:   "assets/thesis/response_surface_lambda.svg",
		PreviewContent: "Response Surface der λ-Komponente, welche die zeitverzögerten Effekte (Lag) modelliert.",
		HTMLContent:    "assets/thesis/response_surface_lambda.html",
		Id:             "response_surface_lambda",
	},
	{
		Date:           "Dezember 2025",
		Header:         "σ-Komponente (Varianz)",
		PreviewImage:   "assets/thesis/response_surface_sigma.svg",
		PreviewContent: "3D Darstellung der modellierten Varianzstruktur (σ-Komponente) des Gesamtmodells.",
		HTMLContent:    "assets/thesis/response_surface_sigma.html",
		Id:             "response_surface_sigma",
	},
	{
		Date:           "Dezember 2025",
		Header:         "Gesamtmodell (ohne Lag)",
		PreviewImage:   "assets/thesis/response_surface_full_model.svg",
		PreviewContent: "3D Response Surface der Fixed Effects des Gesamtmodells mit den Komponenten μ (Baseline), λ (Lag-Effekt) und σ (Varianz) bei durchschnittlichem Affekt.",
		HTMLContent:    "assets/thesis/response_surface_full_model.html",
		Id:             "response_surface_full_model",
	},
	{
		Date:           "Dezember 2025",
		Header:         "Gesamtmodell (t-1: −1 SD)",
		PreviewImage:   "assets/thesis/response_surface_full_model_lag-1.svg",
		PreviewContent: "Fixed Effects des Gesamtmodells bei −1 SD des positiven Affekts zum Zeitpunkt t−1.",
		HTMLContent:    "assets/thesis/response_surface_full_model_lag-1.html",
		Id:             "response_surface_full_model_lag-1",
	},
	{
		Date:           "Dezember 2025",
		Header:         "Gesamtmodell (t-1: +1 SD)",
		PreviewImage:   "assets/thesis/response_surface_full_model_lag1.svg",
		PreviewContent: "Fixed Effects des Gesamtmodells bei +1 SD des positiven Affekts zum Zeitpunkt t−1.",
		HTMLContent:    "assets/thesis/response_surface_full_model_lag1.html",
		Id:             "response_surface_full_model_lag1",
	},
}
