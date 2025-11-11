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
			"doctorName": "Dr. Adaobi Okeke",
			"welcome":    "Your Health, Our Priority",
			"tagline":    "Personalized medical consultations, health insights, and wellness programs — all in one place.",
			"cta":        "Book your consultation today and take charge of your health.",
		},
		IntMap: map[string]int{
			"consultationsCompleted": 1200,
			"happyPatients":          980,
			"articlesPublished":      75,
		},
		FloatMap: map[string]float32{
			"patientSatisfaction": 97.5,
		},
		Data: map[string]any{
			"services": []map[string]string{
				{
					"title":       "Online Consultation",
					"description": "Schedule video sessions and receive professional medical advice from the comfort of your home.",
					"icon":        "💻",
				},
				{
					"title":       "In-Person Consultation",
					"description": "Visit our clinic for personalized checkups and wellness assessments.",
					"icon":        "🏥",
				},
				{
					"title":       "Follow-up & Monitoring",
					"description": "Stay on track with continuous follow-ups and progress tracking plans.",
					"icon":        "📋",
				},
			},
			"testimonials": []map[string]string{
				{
					"name":    "Chinedu M.",
					"comment": "Dr. Adaobi is not only professional but also genuinely caring. My lifestyle has completely improved!",
				},
				{
					"name":    "Ngozi E.",
					"comment": "Booking a consultation was so easy. The follow-up and guidance were excellent.",
				},
				{
					"name":    "Uche Health Blog Reader",
					"comment": "Her articles helped me understand nutrition and mental balance better.",
				},
			},
			"featuredArticles": []map[string]string{
				{
					"title": "Healthy Nutrition for Busy People",
					"slug":  "healthy-nutrition-for-busy-people",
					"image": "/images/nutrition.jpg",
				},
				{
					"title": "Managing Stress the Smart Way",
					"slug":  "managing-stress-the-smart-way",
					"image": "/images/stress.jpg",
				},
				{
					"title": "Why Regular Checkups Matter",
					"slug":  "why-regular-checkups-matter",
					"image": "/images/checkup.jpg",
				},
			},
		},
		Active: "home",
	}

	render.RenderTemplate(w, r, "home.page.html", td)
}


func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	td := &models.TemplateData{}

	render.RenderTemplate(w, r, "about.page.html", td)
}
