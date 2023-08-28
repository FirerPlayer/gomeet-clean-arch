<script lang="ts">
	import { elasticOut } from 'svelte/easing';
	import { Computer, Moon, Sun } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { writable, type Writable } from 'svelte/store';

	enum WebTheme {
		Light,
		Dark,
		System
	}

	let currTheme: Writable<WebTheme>;

	onMount(() => {
		if (localStorage.getItem('theme') !== null) {
			currTheme = writable(Number.parseInt(localStorage.getItem('theme') as string) as WebTheme);
		} else {
			currTheme = writable(WebTheme.System);
		}
		currTheme.subscribe((value) => {
			if (value === WebTheme.Dark) {
				document.documentElement.classList.add('dark');
				localStorage.setItem('theme', `${WebTheme.Dark}`);
			} else if (value === WebTheme.Light) {
				document.documentElement.classList.remove('dark');
				localStorage.setItem('theme', `${WebTheme.Light}`);
			} else {
				localStorage.setItem('theme', `${WebTheme.System}`);
				const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
				if (isDark) {
					document.documentElement.classList.add('dark');
				} else {
					document.documentElement.classList.remove('dark');
				}
			}
		});
	});

	const handleThemeChange = () => {
		if ($currTheme === WebTheme.System) {
			currTheme.set(WebTheme.Light);
		} else if ($currTheme === WebTheme.Light) {
			currTheme.set(WebTheme.Dark);
		} else {
			currTheme.set(WebTheme.System);
		}
	};

	const transverse = (node: Node, { duration = 800, delay = 0, easing = elasticOut } = {}) => {
		return {
			duration,
			delay,
			easing,
			css: (t: number) => `transform: rotate(${360 * t}deg);`
		};
	};
</script>

<button
	on:click={handleThemeChange}
	class="ring-1 ring-current rounded-full p-2 w-fit hover:bg-gray-5/50"
	class:ring-amber={$currTheme === WebTheme.Dark}
>
	{#key $currTheme}
		<div class="w-6 h-6" in:transverse>
			<svelte:component
				this={$currTheme === WebTheme.System ? Computer : $currTheme === WebTheme.Dark ? Sun : Moon}
				size="100%"
				class={$currTheme === WebTheme.Dark ? 'fill-white color-amber' : 'fill-white'}
			/>
		</div>
	{/key}
</button>
