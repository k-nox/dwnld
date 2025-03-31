<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import type { config } from '$lib/wailsjs/go/models';
	import Directory from '../download/fields/directory.svelte';
	import OutputTemplate from '../download/fields/outputTemplate.svelte';
	import Resolution from '../download/fields/resolution.svelte';
	import { toast } from 'svelte-sonner';
	import { WindowReload } from '$lib/wailsjs/runtime/runtime';
	import { UpdateSettings } from '$lib/wailsjs/go/app/App';

	interface Props {
		isOpen: boolean;
		defaults: config.Download;
	}
	let { isOpen = $bindable(), defaults }: Props = $props();

	let form = $state({ ...defaults });

	let isChanged = $derived.by(() => {
		let key: keyof config.Download;
		for (key in defaults) {
			if (form[key] !== defaults[key]) {
				return true;
			}
		}
		return false;
	});

	const updateDefaults = async (e: SubmitEvent) => {
		e.preventDefault();
		try {
			await UpdateSettings({ download: form });
			// TODO: this doesn't feel like the "correct" way to do this
			WindowReload();
		} catch (e: unknown) {
			let msg = '';
			if (typeof e === 'string') {
				msg = e;
			} else if (e instanceof Error) {
				msg = e.message;
			}
			toast.error(msg);
		}
	};
</script>

<Dialog.Root bind:open={isOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Edit Default Settings</Dialog.Title>
		</Dialog.Header>
		<form onsubmit={updateDefaults}>
			<div class="flex flex-col gap-y-4">
				<OutputTemplate
					bind:value={form.outputTemplate}
					defaultValue={defaults.outputTemplate}
					disabled={false}
					label="defaultOutputTemplate"
				/>
				<Resolution
					bind:value={form.targetResolution}
					disabled={false}
					defaultValue={defaults.targetResolution}
				/>
				<Directory
					bind:value={form.outputDirectory}
					defaultValue={defaults.outputDirectory}
					disabled={false}
					label="defaultOutputDirectory"
				/>
				<Button type="submit" class="self-center" disabled={!isChanged}>Save Defaults</Button>
			</div>
		</form>
	</Dialog.Content>
</Dialog.Root>
