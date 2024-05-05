/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    // "./views/**/*.templ",
    // "./views/**/*.go",
    // "./views/**/*.html",
    // "./views/components/**/*.templ",
    // "./views/components/**/*.go",
    // "./views/components/**/*.html",
    // "./views/**/*",
    // "./**/*.html",
    "./view/**/*.templ",
    "./view/**/*.go",
  ],
  theme: {
    extend: {
      fontFamily: {
        oxygen: ["Oxygen", 'sans-serif'],
      },
      colors: {
        space: '#1B1B3A',
        finn: '#693668',
        haze: '#A74482',
        rose: '#F84AA7',
        folly: '#F84AA7',
      },
    },
  },
  plugins: [],
};
