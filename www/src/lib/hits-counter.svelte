<script>
	import { onMount } from 'svelte'
	export let count = 0
	const url = 'https://goblin.reaper.im'
	onMount(async () => {
		fetch(
			`https://api.hits.link/v1/hits?border=square&json=true&bgRight=27272a&bgLeft=27272a&url=${url}`
		)
			.then((response) => {
				console.log({ response })

				if (!response.ok) {
					throw response
				}

				return response.json()
			})
			.then((data) => {
				count = data.data.hits
			})
			.catch((error) => {
				console.log(error)
				count = 0
			})
	})
</script>

<div class="h-14 flex justify-end items-center">
	<div class="px-3">
		<p class="text-2xl">{count} <small>hits</small></p>
	</div>
</div>
