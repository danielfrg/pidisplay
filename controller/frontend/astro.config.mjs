// @ts-check
import { defineConfig } from "astro/config";

import tailwindcss from "@tailwindcss/vite";

import alpinejs from "@astrojs/alpinejs";

// https://astro.build/config
export default defineConfig({
  output: "static",
  outDir: "../internal/server/dist",

  vite: {
    plugins: [tailwindcss()],
  },

  integrations: [alpinejs()],
});
