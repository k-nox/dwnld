<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import Info from 'lucide-svelte/icons/info';
	import { buttonVariants } from './ui/button';
	import type { Snippet } from 'svelte';

	let { children }: { children: Snippet } = $props();
</script>

<Popover.Root>
	<Popover.Trigger class={buttonVariants({ variant: 'ghost', size: 'icon' })}>
		<Info />
		<span class="sr-only">More informtation on output formats</span>
	</Popover.Trigger>
	<Popover.Content class="text-sm text-primary" side="top">{@render children()}</Popover.Content>
</Popover.Root>
