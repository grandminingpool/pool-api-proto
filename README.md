# pool-api-proto

Protobuf schemas for a single pool API in Grand Pool.

## Usage
```shell
go get github.com/grandminingpool/pool-api-proto/v2
```

Import the generated packages from `github.com/grandminingpool/pool-api-proto/v2/generated/...`,
e.g. `github.com/grandminingpool/pool-api-proto/v2/generated/pool`.

## Merged mining
A pool mines one primary blockchain (`PoolInfo.blockchain`) and zero or more auxiliary (aux)
blockchains with the same work.

- The keys of `PoolInfo.aux_payout_terms` are the aux blockchains the pool merge-mines.
- Pool block counters, miner balances and profitability come for every blockchain at once:
  `PoolStats.blocks` / `aux_blocks`, `MinerBalance.balance` / `aux_balances` and
  `MinerProfitabilityPoint.balance` / `aux_balances`. The aux maps have an entry for every aux
  blockchain.
- The miners list (`GetMiners`) is one list for all blockchains. `Miner.blocks_count` and
  `solo_blocks_count` are totals over all blockchains; aux blocks found by the miner's shares count
  whether or not the miner was rewarded for them. The per-blockchain breakdown is in the blocks
  lists (`GetBlocks`, `GetSoloBlocks`) and `PoolStats`.
- Other per-blockchain data (payouts, blocks, solo blocks, rounds, network difficulty) is requested
  one blockchain at a time through the `blockchain` request field. An unset or empty one means the
  primary blockchain.
- `NetworkInfo` describes the primary blockchain only.
- Clients send only `blockchain` values listed in `PoolInfo.aux_payout_terms`. A server returns
  `InvalidArgument` for any other non-empty value, including on a pool without merged mining.
- Amounts are in each blockchain's own smallest unit.
- Miner logins use `COIN1,COIN2.RIG` (primary address, aux address, worker). `Payout.address` is
  the address paid when it differs from `miner` (e.g. the miner's aux address); unset otherwise.
- `ValidateAddress` validates primary addresses only.
- The history of an aux blockchain the pool no longer merge-mines is not available: requests for a
  blockchain not in `PoolInfo.aux_payout_terms` return `InvalidArgument`.

## Generating code
```shell
go generate
```

Use the generator versions the committed code was generated with, so regenerating produces no
unrelated changes:
- protoc 26.1 (reported as `v5.26.1` in the generated files);
- protoc-gen-go v1.28.1;
- protoc-gen-go-grpc v1.2.0.

## Migrating from v1
v2 is a major version: the module path is now `github.com/grandminingpool/pool-api-proto/v2`, so
imports change from `github.com/grandminingpool/pool-api-proto/generated/...` to
`github.com/grandminingpool/pool-api-proto/v2/generated/...`.

Wire compatibility with v1 is not kept. Moved and changed fields:

| v1 | v2 |
|---|---|
| `PoolInfo.address_format` / `fee` / `payouts_info` | `PoolInfo.payout_terms.address_format` / `fee` / `payouts_info` |
| `PoolStats.blocks_count_24h` / `last_found_block_height` / `last_found_block_at` | `PoolStats.blocks.blocks_count_24h` / `last_found_block_height` / `last_found_block_at` |
| `Miner.blocks_count` / `solo_blocks_count`: primary blocks only | Same fields: totals over all blockchains (meaning only; names and numbers unchanged) |
| `MinerProfitabilityPoint.date` = 2 | `MinerProfitabilityPoint.date` = 3 (wire only; the name is unchanged) |
| `GetPayoutsRequest` / `GetBlocksRequest` / `GetSoloBlocksRequest`: `pagination` = 3 | `pagination` = 4; the new `blockchain` takes 3 (wire only) |
