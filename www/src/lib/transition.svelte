<script>
	import { onMount, tick } from 'svelte'

	/**
	 * fallback to display content immediately if javascript is disabled
	 */
	export let showWithoutJavascript = true
	export let active = ''
	export let from = ''
	export let to = ''

	let mounted = false
	onMount(async () => {
		mounted = true
		await tick()
	})
</script>

<div class="{active} {mounted ? to : from}" class:showWithoutJavascript>
	<slot />
</div>
