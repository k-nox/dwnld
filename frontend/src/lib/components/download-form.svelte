<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Download } from '$lib/wailsjs/go/app/Downloader';
	import { toast } from 'svelte-sonner';
	import DirectoryInput from './directory-input.svelte';

	const form = $state({
		url: '',
		outputTempl: '',
		outputDir: ''
	});

	const resetForm = () => {
		form.url = '';
		form.outputTempl = '';
		form.outputDir = '';
	};

	const download = async () => {
		toast.promise(Download(form.url, form.outputDir, form.outputTempl), {
			loading: 'Downloading...',
			success: () => {
				resetForm();
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
			}
		});
	};
</script>

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
		<DirectoryInput bind:value={form.outputDir} />
	</div>
	<Button type="submit" onclick={download}>Download</Button>
</form>
