<script>
	import Command from '$lib/command.svelte'
	import InlineCode from '$lib/inline-code.svelte'
	import Transition from '$lib/transition.svelte'
</script>

<section class="py-20 min-h-[calc(100vh-5rem)]">
	<Transition
		active="delay-[100ms] duration-500 transition-[opacity,transform] ease-out"
		from="opacity-0 translate-y-3"
		to="opacity-1 translate-y-0"
	>
		<h2 class="font-bold text-4xl sm:text-6xl leading-[1.1] sm:leading-[1.25]">
			Golang binaries in a curl, built by <span
				class="bg-clip-text text-transparent bg-gradient-to-r from-[#fd7e14] to-[#fab005]"
			>
				goblins
			</span>
		</h2>
	</Transition>

	<div class="pt-10 space-y-4">
		<Transition
			active="delay-[200ms] duration-500 transition-[opacity,transform] ease-out"
			from="opacity-0 translate-y-3"
			to="opacity-1 translate-y-0"
		>
			<Command
				command="curl -sf https://goblin.reaper.im/github.com/rakyll/hey@v0.1.3 | sh"
			/>
		</Transition>

		<div class="pt-4 space-x-4 flex items-center">
			<Transition
				active="delay-[250ms] duration-500 transition-[opacity,transform] ease-out"
				from="opacity-0 translate-y-3"
				to="opacity-1 translate-y-0"
			>
				<a
					href="#intro"
					class="px-4 py-2 font-medium text-sm text-on-accent bg-accent border border-accent rounded-md flex items-center hover:opacity-90"
				>
					Get Started
				</a>
			</Transition>

			<Transition
				active="delay-[300ms] duration-500 transition-[opacity,transform] ease-out"
				from="opacity-0 translate-y-3"
				to="opacity-1 translate-y-0"
			>
				<a
					rel="external"
					target="_blank"
					href="https://github.com/barelyhuman/goblin"
					class="px-4 py-2 font-medium text-sm text-subtle border rounded-md flex items-center hover:bg-surface"
				>
					Source
				</a>
			</Transition>
		</div>
	</div>
</section>

<section class="space-y-20">
	<div class="space-y-6">
		<h3 id="intro" class="font-semibold text-2xl">Introduction</h3>

		<div>
			<p>
				Go binaries install from the command line without requiring Go to be
				installed locally. Goblin streamlines this process by cross-compiling
				binaries on request and caching for subsequent installations.
			</p>
		</div>
	</div>

	<div class="space-y-6">
		<h3 id="usage" class="font-semibold text-2xl">Usage</h3>

		<div class="space-y-6">
			<div class="space-y-3">
				<p>
					Install <InlineCode>PKG</InlineCode> with optional <InlineCode
						>@VERSION</InlineCode
					>
				</p>

				<Command
					command="curl -sf https://goblin.reaper.im/github.com/<PKG>[@VERSION] | sh"
				/>

				<p class="text-muted text-xs">
					Note: <span class="font-medium text-subtle">github.com</span> is required
				</p>
			</div>

			<div class="space-y-3">
				<p>
					Install <InlineCode>PKG</InlineCode> to directory <InlineCode
						>PREFIX</InlineCode
					>
				</p>

				<Command
					command="curl -sf https://goblin.reaper.im/github.com/<PKG>[@VERSION] | PREFIX=/tmp sh"
				/>

				<p class="text-muted text-xs">
					The directory will be created if it does not exist
				</p>
			</div>
		</div>
	</div>

	<div class="space-y-6">
		<h3 id="examples" class="font-semibold text-2xl">Examples</h3>

		<div class="space-y-6">
			<div class="space-y-3">
				<p>
					Install the latest version of <InlineCode>hey</InlineCode>
				</p>

				<Command
					command="curl -sf https://goblin.reaper.im/github.com/rakyll/hey | sh"
				/>
			</div>

			<div class="space-y-3">
				<p>
					Install <InlineCode>statico@v0.0.7</InlineCode>
				</p>

				<Command
					command="curl -sf https://goblin.reaper.im/github.com/barelyhuman/statico@v0.0.7 | sh"
				/>

				<p class="text-muted text-xs">
					Version can be exact –
					<span class="font-medium text-subtle">v1.2.3</span> or implicit –
					<span class="font-medium text-subtle">v1</span> with the
					<span class="font-medium text-subtle">v</span> being optional
				</p>
			</div>
		</div>
	</div>

	<div class="space-y-6">
		<h3 id="how-does-it-work" class="font-semibold text-2xl">
			How does it work?
		</h3>

		<div class="space-y-6">
			<div class="space-y-3">
				<p>
					Each request fetches the latest tag from GitHub and responds with a
					shell script which performs a second request, populated with the
					resolved version and architecture as shown here:
				</p>

				<Command
					hidePrefix
					command="https://goblin.reaper.im/binary/github.com/rakyll/hey?os=darwin&arch=amd64&version=v0.1.3"
				/>

				<p>
					The response of this request is a Golang binary compiled for the
					requested os, architecture, and package version.
				</p>

				<p class="text-muted text-xs">
					The result will be cached in a CDN for subsequent requests in one of a
					future version
				</p>
			</div>
		</div>
	</div>

	<div class="space-y-6">
		<h3 id="caveats" class="font-semibold text-2xl">Caveats</h3>

		<ul class="list-disc list-inside">
			<li>Go package must have at least one git tag</li>
			<li>
				Go package must compile in under 100 seconds due to CDN limitations
			</li>
		</ul>
	</div>

	<div class="space-y-6">
		<h3 id="faq" class="font-semibold text-2xl">FAQ</h3>

		<div class="space-y-6">
			<div class="space-y-3">
				<h4 class="font-semibold">What's wrong with go-get?</h4>
				<p>
					The user must have Go installed to use go-get. Goblin makes your
					program accessible to all, including situations where the Go toolchain
					may be unavailable.
				</p>
			</div>

			<div class="space-y-3">
				<h4 class="font-semibold">Which version of Go is used?</h4>
				<p>Currently Go 1.16.x via the official golang:1.16 Docker image.</p>
			</div>
		</div>
	</div>
</section>
