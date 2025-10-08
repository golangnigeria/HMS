package handlers

import (
	"net/http"

	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/driver"
	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
	"github.com/golangnigeria/kinicart/internals/repository"
	"github.com/golangnigeria/kinicart/internals/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB: dbrepo.NewPostgreRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{
		StringMap: map[string]string{
			"welcome": "Welcome to KiLab",
			"tagline": "Simplifying healthcare connections between patients, doctors, and laboratories.",
		},
		IntMap: map[string]int{
			"totalPatients": 1342,
			"totalDoctors":  245,
			"totalLabs":     82,
		},
		FloatMap: map[string]float32{
			"referralSuccessRate": 98.4,
		},
		Data: map[string]any{
			"testimonials": []map[string]string{
				{
					"name":    "Dr. Adaobi Okeke",
					"comment": "KiLab has made patient referrals effortless and transparent.",
				},
				{
					"name":    "Emeka, Patient",
					"comment": "I received my test results faster than ever before. Highly recommend!",
				},
				{
					"name":    "LabOne Diagnostics",
					"comment": "Our lab receives consistent, trusted referrals daily.",
				},
			},
			"roles": []map[string]string{
				{"name": "Patient", "description": "Access test results and manage referrals easily."},
				{"name": "Doctor", "description": "Refer patients and track lab results seamlessly."},
				{"name": "Laboratory", "description": "Receive referrals and expand your partnerships."},
			},
		},
	}

	render.RenderTemplate(w, r, "home.page.html", td)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "about.page.html", td)
}
