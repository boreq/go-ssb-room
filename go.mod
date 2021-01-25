module go.mindeco.de/ssb-rooms

go 1.15

require (
	github.com/go-kit/kit v0.10.0
	github.com/gorilla/websocket v1.4.1
	github.com/keks/nocomment v0.0.0-20181007001506-30c6dcb4a472
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	go.cryptoscope.co/muxrpc/v2 v2.0.0-20210118124540-d38ef6e0e8c7
	go.cryptoscope.co/netwrap v0.1.1
	go.cryptoscope.co/secretstream v1.2.2
	go.mindeco.de/ssb-refs v0.1.1-0.20210108133850-cf1f44fea870
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
)

// We need our internal/extra25519 since agl pulled his repo recently.
// Issue: https://github.com/cryptoscope/ssb/issues/44
// Ours uses a fork of x/crypto where edwards25519 is not an internal package,
// This seemed like the easiest change to port agl's extra25519 to use x/crypto
// Background: https://github.com/agl/ed25519/issues/27#issuecomment-591073699
// The branch in use: https://github.com/cryptix/golang_x_crypto/tree/non-internal-edwards
replace golang.org/x/crypto => github.com/cryptix/golang_x_crypto v0.0.0-20200924101112-886946aabeb8

replace go.cryptoscope.co/muxrpc/v2 => /home/cryptix/go-repos/muxrpc
