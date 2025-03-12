<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Button } from '$lib/components/ui/button';
	import { ChooseDirectory } from '$lib/wailsjs/go/app/Downloader';
	import { LogDebug } from '$lib/wailsjs/runtime/runtime';

	let { value = $bindable() }: { value: string } = $props();
	const chooseDirectory = async () => {
		try {
			value = await ChooseDirectory();
		} catch (e: unknown) {
			if (typeof e === 'string') {
				LogDebug(e);
			} else if (e instanceof Error) {
				LogDebug(e.message);
			}
		}
	};
</script>

<Label for="dir" class="m-1">Output Directory</Label>
<div class="flex w-1/2 rounded-md border border-input shadow-sm">
	<Button variant="secondary" type="button" onclick={chooseDirectory}>Choose Directory</Button>
	<Input id="dir" class="border-0 focus-visible:ring-0" readonly type="text" bind:value />
</div>
