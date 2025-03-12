<script lang="ts">
	import ThemeToggle from '$lib/components/theme-toggle.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	let { children } = $props();
</script>

<ModeWatcher />
<div class="grid h-screen grid-cols-12 grid-rows-12">
	<header class="col-span-12 row-span-1 m-4 flex flex-row items-center justify-end">
		<ThemeToggle />
	</header>
	<main class="col-span-12 row-span-11">
		{@render children()}
	</main>
</div>
