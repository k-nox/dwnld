<script lang="ts">
	import type { Snippet } from 'svelte';
	import Info from 'lucide-svelte/icons/info';
	import * as Popover from '$lib/components/ui/popover';
	import { buttonVariants } from '$lib/components/ui/button';

	let { children }: { children: Snippet } = $props();
</script>

<Popover.Root>
	<Popover.Trigger
		class={[buttonVariants({ variant: 'ghost', size: 'icon' }), 'h-[1.1lh] w-[1.1lh]']}
	>
		<Info />
		<span class="sr-only">More informtation on output formats</span>
	</Popover.Trigger>
	<Popover.Content class="text-sm text-primary" side="top">{@render children()}</Popover.Content>
</Popover.Root>
