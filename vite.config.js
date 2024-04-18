import path from "node:path";

export default {
    root: path.join(__dirname, "./resources/"),
    build: {
        manifest: "manifest.json",
        outDir: path.join(__dirname, "./public"),
        rollupOptions: {
            input: "resources/js/main.js",
        },
    },
};
