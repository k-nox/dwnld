<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import { app } from '$lib/wailsjs/go/models';

	let { value = $bindable() }: { value: app.Resolution | undefined } = $props();
</script>

<Select.Root type="single" bind:value name="resolution" required>
	<Select.Trigger>
		{value ? value : 'Select a resolution.'}
	</Select.Trigger>
	<Select.Content>
		{#each Object.values(app.Resolution) as res}
			<Select.Item value={res}>{res}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>
