# Usage

Mnemonic for private key in environment var VRFPRIV. Round and blockseed in base32 format.

```sh
vrfprove [vrf_public_key] [round] [blockseed]

```

Runs VRF proof and writes the base32 encoded proof and random output to STDOUT separated by a space.

# Build

Edit `go.mod` to specify actual `go-algorand` source directory.
