<div class="mx-auto max-w-6xl">
	<header class="px-6 sm:px-10 h-20 flex items-center justify-between">
		<h1 class="font-semibold text-subtle"><a href="/">Goblin</a></h1>

		<nav>
			<ul>
				<li>
					<a
						rel="external"
						target="_blank"
						href="https://github.com/barelyhuman/goblin"
						class="font-medium text-sm text-muted hover:text-text"
					>
						GitHub
					</a>
				</li>
			</ul>
		</nav>
	</header>
</div>

<div class="mx-auto max-w-content">
	<main class="mt-0 sm:mt-10 p-6 sm:p-10">
		<slot />
	</main>
</div>

<style lang="postcss" global>
	@import url('https://rsms.me/inter/inter.css');

	@tailwind base;
	@tailwind components;
	@tailwind utilities;

	:root {
		--base: theme('colors.zinc.50');
		--surface: theme('colors.zinc.100');
		--overlay: theme('colors.zinc.200');
		--muted: theme('colors.zinc.500');
		--subtle: theme('colors.zinc.700');
		--text: theme('colors.zinc.900');
		--accent: #fd7e14;
		--on-accent: white;
	}

	@media (prefers-color-scheme: dark) {
		:root {
			--base: theme('colors.zinc.900');
			--surface: theme('colors.zinc.800');
			--overlay: theme('colors.zinc.700');
			--muted: theme('colors.zinc.500');
			--subtle: theme('colors.zinc.400');
			--text: theme('colors.zinc.200');
			--accent: #ffa94d;
			--on-accent: black;
		}
	}

	body {
		@apply text-text bg-base antialiased;
	}
</style>
