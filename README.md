# Usage

Mnemonic for private key in environment var VRFPRIV. Round and blockseed in base32 format.

```sh
vrfprove [round] [blockseed]

```

Runs VRF proof and outputs the base32 encoded proof and hash to STDOUT separated by a space.

# Build

Edit `go.mod` to specify actual `go-algorand` source directory.
