<script lang="ts">
	import ThemeToggle from '$lib/components/theme-toggle.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	let { children } = $props();
</script>

<ModeWatcher />
<Toaster richColors />
<div class="flex h-screen flex-col">
	<header class="flex flex-row items-center justify-end p-6" style="--wails-draggable:drag">
		<div style="--wails-draggable:nodrag">
			<ThemeToggle />
		</div>
	</header>
	<main class="h-full">
		{@render children()}
	</main>
</div>
