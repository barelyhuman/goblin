const { fontFamily } = require('tailwindcss/defaultTheme')

module.exports = {
	content: ['./src/**/*.svelte', './src/**/*.html'],
	theme: {
		extend: {
			maxWidth: {
				content: '48rem',
			},
			fontSize: {
				sm: '13px',
				md: '15px',
			},
			fontFamily: {
				sans: ['Inter', ...fontFamily.sans],
			},
			colors: {
				base: 'var(--base)',
				surface: 'var(--surface)',
				overlay: 'var(--overlay)',
				muted: 'var(--muted)',
				subtle: 'var(--subtle)',
				text: 'var(--text)',
				accent: 'var(--accent)',
				'on-accent': 'var(--on-accent)',
			},
			boxShadow: {
				DEFAULT: '0 10px 30px -20px rgba(87, 82, 121, 0.2)',
			},
			borderColor: (theme) => ({
				DEFAULT: theme('colors.overlay', 'currentColor'),
			}),
		},
	},
	plugins: [
		require('tailwind-scrollbar-hide')
	],
}
