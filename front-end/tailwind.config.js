/** @type {import('tailwindcss').Config} */
export default {
  darkMode: ["class", '[data-mode="dark"]'],
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        custom: ['Poppins-Regular', 'sans-serif'],
        // Ajoutez autant de polices personnalisées que nécessaire
      }
    }
  },
  plugins: []
};
