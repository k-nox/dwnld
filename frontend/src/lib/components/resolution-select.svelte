<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import { app } from '$lib/wailsjs/go/models';

	let { value = $bindable(), disabled }: { value: app.Resolution | undefined; disabled: boolean } =
		$props();
</script>

<Select.Root type="single" bind:value name="resolution" required {disabled}>
	<Select.Trigger>
		{value ? value : 'Select a resolution.'}
	</Select.Trigger>
	<Select.Content>
		{#each Object.values(app.Resolution) as res}
			<Select.Item value={res}>{res}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>
