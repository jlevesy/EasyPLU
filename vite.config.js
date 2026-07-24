import { dirname, resolve } from 'path'
import { fileURLToPath } from 'node:url'
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

const __dirname = dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  root: resolve(__dirname, 'assets'),
  build: {
    outDir: "static",
    assetsDir: "",
    manifest: "assets-manifest.json",
    rollupOptions: {
      input: [
        "src/js/layout.js",
        "src/css/style.css",
      ],
    },
    copyPublicDir: false,
  },
  plugins: [tailwindcss()],
})
