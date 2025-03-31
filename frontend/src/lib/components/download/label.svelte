<script lang="ts">
	import type { ComponentProps, Snippet } from 'svelte';
	import { Label } from '../ui/label';
	import InfoPopover from './info-popover.svelte';

	interface Props extends ComponentProps<typeof Label> {
		info?: Snippet;
		children: Snippet;
	}

	let { children, info, class: className, ...props }: Props = $props();
</script>

<Label {...props} class={['flex w-full items-start justify-between gap-x-3  pb-1', className]}>
	{@render children()}
	{#if info}
		<span class="flex h-[1lh] items-center">
			<InfoPopover>{@render info()}</InfoPopover>
		</span>
	{/if}
</Label>
