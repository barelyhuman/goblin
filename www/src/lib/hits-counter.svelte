<script>
	import { onMount } from 'svelte'
	export let count = 0
	const urls = ['https://goblin.barelyhuman.xyz', 'https://goblin.run']
	onMount(async () => {
		urls.forEach((url) => {
			fetch(`https://hits.goblin.run/hits?url=${url}`)
				.then((response) => {
					if (!response.ok) {
						throw response
					}

					return response.json()
				})
				.then((data) => {
					count += data.count
				})
				.catch((error) => {
					console.log(error)
					count += 0
				})
		})
	})
</script>

<div class="h-14 flex justify-end items-center">
	<div class="px-3">
		<p class="text-2xl">{count} <small>hits</small></p>
	</div>
</div>
