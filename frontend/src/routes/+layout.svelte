<script lang="ts">
	import Icon from '$lib/components/Icon.svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import ThemeChanger from '$lib/components/ThemeChanger.svelte';
	import '@unocss/reset/tailwind.css';
	import { Candy, User } from 'lucide-svelte';
	import 'uno.css';
</script>

<div
	class="min-h-screen bg-white dark:bg-primary-900 text-black dark:text-white flex-(~ col) items-center"
>
	<header
		class="bg-primary-500 dark:bg-primary-700 rounded-(b-10) min-h-12 min-w-sm max-w-2xl w-full mx-4 p-4 flex items-center justify-between border-(b x current)"
	>
		<div class="centered space-x-3">
			<div class="bg-red rounded-full p-1 shadow-inner shadow-white/80">
				<Candy class="w-9 h-9 color-white fill-black" />
			</div>
			<h1 class="text-4xl w-fit h-fit font-normal">Gomeet</h1>
		</div>
		<div class="w-fit centered space-x-3">
			<ThemeChanger />
			<div class="hover:bg-gray-5/50 border-(~ current) w-fit rounded-full p-2">
				<User class="w-7 h-7" />
			</div>
		</div>
	</header>
	<div class="centered w-90% min-w-max h-[calc(100vh-80px)] rounded-3 p-2 border-(~ black)">
		<div class="w-40% max-w-xs h-full">
			<Sidebar />
		</div>
		<div class="w-full h-full">
			<slot />
		</div>
	</div>
</div>

<style>
	:global(body) {
		font-family: 'IBM Plex Sans', sans-serif;
	}
	:global(.centered) {
		--at-apply: flex items-center justify-center;
	}
</style>
