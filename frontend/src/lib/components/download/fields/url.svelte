<script lang="ts">
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '../label.svelte';
	import type { LabeledFieldProps } from './types';

	let { value = $bindable(), disabled, label }: LabeledFieldProps<string> = $props();
</script>

<Label for={label}>
	URL
	{#snippet info()}
		Accepts both single video and playlist URLs.
	{/snippet}
</Label>
<Input id={label} type="url" placeholder="video url" bind:value required {disabled} />
