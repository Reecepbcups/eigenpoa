# Eigen PoA

Eigenlayer AVS meets Cosmos-SDK Proof of Authority.

Contract is based off of `https://github.com/Layr-Labs/hello-world-avs` [001dc6e944280559dfb44f75faf5102349a61d8e](https://github.com/Layr-Labs/hello-world-avs/commit/001dc6e944280559dfb44f75faf5102349a61d8e)

```bash
git clone git@github.com:Reecepbcups/eigen-poa-avs.git

# https://goethereumbook.org/smart-contract-compile/
# https://goethereumbook.org/en/block-query/

cd eigen-poa-avs
npm run extract:abis # to the ./abis dir

cd ..
abigen --abi=./eigen-poa-avs/abis/HelloWorldServiceManager.json --pkg=manager --out=./x/avs/keeper/manager/HelloWorldServiceManager.go  # TODO: how to do others as well?
abigen --abi=./eigen-poa-avs/abis/ECDSAStakeRegistry.json --pkg=stake_registry --out=./x/avs/keeper/stake_registry/ECDSAStakeRegistry.go  # TODO: how to do others as well?

```

---

## Generated With [Spawn](https://github.com/rollchains/spawn)

## Module Scaffolding

- `spawn module new <name>` *Generates a Cosmos module template*

## Content Generation

- `make proto-gen` *Generates go code from proto files, stubs interfaces*

## Testnet

- `make testnet` *IBC testnet from chain <-> local cosmos-hub*
- `make sh-testnet` *Single node, no IBC. quick iteration*
- `local-ic chains` *See available testnets from the chains/ directory*
- `local-ic start <name>` *Starts a local chain with the given name*

## Local Images

- `make install`      *Builds the chain's binary*
- `make local-image`  *Builds the chain's docker image*

## Testing

- `go test ./... -v` *Unit test*
- `make ictest-*`  *E2E testing*

## Webapp Template

Generate the template base with spawn. Requires [npm](https://nodejs.org/en/download/package-manager) and [yarn](https://classic.yarnpkg.com/lang/en/docs/install) to be installed.

- `make generate-webapp` *[Cosmology Webapp Template](https://github.com/cosmology-tech/create-cosmos-app)*

Start the testnet with `make testnet`, and open the webapp `cd ./web && yarn dev`
