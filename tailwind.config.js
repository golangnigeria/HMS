/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.html",   // Go templates
    "./static/js/**/*.js",     // Alpine.js or custom scripts
  ],
  darkMode: "class",           // enables dark mode via `class`
  theme: {
    extend: {
      colors: {
        primary: "#11a4d4",
        "primary-dark": "#0e8db5",
        secondary: "#f0f9ff",
        "background-light": "#f6f8f8",
        "background-dark": "#101d22",
        "content-light": "#ffffff",
        "content-dark": "#1a2a31",
        "text-primary-light": "#0d181b",
        "text-primary-dark": "#e8f3f6",
        "text-secondary-light": "#4c869a",
        "text-secondary-dark": "#a0c7d4",
        "border-light": "#e7f0f3",
        "border-dark": "#2b4754",
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
}
