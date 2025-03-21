import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  test: {
    globals: true,
    environment: "jsdom",
    setupFiles: "./setupTests.js",
    coverage: {
      reporter: ['html', 'text', 'lcov'], // Ensure proper coverage reporters are specified
      reportsDirectory: './coverage',    // Set the directory where coverage will be saved
    },
  },
});
