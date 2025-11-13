package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golangnigeria/kinicart/internals/models"
	"github.com/golangnigeria/kinicart/internals/render"
)

// BlogList displays all blog posts
func (m *Repository) BlogList(w http.ResponseWriter, r *http.Request) {
	// Example blog posts (later can be loaded from DB)
	posts := []models.BlogPost{
		{
			Title:     "Healthy Nutrition for Busy People",
			Slug:      "healthy-nutrition-for-busy-people",
			Image:     "/images/nutrition.jpg",
			Body:      "Learn how to maintain a balanced diet even with a busy schedule.",
			CreatedAt: "Nov 5, 2025",
		},
		{
			Title:     "Managing Stress the Smart Way",
			Slug:      "managing-stress-the-smart-way",
			Image:     "/images/stress.jpg",
			Body:      "Practical steps for managing stress and maintaining mental balance.",
			CreatedAt: "Nov 8, 2025",
		},
		{
			Title:     "Why Regular Checkups Matter",
			Slug:      "why-regular-checkups-matter",
			Image:     "/images/checkup.jpg",
			Body:      "Discover why regular medical checkups are essential for long-term health.",
			CreatedAt: "Nov 10, 2025",
		},
	}

	td := &models.TemplateData{
		Data: map[string]any{
			"posts": posts,
		},
		Active: "blog",
	}

	render.RenderTemplate(w, r, "blog.page.html", td)
}

// BlogSingle displays a single blog post dynamically using slug from the URL
func (m *Repository) BlogSingle(w http.ResponseWriter, r *http.Request) {
	// Normally you'd get this from chi router like chi.URLParam(r, "slug")
	// But for simplicity, we’ll extract it manually here:
	slug := chi.URLParam(r, "slug")

	// Temporary static data (later replace with DB query)
	posts := []models.BlogPost{
		{
			Title:     "Healthy Nutrition for Busy People",
			Slug:      "healthy-nutrition-for-busy-people",
			Image:     "/images/nutrition.jpg",
			Body:      `Learn how to maintain a balanced diet even with a busy schedule. Maintaining a healthy diet is crucial for overall well-being, especially for those with hectic lifestyles.

Here are some tips to help you stay on track:

1. Plan Your Meals: Take some time each week to plan your meals. This will help you avoid last-minute unhealthy choices.

2. Prepare Snacks: Keep healthy snacks like fruits, nuts, and yogurt on hand to curb hunger between meals.

3. Stay Hydrated: Drink plenty of water throughout the day to stay hydrated and avoid mistaking thirst for hunger.

4. Choose Whole Foods: Opt for whole grains, lean proteins, and plenty of vegetables to ensure you're getting essential nutrients.

5. Limit Processed Foods: Try to minimize your intake of processed and fast foods, which are often high in unhealthy fats and sugars.

By following these tips, you can maintain a balanced diet even with a busy schedule.`,
			CreatedAt: "Nov 5, 2025",
		},
		{
	Title:     "Managing Stress the Smart Way",
	Slug:      "managing-stress-the-smart-way",
	Image:     "/images/stress.jpg",
	Body:      `Stress is a normal part of life, but when it becomes overwhelming, it can affect your mental, emotional, and physical health. Managing stress in a healthy way helps you stay focused, productive, and balanced.

1. **Identify Your Stress Triggers:**  
   Take note of the situations, people, or habits that cause stress. Awareness is the first step to managing it effectively.

2. **Practice Deep Breathing or Meditation:**  
   Techniques like slow breathing, mindfulness, or meditation can calm your nervous system and help you regain focus.

3. **Stay Active:**  
   Regular exercise, such as walking, yoga, or cycling, releases endorphins—your body’s natural stress relievers.

4. **Maintain Healthy Connections:**  
   Don’t isolate yourself. Talking with friends, family, or a counselor provides emotional support and fresh perspectives.

5. **Prioritize Sleep and Nutrition:**  
   Poor sleep or an unhealthy diet can worsen stress. Aim for 7–8 hours of sleep and eat balanced meals.

6. **Learn to Say No:**  
   Overcommitting can lead to burnout. Set boundaries and make time for rest and hobbies.

Remember, stress management isn’t about avoiding challenges but learning how to respond calmly and effectively. A balanced mind leads to a healthier body and a happier life.`,
	CreatedAt: "Nov 8, 2025",
},
		{
	Title:     "Why Regular Checkups Matter",
	Slug:      "why-regular-checkups-matter",
	Image:     "/images/checkup.jpg",
	Body:      `In today’s busy world, it’s easy to overlook regular medical checkups—especially when you feel healthy. However, preventive health care is one of the most effective ways to detect problems early and maintain long-term wellness.

1. **Early Detection Saves Lives:**  
   Many serious conditions, like diabetes, high blood pressure, and certain cancers, show few or no symptoms in their early stages. Routine screenings can identify these issues before they become severe.

2. **Build a Health History:**  
   Regular visits help your doctor understand your personal and family health trends, making it easier to detect changes or risks over time.

3. **Stay Updated on Vaccinations:**  
   Immunizations aren’t just for children. Adults also need vaccines to protect against preventable diseases like influenza, tetanus, and hepatitis.

4. **Promote Healthy Habits:**  
   Checkups provide a great opportunity to discuss diet, exercise, sleep, and stress management with your healthcare provider.

5. **Peace of Mind:**  
   Knowing that you’re taking proactive steps for your health reduces anxiety and helps you make informed lifestyle choices.

Your health is your most valuable asset. Don’t wait for symptoms to appear—make regular checkups a part of your wellness routine and invest in a healthier future.`,
	CreatedAt: "Nov 10, 2025",
},
	}

	// Find matching post
	var found *models.BlogPost
	for _, p := range posts {
		if p.Slug == slug {
			found = &p
			break
		}
	}

	if found == nil {
		http.NotFound(w, r)
		return
	}

	td := &models.TemplateData{
		Data: map[string]any{
			"post": found, // ✅ singular for single blog
		},
		Active: "blog",
	}

	render.RenderTemplate(w, r, "blog_single.page.html", td)
}
