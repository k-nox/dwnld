<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import type { config } from '$lib/wailsjs/go/models';
	import Directory from '../download/fields/directory.svelte';
	import OutputTemplate from '../download/fields/outputTemplate.svelte';
	import Resolution from '../download/fields/resolution.svelte';
	interface Props {
		isOpen: boolean;
		defaults: config.Download;
	}
	let { isOpen = $bindable(), defaults }: Props = $props();

	let form = $state({ ...defaults });

	const updateDefaults = async (e: SubmitEvent) => {
		e.preventDefault();
		console.log('update defaults here');
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
			</div>
		</form>
	</Dialog.Content>
</Dialog.Root>
