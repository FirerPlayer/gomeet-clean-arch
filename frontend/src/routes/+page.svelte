<script lang="ts">
	import { ArrowRight } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { sineIn, sineOut } from 'svelte/easing';
	import { fly } from 'svelte/transition';

	const handleClick = (ev: MouseEvent) => {
		dialogOpen = false;
	};
	let dialogOpen = false;
	onMount(() => {
		setTimeout(() => {
			dialogOpen = true;
		}, 1500);
	});
</script>

{#if dialogOpen}
	<div class="w-[100lvw] h-[100lvh] bg-black/40 absolute z-10 top-0 left-0 pointer-events-none" />
	<div
		in:fly={{ y: '50%', duration: 700, opacity: 0, easing: sineOut }}
		out:fly={{ y: '-50%', duration: 700, opacity: 0, easing: sineIn }}
		class="bg-primary-100 dark:bg-primary-900 absolute
		top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 centered
		w-90% flex-col gap-2 p-5 rounded-3 border-(~ black) shadow-xl z-20"
	>
		<h1 class="text-4xl font-bold text-cyan">Hello !!</h1>
		<h2 class="text-lg [&>span]:font-semibold">
			Welcome to Gomeet, a realtime chat app built with <span class="text-cyan">Golang</span> and
			<span class="text-orange-6">SvelteKit</span>. Take some tour trhough the app to get a feel of
			it. You now have access to the following features:
			<ul class="ml-3 [&>li]:font-semibold [&>li]:flex [&>li]:items-center [&>li]:gap-3">
				<li>
					<ArrowRight class="w-4 h-4" />
					Realtime chat
				</li>
				<li>
					<ArrowRight class="w-4 h-4" />
					File sharing
				</li>
				<li>
					<ArrowRight class="w-4 h-4" />
					Group chat
				</li>
			</ul>
		</h2>
		<form method="post">
			<button
				class="p-3 bg-cyan-3 rounded-3 border-black hover:bg-cyan-4 transition-all duration-300 mt-3 focus:outline focus:outline-orange"
				on:click={handleClick}
			>
				Awesome ! Let's go
			</button>
		</form>
	</div>
{/if}
Content do chat
