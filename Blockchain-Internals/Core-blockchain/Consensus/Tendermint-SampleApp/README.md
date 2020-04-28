How to Start
- clone tendermint and run `make install`
- initialize the keys - `tendermint init`
- add dependencies to vendor - `go mod vendor`
- build kvstore - `go build -mod=vendor`
- start the node - `./kvstore -config "/tmp/kvstore/config/config.toml"`

Should see following logs 
```venkateshmankena@MacBook-Pro kvstore % ./kvstore -config "/tmp/kvstore/config/config.toml"
badger 2020/03/10 15:24:10 INFO: All 0 tables opened in 0s
badger 2020/03/10 15:24:10 INFO: Replaying file id: 0 at offset: 0
badger 2020/03/10 15:24:10 INFO: Replay took: 209.199µs
Starting tendermint kvstore - vendor
Starting tendermint kvstore - vendor
Starting tendermint kvstore - vendor
E[2020-03-10|15:24:11.107] Failed to save AddrBook to file              module=p2p book=config/addrbook.json file=config/addrbook.json err="open config/write-file-atomic-03895281980430084905: no such file or directory"
I[2020-03-10|15:24:12.194] Executed block                               module=state height=180 validTxs=0 invalidTxs=0
I[2020-03-10|15:24:12.205] Committed state                              module=state height=180 txs=0 appHash=
I[2020-03-10|15:24:13.253] Executed block                               module=state height=181 validTxs=0 invalidTxs=0
I[2020-03-10|15:24:13.263] Committed state                              module=state height=181 txs=0 appHash=
I[2020-03-10|15:24:14.311] Executed block                               module=state height=182 validTxs=0 invalidTxs=0
```


Send Transaction 
curl http://localhost:26657/broadcast_tx_commit?tx="tendermint=rocks"

Response
```{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "check_tx": {
      "code": 0,
      "data": null,
      "log": "",
      "info": "",
      "gasWanted": "1",
      "gasUsed": "0",
      "events": [],
      "codespace": ""
    },
    "deliver_tx": {
      "code": 0,
      "data": null,
      "log": "",
      "info": "",
      "gasWanted": "0",
      "gasUsed": "0",
      "events": [],
      "codespace": ""
    },
    "hash": "1B3C5A1093DB952C331B1749A21DCCBB0F6C7F4E0055CD04D16346472FC60EC6",
    "height": "132"
  }
}```


query transaction
curl http://localhost:26657/abci_query?data="tendermint"
Response
```
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "response": {
      "code": 0,
      "log": "exists",
      "info": "",
      "index": "0",
      "key": "dGVuZGVybWludA==",
      "value": "cm9ja3M=",
      "proof": null,
      "height": "0",
      "codespace": ""
    }
  }
}
```


References 
- https://blog.cosmos.network/writing-a-built-in-tendermint-core-app-in-go-a52f3a35ec09


