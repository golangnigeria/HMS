/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.html",
    "./static/js/**/*.js",
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        // 🔥 Main brand colors
        primary: "#f97316", // orange-500
        "primary-dark": "#c2410c",
        secondary: "#ffedd5",

        // 🌑 Surfaces
        "bg-dark": "#0d0d0d",
        "bg-light": "#ffffff",
        "card-dark": "#1a1a1a",
        "card-light": "#f9fafb",

        // ✨ Text
        "text-light": "#0d181b",
        "text-dark": "#f5f5f5",
        "text-muted-light": "#4c869a",
        "text-muted-dark": "#a0c7d4",

        // 🧱 Borders
        "border-light": "#e7f0f3",
        "border-dark": "#2b4754",

        // 🌈 Accent gradients
        "accent-1": "#fb923c",
        "accent-2": "#ea580c",
        "accent-3": "#c2410c",
      },
      backgroundImage: {
        "gradient-orange": "linear-gradient(to right, #f97316, #ea580c, #c2410c)",
        "gradient-radial": "radial-gradient(circle at top right, #f97316, #c2410c 70%)",
      },
      fontFamily: {
        display: ["Inter", "sans-serif"],
      },
      borderRadius: {
        DEFAULT: "0.5rem",
        lg: "0.75rem",
        xl: "1rem",
        "2xl": "1.5rem",
        full: "9999px",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
