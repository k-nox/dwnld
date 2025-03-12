<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import Info from 'lucide-svelte/icons/info';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import type { MouseEventHandler } from 'svelte/elements';
	import { buttonVariants } from './ui/button';

	const handleClick: MouseEventHandler<HTMLAnchorElement> = (e) => {
		e.preventDefault();
		BrowserOpenURL(e.currentTarget.href);
	};
</script>

<Popover.Root>
	<Popover.Trigger class={buttonVariants({ variant: 'ghost', size: 'icon' })}>
		<Info />
		<span class="sr-only">More informtation on output formats</span>
	</Popover.Trigger>
	<Popover.Content class="text-sm text-primary" side="top"
		>For details and examples of acceptable formats, check out <a
			href="https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#output-template"
			class="underline underline-offset-4"
			onclick={handleClick}>yt-dlp's docs</a
		>, or leave the field as is to use the default.</Popover.Content
	>
</Popover.Root>
