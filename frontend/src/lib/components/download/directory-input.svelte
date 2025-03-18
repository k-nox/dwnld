<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';

	import { ChooseDirectory } from '$lib/wailsjs/go/download/Downloader';
	import { LogDebug } from '$lib/wailsjs/runtime/runtime';

	interface Props {
		value?: string;
		placeholder?: string;
		disabled: boolean;
	}

	let { value = $bindable(), disabled, placeholder }: Props = $props();
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

<div class="flex rounded-md border border-input shadow-sm">
	<Button variant="secondary" type="button" onclick={chooseDirectory} {disabled}
		>Choose Directory</Button
	>
	<Input
		id="dir"
		class="border-0 focus-visible:ring-0"
		required
		readonly
		{placeholder}
		{disabled}
		type="text"
		bind:value
	/>
</div>
