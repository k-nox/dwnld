<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import { inputClassNames } from '$lib/utils';
	import { config } from '$lib/wailsjs/go/models';
	import Label from '../label.svelte';
	import type { FieldProps } from './types';

	let { value = $bindable(), disabled, defaultValue }: FieldProps<config.Resolution> = $props();
	let classes = $derived(inputClassNames(value, defaultValue));
</script>

<div>
	<Label>
		Resolution
		{#snippet info()}
			There's no guarantee the selected resolution will be available. This functions as an upper
			limit for the available resolutions.
		{/snippet}
	</Label>
	<Select.Root type="single" bind:value name="resolution" {disabled}>
		<Select.Trigger class={classes}>
			{value || defaultValue || 'Select a resolution'}
		</Select.Trigger>
		<Select.Content>
			{#each Object.values(config.Resolution) as res}
				<Select.Item value={res}>{res}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>
</div>
