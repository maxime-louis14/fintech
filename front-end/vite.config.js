import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": "/src"
    }
  },
  server: {
    watch: {
      usePolling: true, // Désactiver le sondage
      poll: 1000 // Définir un délai de sondage (en millisecondes)
    },
    host: true,
    strictPort: true,
    port: 3000
  }
});
