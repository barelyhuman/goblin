module github.com/barelyhuman/goblin

go 1.24.0

toolchain go1.24.2

require (
	github.com/Masterminds/semver v1.5.0
	github.com/barelyhuman/go v0.2.2
	github.com/google/go-github/v53 v53.2.0
	github.com/joho/godotenv v1.5.1
	github.com/minio/minio-go/v7 v7.0.77
	github.com/tj/go-semver v1.0.0
	go.uber.org/ratelimit v0.3.1
	golang.org/x/oauth2 v0.27.0
)

// Sec patches
require golang.org/x/crypto v0.37.0 // indirect

require (
	github.com/ProtonMail/go-crypto v1.1.3 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/net v0.36.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)
