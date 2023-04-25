import { defineConfig } from "astro/config";

import node from "@astrojs/node";

// https://astro.build/config
export default defineConfig({
  integrations: [],
  site: `data-dojo-api-production-9c3d.up.railway.app`,
  output: "server",
  adapter: node({
    mode: "standalone"
  })
});
