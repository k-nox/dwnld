<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { toast } from 'svelte-sonner';
	import { config } from '$lib/wailsjs/go/models';

	import Resolution from './fields/resolution.svelte';
	import Url from './fields/url.svelte';
	import Directory from './fields/directory.svelte';
	import OutputTemplate from './fields/outputTemplate.svelte';
	import { Download } from '$lib/wailsjs/go/app/App';

	import Switch from './switch.svelte';
	let { defaults }: { defaults: config.Download } = $props();

	const form: config.Download = $state({
		embedSubtitles: defaults.embedSubtitles,
		writeInfoJSON: defaults.writeInfoJSON
	});
	let url = $state('');

	let loading = $state(false);

	const downloadHandler = async (e: SubmitEvent) => {
		e.preventDefault();
		loading = true;
		toast.promise(Download(url, form), {
			loading: 'Downloading...',
			success: () => {
				url = '';
				return 'Video successfully downloaded.';
			},
			error: (e: unknown) => {
				let msg = '';
				if (typeof e === 'string') {
					msg = e;
				} else if (e instanceof Error) {
					msg = e.message;
				}
				return msg;
			},
			finally: () => {
				loading = false;
			}
		});
	};
</script>

<form class="flex min-h-full flex-col items-center justify-center gap-4" onsubmit={downloadHandler}>
	<div class="flex w-full flex-wrap items-end justify-center gap-4">
		<div class="basis-1/4">
			<Url bind:value={url} disabled={loading} label="url" />
		</div>
		<div class="basis-1/4">
			<OutputTemplate
				bind:value={form.outputTemplate}
				disabled={loading}
				defaultValue={defaults.outputTemplate}
				label="outputTemplate"
			/>
		</div>
		<div class="basis-1/4">
			<Resolution
				disabled={loading}
				bind:value={form.targetResolution}
				defaultValue={defaults.targetResolution}
			/>
		</div>
		<div class="basis-1/2">
			<Directory
				bind:value={form.outputDirectory}
				disabled={loading}
				defaultValue={defaults.outputDirectory}
				label="outputDirectory"
			/>
		</div>
		<div class="flex basis-1/2 flex-col items-start gap-2">
			<Switch
				id="write-info-json"
				bind:checked={form.writeInfoJSON}
				disabled={loading}
				label="Write info.json file"
			>
				{#snippet info()}
					Whether or not the info.json file should be downloaded alongside the video.
				{/snippet}
			</Switch>
			<Switch
				id="embed-subtitles"
				bind:checked={form.embedSubtitles}
				disabled={loading}
				label="Embed Subtitles"
			>
				{#snippet info()}
					Currently only English subtitles can be selected, and auto-generated subtitles will be
					ignored.
				{/snippet}
			</Switch>
		</div>
	</div>
	<Button type="submit" disabled={loading}>Download</Button>
</form>
