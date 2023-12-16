import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      backgroundImage: {},
    },
  },
  plugins: [],
  corePlugins: {
    preflight: false,
  },
};
export default config;
