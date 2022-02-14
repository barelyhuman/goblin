<script>
	const primary =
		'px-4 py-2 font-medium text-md text-on-accent bg-accent border border-accent rounded-md flex items-center hover:opacity-90'
	const secondary =
		'px-4 py-2 font-medium text-md text-subtle border rounded-md flex items-center hover:bg-surface'
	export let href
	export let variant = 'primary'
	export let external = false
	export let additionalClasses = ''
	let classList = [
		variant === 'primary' ? primary : secondary,
		additionalClasses,
	].join(' ')
</script>

<a
	{href}
	target={external ? '_blank' : null}
	rel={external ? 'noopener noreferrer' : null}
	class={classList}><slot /></a
>
