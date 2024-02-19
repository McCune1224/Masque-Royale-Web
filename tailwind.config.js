/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    // "./views/**/*.templ",
    // "./views/**/*.go",
    // "./views/**/*.html",
    // "./views/components/**/*.templ",
    // "./views/components/**/*.go",
    // "./views/components/**/*.html",
    "./views/**/*",
  ],
  theme: {
    extend: {
      primary: "#FF6363",
      secondary: {
        100: "#E2E2D5",
        200: "#888883",
      },
    },
  },
  plugins: [],
};
