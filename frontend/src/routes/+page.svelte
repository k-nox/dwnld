<script lang="ts">
	import DownloadForm from '$lib/components/download/download-form.svelte';
	import { onMount } from 'svelte';
	import type { PageProps } from './$types';
	import { EventsOff, EventsOn } from '$lib/wailsjs/runtime/runtime';
	import Settings from '$lib/components/settings/settings.svelte';

	let { data }: PageProps = $props();
	let { defaults } = data;
	let settingsOpen = $state(false);

	onMount(() => {
		EventsOn('openSettings', () => {
			settingsOpen = true;
		});
		return () => {
			EventsOff('openSettings');
		};
	});
</script>

<Settings bind:isOpen={settingsOpen} />
<div class="h-full">
	<DownloadForm {defaults} />
</div>
