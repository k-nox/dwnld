<script lang="ts">
	import type { ComponentProps, Snippet } from 'svelte';
	import { Label } from '../ui/label';
	import InfoPopover from './info-popover.svelte';

	interface Props extends ComponentProps<typeof Label> {
		info?: Snippet;
		children: Snippet;
	}

	let { children, info, ...props }: Props = $props();
</script>

<Label {...props}>
	{@render children()}
	{#if info}
		<InfoPopover>{@render info()}</InfoPopover>
	{/if}
</Label>
