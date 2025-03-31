<script lang="ts">
	import { Switch } from '$lib/components/ui/switch';
	import type { Snippet } from 'svelte';
	import Label from './label.svelte';

	interface Props {
		id: string;
		checked: boolean;
		disabled: boolean;
		info?: Snippet;
		label: string;
	}

	let { id, info, label, checked = $bindable(), disabled }: Props = $props();
</script>

<div class="flex w-1/2 items-center gap-x-2">
	<Switch {id} bind:checked {disabled} />
	<Label for={id} {info}>
		{label}
	</Label>
</div>
