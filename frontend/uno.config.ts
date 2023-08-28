// uno.config.ts
import { presetUno, transformerDirectives } from 'unocss';
import { defineConfig, presetWebFonts, transformerVariantGroup } from 'unocss';
import { presetScrollbar } from 'unocss-preset-scrollbar';
import extractorSvelte from '@unocss/extractor-svelte';

export default defineConfig({
	presets: [
		presetUno(),
		presetWebFonts({
			provider: 'bunny',
			fonts: {
				sans: 'IBM Plex Sans',
				mono: 'IBM Plex Mono'
			}
		}),
		presetScrollbar()
	],
	theme: {
		colors: {
			primary: {
				DEFAULT: '#06b6d4',
				100: '#cdf0f6',
				200: '#9be2ee',
				300: '#6ad3e5',
				400: '#38c5dd',
				500: '#06b6d4',
				600: '#0592aa',
				700: '#046d7f',
				800: '#024955',
				900: '#01242a'
			}
		}
	},
	transformers: [transformerVariantGroup(), transformerDirectives()],
	extractors: [extractorSvelte()]
});
