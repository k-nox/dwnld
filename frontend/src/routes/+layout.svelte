<script lang="ts">
	import ThemeToggle from '$lib/components/theme-toggle.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	let { children } = $props();
</script>

<ModeWatcher />
<Toaster richColors closeButton />
<div class="grid h-screen grid-cols-12 grid-rows-12">
	<header class="col-span-12 row-span-1 flex flex-row items-center justify-end p-6">
		<ThemeToggle />
	</header>
	<main class="col-span-12 row-span-11">
		{@render children()}
	</main>
</div>
