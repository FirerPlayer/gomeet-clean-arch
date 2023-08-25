import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import UnoCSS from 'unocss/vite';
// import { presetAttributify, presetIcons, presetUno } from 'unocss';
import extractorSvelte from '@unocss/extractor-svelte';

export default defineConfig({
	plugins: [
		UnoCSS({
			extractors: [extractorSvelte()]
			// presets: [presetUno(), presetAttributify(), presetIcons()]
		}),
		sveltekit()
	],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});