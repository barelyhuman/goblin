## Usage

Install `package` with optional `@version` and `options`:

```command
curl -sf {{.ORIGIN_URL}}/<package>[@version] | [...options] sh
```

## API

`package` - Complete module path

```sh
github.com/barely## Usage

Install `package` with optional `@version` and `options`:

```command
curl -sf {{.ORIGIN_URL}}/<package>[@version] | [...options] sh
```

## API

`package` - Complete module path

```sh
github.com/barelyhuman/commitlog
gopkg.in/yaml.v2
```

`version` - Exact or partial version range, optionally prefixed with "v"

```sh
# Install the latest version
<package>

# Install v1.2.3
<package>@v1.2.3

# Install v3.x.x
<package>@v3
```

### Options

> Control Goblin's behavior with environment variables

`PREFIX` - Change installation location (default: `/usr/local/bin`)

```sh
# Install to /tmp
... | PREFIX=/tmp sh
```

`OUT` - Rename the resulting binary (default: `<package name>`)

```sh
# Export Windows executable
... | OUT=example.exe sh
```

`CMD_PATH` - Path to the binary package (default: "")

```sh
# Export Windows executable
... | CMD_PATH="/cmd/example" sh
```

## Examples

Install the latest version:

```command
curl -sf {{.ORIGIN_URL}}/github.com/rakyll/hey | sh
```

Specify package version:

```command
curl -sf {{.ORIGIN_URL}}/github.com/barelyhuman/statico@v0.0.7 | sh
```

Or use commit hashes:

```command
curl -sf {{.ORIGIN_URL}}/github.com/barelyhuman/commitlog@bba8d7a63d622e4f12dbea9722b647cd985be8ad | sh
```

Use alternative sources:

```command
curl -sf {{.ORIGIN_URL}}/golang.org/x/tools/godoc | sh
```

Specify nested packages

> Note: nested package expect the path to be a `package main` file with a `main`
> call. If you use something like `spf13/cobra` then check the 2nd example.

```command
curl -sf {{.ORIGIN_URL}}/vito/bass/cmd/bass | sh
```

```command
curl -sf {{.ORIGIN_URL}}/gnorm.org/gnorm | CMD_PATH="/cli" PREFIX=./bin sh
```

## How does it work?

Each request resolves the needed tags and versions from
[proxy.golang.org](https://proxy.golang.org). If no module is found, you can try
replacing the version with a commit hash on supported platforms, e.g. GitHub.

The response of this request is a Golang binary compiled for the requested
operating system, architecture, package version, and the binary's name—using Go
1.17.x via the official [Docker image](https://hub.docker.com/_/golang).

**Example response**

```sh
{{.ORIGIN_URL}}/binary/github.com/rakyll/hey?os=darwin&arch=amd64&version=v0.1.3&out=hey
```

_Note: compilation is limited to 200 seconds due to timeout restrictions._
human/commitlog
gopkg.in/yaml.v2
```

`version` - Exact or partial version range, optionally prefixed with "v"

```sh
# Install the latest version
<package>

# Install v1.2.3
<package>@v1.2.3

# Install v3.x.x
<package>@v3
```

### Options

> Control Goblin's behavior with environment variables

`PREFIX` - Change installation location (default: `/usr/local/bin`)

```sh
# Install to /tmp
... | PREFIX=/tmp sh
```

`OUT` - Rename the resulting binary (default: `<package name>`)

```sh
# Export Windows executable
... | OUT=example.exe sh
```

`CMD_PATH` - Path to the binary package (default: "")

```sh
# Export Windows executable
... | CMD_PATH="/cmd/example" sh
```

## Examples

Install the latest version:

```command
curl -sf {{.ORIGIN_URL}}/github.com/rakyll/hey | sh
```

Specify package version:

```command
curl -sf {{.ORIGIN_URL}}/github.com/barelyhuman/statico@v0.0.7 | sh
```

Or use commit hashes:

```command
curl -sf {{.ORIGIN_URL}}/github.com/barelyhuman/commitlog@bba8d7a63d622e4f12dbea9722b647cd985be8ad | sh
```

Use alternative sources:

```command
curl -sf {{.ORIGIN_URL}}/golang.org/x/tools/godoc | sh
```

Specify nested packages

> Note: nested package expect the path to be a `package main` file with a `main`
> call. If you use something like `spf13/cobra` then check the 2nd example.

```command
curl -sf {{.ORIGIN_URL}}/vito/bass/cmd/bass | sh
```

```command
curl -sf {{.ORIGIN_URL}}/gnorm.org/gnorm | CMD_PATH="/cli" PREFIX=./bin sh
```

## How does it work?

Each request resolves the needed tags and versions from
[proxy.golang.org](https://proxy.golang.org). If no module is found, you can try
replacing the version with a commit hash on supported platforms, e.g. GitHub.

The response of this request is a Golang binary compiled for the requested
operating system, architecture, package version, and the binary's name—using Go
1.17.x via the official [Docker image](https://hub.docker.com/_/golang).

**Example response**

```sh
{{.ORIGIN_URL}}/binary/github.com/rakyll/hey?os=darwin&arch=amd64&version=v0.1.3&out=hey
```

_Note: compilation is limited to 200 seconds due to timeout restrictions._
