<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import { config } from '$lib/wailsjs/go/models';

	interface Props {
		value?: config.Resolution;
		disabled: boolean;
		defaultValue?: config.Resolution;
	}

	let { value = $bindable(), disabled, defaultValue }: Props = $props();
</script>

<Select.Root type="single" bind:value name="resolution" required {disabled}>
	<Select.Trigger>
		{value || defaultValue || 'Select a resolution'}
	</Select.Trigger>
	<Select.Content>
		{#each Object.values(config.Resolution) as res}
			<Select.Item value={res}>{res}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>
