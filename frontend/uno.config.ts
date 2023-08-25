// uno.config.ts
import { presetUno } from 'unocss';
import { defineConfig, presetAttributify, presetIcons, presetWebFonts } from 'unocss';

export default defineConfig({
	presets: [
		presetUno(),
		presetAttributify(),
		presetIcons(),
		presetWebFonts({
			provider: 'bunny',
			fonts: {
				sans: 'IBM Plex Sans',
				mono: 'IBM Plex Mono'
			}
		})
	]
});
