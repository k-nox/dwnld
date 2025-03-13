<script lang="ts">
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import type { MouseEventHandler } from 'svelte/elements';

	const handleClick: MouseEventHandler<HTMLAnchorElement> = (e) => {
		e.preventDefault();
		BrowserOpenURL(e.currentTarget.href);
	};
</script>

<div>
	For details and examples of acceptable formats, check out <a
		href="https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#output-template"
		class="underline underline-offset-4"
		onclick={handleClick}>yt-dlp's docs</a
	>, or leave the field as is to use the default.
</div>
