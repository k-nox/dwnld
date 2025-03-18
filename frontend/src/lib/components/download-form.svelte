<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Download } from '$lib/wailsjs/go/download/Downloader';
	import { toast } from 'svelte-sonner';
	import DirectoryInput from './directory-input.svelte';
	import InfoPopover from './info-popover.svelte';
	import TemplateInfo from './template-info.svelte';
	import { config } from '$lib/wailsjs/go/models';
	import ResolutionSelect from './resolution-select.svelte';

	let { defaults }: { defaults: config.Download } = $props();

	const form: config.Download = $state({ ...defaults });
	let url = $state('');
	let dir = $state('');
	$inspect(dir);

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

{#snippet urlInput()}
	<Label for="url">URL<InfoPopover>Currently only single videos are supported.</InfoPopover></Label>
	<Input id="url" type="url" placeholder="video url" bind:value={url} required disabled={loading} />
{/snippet}

{#snippet outputTemplateInput()}
	<Label for="templ">Output Template<InfoPopover><TemplateInfo /></InfoPopover></Label>
	<Input
		id="templ"
		type="text"
		placeholder="%(title)s [%(id)s].%(ext)s"
		disabled={loading}
		bind:value={form.outputTemplate}
	/>
{/snippet}

{#snippet resolutionSelect()}
	<Label
		>Resolution<InfoPopover
			>There's no guarantee the selected resolution will be available. This functions as an upper
			limit for the available resolutions.</InfoPopover
		></Label
	>
	<ResolutionSelect disabled={loading} bind:value={form.targetResolution} />
{/snippet}

<form class="flex min-h-full flex-col items-center justify-center gap-4" onsubmit={downloadHandler}>
	<div class="flex w-full flex-wrap items-end justify-center gap-4">
		<div class="min-w-96">
			{@render urlInput()}
		</div>
		<div class="min-w-96">
			{@render outputTemplateInput()}
		</div>
		<div class="min-w-96">
			{@render resolutionSelect()}
		</div>
	</div>
	<div class="flex w-full flex-col items-center">
		<DirectoryInput disabled={loading} bind:value={form.outputDirectory} />
	</div>
	<Button type="submit">Download</Button>
</form>
