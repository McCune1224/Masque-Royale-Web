<script lang="ts">
	import { onMount } from 'svelte';

	onMount(() => {
		document.title = 'Day 0';
	});
</script>

<div>
	<h1 class="h1 text-center text-primary-700 dark:text-primary-400">Day 0</h1>
	<div class="py-32 container h-full mx-auto flex justify-center items-center">
		<div class="space-y-10 text-center flex flex-col items-center">
			<div class="flex justify-center items-center">
				<div
					class="w-[300px] h-[300px] bg-surface-300 dark:bg-surface-200 rounded-full flex justify-center items-center shadow-2xl"
				>
					<div
						class="w-[200px] h-[200px] bg-gray-500 dark:bg-surface-700 rounded-full flex justify-center items-center"
					></div>
				</div>
			</div>
		</div>
	</div>
</div>
