/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {},
	},
	daysui: {
		theme: ['dark', 'light', 'cupcake'],
	},
	plugins: [require('daisyui')],
};
