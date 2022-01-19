<script>
	export let command = ''
	export let hidePrefix = false
</script>

<div
	class="h-14 bg-surface rounded-md flex items-center overflow-x-scroll scrollbar-hide whitespace-nowrap"
>
	<div class="px-3 font-mono text-sm">
		{#if !hidePrefix}<span class="text-muted pointer-events-none">$</span>{/if}
		{command}
	</div>
</div>
