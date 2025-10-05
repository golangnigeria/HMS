package handlers

import (
	"net/http"

	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{
		StringMap: map[string]string{
			"welcome": "This is coming from the backend",
			"tagline": "Building an ecosystem that connects farmers, investors, diaspora, and food markets.",
		},
		IntMap: map[string]int{
			"activeUsers": 124,
			"totalFarms":  58,
		},
		FloatMap: map[string]float32{
			"roi":               12.5,
			"foodCreditsIssued": 2500000, // ₦2.5M worth of credits
		},
		Data: map[string]any{
			"features": []string{
				"Invest in local farms",
				"Track growth in real time",
				"Earn profits securely",
			},
			"featuredFarms": []map[string]string{
				{
					"name":     "Cassava Farm Project",
					"location": "Enugu, Nigeria",
					"roi":      "15",
					"image":    "/static/images/farms/cassava.jpg",
				},
				{
					"name":     "Rice Farm Expansion",
					"location": "Kano, Nigeria",
					"roi":      "18",
					"image":    "/static/images/farms/rice.jpg",
				},
				{
					"name":     "Maize Irrigation Scheme",
					"location": "Oyo, Nigeria",
					"roi":      "20",
					"image":    "/static/images/farms/maize.jpg",
				},
			},
			"testimonials": []map[string]string{
				{
					"name":    "Ada, UK Diaspora",
					"comment": "MyNneFarm made it easy for me to support my family from abroad while earning ROI.",
				},
				{
					"name":    "Chinedu, Farmer",
					"comment": "As a farmer, I no longer worry about bank loans. Investors trust and support me directly.",
				},
			},
		},
	}

	render.RenderTemplate(w, r, "home.page.html", td)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "about.page.html", td)
}
