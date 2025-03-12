<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Download } from '$lib/wailsjs/go/app/Downloader';
	import { LogDebug } from '$lib/wailsjs/runtime/runtime';
	let url = $state('');

	let status = $state('');

	const download = async () => {
		status = 'loading';
		try {
			await Download(url);
		} catch (e) {
			if (typeof e === 'string') {
				LogDebug(e);
			} else if (e instanceof Error) {
				LogDebug(e.message);
			}
		}
		status = '';
		url = '';
	};
</script>

<div class="h-full">
	{#if status === 'loading'}
		<div class="flex min-h-full flex-col items-center justify-center">loading...</div>
	{:else}
		<form class="flex min-h-full flex-col items-center justify-center gap-4">
			<Input class="w-1/3" type="text" placeholder="video url" bind:value={url} />
			<Button type="button" onclick={download}>Download</Button>
		</form>
	{/if}
</div>
