<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { toast } from 'svelte-sonner';

	import { Download } from '$lib/wailsjs/go/download/Downloader';
	import { config } from '$lib/wailsjs/go/models';

	import Label from './label.svelte';
	import DirectoryInput from './directory-input.svelte';
	import ResolutionSelect from './resolution-select.svelte';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import type { MouseEventHandler } from 'svelte/elements';

	let { defaults }: { defaults: config.Download } = $props();

	const form: config.Download = $state({});
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

	const handleOutsideClick: MouseEventHandler<HTMLAnchorElement> = (e) => {
		e.preventDefault();
		BrowserOpenURL(e.currentTarget.href);
	};
</script>

{#snippet urlInput()}
	{#snippet info()}
		Accepts both single video and playlist URLs.
	{/snippet}
	<Label for="url" {info}>URL</Label>
	<Input id="url" type="url" placeholder="video url" bind:value={url} required disabled={loading} />
{/snippet}

{#snippet outputTemplateInput()}
	{#snippet info()}
		<div>
			For details and examples of acceptable formats, check out <a
				href="https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#output-template"
				class="underline underline-offset-4"
				onclick={handleOutsideClick}>yt-dlp's docs</a
			>, or leave the field as is to use the default.
		</div>
	{/snippet}
	<Label for="outputTemplate" {info}>Output Template</Label>
	<Input
		id="outputTemplate"
		type="text"
		placeholder="%(title)s [%(id)s].%(ext)s"
		disabled={loading}
		bind:value={form.outputTemplate}
	/>
{/snippet}

{#snippet resolutionSelect()}
	{#snippet info()}
		There's no guarantee the selected resolution will be available. This functions as an upper limit
		for the available resolutions.
	{/snippet}
	<Label {info}>Resolution</Label>
	<ResolutionSelect disabled={loading} bind:value={form.targetResolution} />
{/snippet}

{#snippet directoryInput()}
	<Label for="outputDirectory">Output Directory</Label>
	<DirectoryInput
		disabled={loading}
		bind:value={form.outputDirectory}
		placeholder={defaults.outputDirectory}
	/>
{/snippet}

<form class="flex min-h-full flex-col items-center justify-center gap-4" onsubmit={downloadHandler}>
	<div class="flex w-full flex-wrap items-end justify-center gap-4">
		<div class="basis-1/4">
			{@render urlInput()}
		</div>
		<div class="basis-1/4">
			{@render outputTemplateInput()}
		</div>
		<div class="basis-1/4">
			{@render resolutionSelect()}
		</div>
		<div class="basis-1/2">
			{@render directoryInput()}
		</div>
	</div>
	<Button type="submit" disabled={loading}>Download</Button>
</form>
