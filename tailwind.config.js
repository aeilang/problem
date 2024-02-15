/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/temp/**/*{.tmpl, html}}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: false
  },
}


