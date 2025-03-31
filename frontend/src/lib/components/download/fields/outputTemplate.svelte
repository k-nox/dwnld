<script lang="ts">
	import Label from '../label.svelte';
	import type { config } from '$lib/wailsjs/go/models';
	import type { MouseEventHandler } from 'svelte/elements';
	import type { LabeledFieldProps } from './types';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import Input from '$lib/components/ui/input/input.svelte';
	import { inputClassNames } from '$lib/utils';

	let {
		value = $bindable(),
		disabled,
		label,
		defaultValue
	}: LabeledFieldProps<config.Download['outputTemplate']> = $props();

	let classes = $derived(inputClassNames(value, defaultValue));

	const handleOutsideClick: MouseEventHandler<HTMLAnchorElement> = (e) => {
		e.preventDefault();
		BrowserOpenURL(e.currentTarget.href);
	};
</script>

<div>
	<Label for={label}>
		Output Template
		{#snippet info()}
			<div>
				For details and examples of acceptable formats, check out <a
					href="https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#output-template"
					class="underline underline-offset-4"
					onclick={handleOutsideClick}>yt-dlp's docs</a
				>, or leave the field as is to use the default.
			</div>
		{/snippet}
	</Label>

	<Input class={classes} id={label} type="text" placeholder={defaultValue} {disabled} bind:value />
</div>
