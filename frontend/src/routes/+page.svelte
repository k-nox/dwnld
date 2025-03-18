<script lang="ts">
	import DownloadForm from '$lib/components/download/download-form.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	let { defaults } = data;
</script>

<div class="h-full">
	<DownloadForm {defaults} />
</div>
