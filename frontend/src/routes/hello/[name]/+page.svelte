<script lang="ts">
	import { page } from '$app/stores';

	let name: string = $page.params.name;
</script>

<h1 text-center text-amber>
	Hello {name} !!!!!!
</h1>
