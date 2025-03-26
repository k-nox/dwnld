<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	let { isOpen = $bindable() }: { isOpen: boolean } = $props();
</script>

<Dialog.Root bind:open={isOpen}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Edit Default Settings</Dialog.Title>
		</Dialog.Header>
		<div>Settings will go here</div>
	</Dialog.Content>
</Dialog.Root>
