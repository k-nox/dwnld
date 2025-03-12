<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { ChooseDirectory, Download } from '$lib/wailsjs/go/app/Downloader';
	import { LogDebug } from '$lib/wailsjs/runtime/runtime';
	const form = $state({
		url: '',
		outputTempl: '',
		outputDir: ''
	});

	let status = $state('');

	const resetForm = () => {
		form.url = '';
		form.outputTempl = '';
		form.outputDir = '';
	};

	const download = async () => {
		status = 'loading';
		try {
			await Download(form.url, form.outputDir, form.outputTempl);
		} catch (e: unknown) {
			if (typeof e === 'string') {
				LogDebug(e);
			} else if (e instanceof Error) {
				LogDebug(e.message);
			}
		}
		resetForm();
		status = '';
	};

	const chooseDirectory = async () => {
		try {
			form.outputDir = await ChooseDirectory();
		} catch (e: unknown) {
			if (typeof e === 'string') {
				LogDebug(e);
			} else if (e instanceof Error) {
				LogDebug(e.message);
			}
		}
	};
</script>

<div class="h-full">
	{#if status === 'loading'}
		<div class="flex min-h-full flex-col items-center justify-center">loading...</div>
	{:else}
		<form class="flex min-h-full flex-col items-center justify-center gap-4">
			<div class="flex w-full justify-center gap-4">
				<div class="w-1/3">
					<Label for="url" class="m-1">URL</Label>
					<Input id="url" type="url" placeholder="video url" bind:value={form.url} />
				</div>
				<div class="w-1/3">
					<Label for="templ" class="m-1">Output Template</Label>
					<Input
						id="templ"
						type="text"
						placeholder="%(title)s [%(id)s].%(ext)s"
						bind:value={form.outputTempl}
					/>
				</div>
			</div>
			<div class="flex w-full flex-col items-center">
				<Label for="dir" class="m-1">Output Directory</Label>
				<div class="flex w-1/2 rounded-md border border-input shadow-sm">
					<Button variant="secondary" type="button" onclick={chooseDirectory}
						>Choose Directory</Button
					>
					<Input
						id="dir"
						class="border-0 focus-visible:ring-0"
						readonly
						type="text"
						bind:value={form.outputDir}
					/>
				</div>
			</div>
			<Button type="submit" onclick={download}>Download</Button>
		</form>
	{/if}
</div>
