/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./www/pages/**/*.{html,md}"],
	theme: {
		extend: {
			spacing: {
				"page-top": "var(--page-top)",
				"page-gutter": "var(--page-gutter)",
				"header-height": "var(--header-height)",
			},
			maxWidth: {
				content: "var(--content-width)",
			},
			minHeight: {
				content: "var(--content-height)",
			},
			colors: {
				base: "rgb(var(--color-base) / <alpha-value>)",
				surface: "rgb(var(--color-surface) / <alpha-value>)",
				overlay: "rgb(var(--color-overlay) / <alpha-value>)",
				subtle: "rgb(var(--color-subtle) / <alpha-value>)",
				text: "rgb(var(--color-text) / <alpha-value>)",
				primary: "rgb(var(--color-primary) / <alpha-value>)",
				secondary: "rgb(var(--color-secondary) / <alpha-value>)",
			},
			borderColor: {
				DEFAULT: "rgb(var(--color-overlay))",
			},
			ringColor: {
				DEFAULT: "rgb(var(--color-primary) / 0.2)",
			},
		},
	},
	plugins: [],
};
