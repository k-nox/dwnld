<script lang="ts">
	import ThemeToggle from '$lib/components/theme-toggle.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import Settings from 'lucide-svelte/icons/settings';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button';
	import { EventsEmit } from '$lib/wailsjs/runtime/runtime';
	let { children } = $props();
</script>

<ModeWatcher />
<Toaster richColors />
<div class="flex h-screen flex-col">
	<header class="flex flex-row items-center justify-end gap-6 p-6" style="--wails-draggable:drag">
		<div style="--wails-draggable:nodrag">
			<Button variant="outline" size="icon" onclick={() => EventsEmit('openSettings')}
				><Settings /></Button
			>
		</div>
		<div style="--wails-draggable:nodrag">
			<ThemeToggle />
		</div>
	</header>
	<main class="h-full">
		{@render children()}
	</main>
</div>
