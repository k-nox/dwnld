<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';

	import { ChooseDirectory } from '$lib/wailsjs/go/app/App';
	import { LogDebug } from '$lib/wailsjs/runtime/runtime';
	import type { LabeledFieldProps } from './types';
	import type { config } from '$lib/wailsjs/go/models';
	import Label from '../label.svelte';
	import { inputClassNames } from '$lib/utils';

	let {
		value = $bindable(),
		disabled,
		defaultValue,
		label
	}: LabeledFieldProps<config.Download['outputDirectory']> = $props();

	let classes = $derived(inputClassNames(value, defaultValue));

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

<div>
	<Label for={label}>Output Directory</Label>
	<div class={['flex rounded-md border border-input shadow-sm', classes]}>
		<Button variant="secondary" type="button" onclick={chooseDirectory} {disabled}>
			Choose Directory
		</Button>
		<Input
			id={label}
			class="border-0 focus-visible:ring-0"
			readonly
			placeholder={defaultValue}
			{disabled}
			type="text"
			bind:value
		/>
	</div>
</div>
