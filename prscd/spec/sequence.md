![graph](https://www.websequencediagrams.com/cgi-bin/cdraw?lz=dGl0bGUgUHJlc2VuY2VqcyB2MgoKYWN0b3IgQWxpY2UABQdCb2IKAAsFLT5Cb2I6IEJvYiBpcyBvbmxpbmUsACUGIGpvaW4gdGhpcyBjaGFubmVsLCBzeW5jIHN0YXRlLgoKbm90ZSBvdmUAVAcsIEJvYiwgcHJzY2Q6IFN0ZXAgMSAtIEF1dGgAZAgAFQdXZWJTb2NrZXQvV2ViVHJhbnNwb3J0LCB3aXRoIGBhdXRoYCBhbmQgYGlkYABgBnJpZ2h0IG9mAFYIAFAFAGUFLT4rQXV0aFN2YwAQBlN2YwoACgctLT4tACULUmVzdWx0AC4IAIIJBTogNHh4IGlmIGF1dGggZmFpbCwgVXBncmFkZSBpZiBzdWNjZXNzAIFGJDIgLSBKb2luIEMAgiUGAIJWCCsAgggHYACCPAdfam9pbmAAgTQHLT4tAIEABwARDiBBQ0sAgxcIKDIpAIMdBWBwZWVyXwCDGwZgAIIkBnN0YXRfb2JqIHBheWxvYWQAghwPAINUBWFkZACDRwd0bwCDWAcgdXNlcnMgY2FjaGUAHxRzaG91bGQgcmVzcG9uZAByBwCDbwVgCkJvYgCBDgUAgS8IABAMAIQKBWxlZgCDIwUADhMAUQhhZGQAhGsFdG8AfQcAgVUWAC4HYXQgYW55IHRpbQCBGBV1cGRhdGUAhSoHAIIJCACEdyQzIC0gQnJvYWRjYXN0IGN1c3RvbWUgZGF0YQCCaRFkYXRhYACFSyQ0IC0gTGVhdmUAg3YQAIM2DGZmAIM-BQCDeQoACRQAgzcTcmVtb3YAgWIIZnJvbSBsb2NhbACCPAg&s=default)

Source file:

Open: https://www.websequencediagrams.com/#, Source code is:

```text
title Presencejs v2 - 20221009

actor Alice
actor Bob
Alice->Bob: Bob is online, Alice join this channel, sync state.

note over Alice, Bob, prscd: Step 1 - Auth
Alice->prscd: WebSocket/WebTransport, with `auth` and `id`
note right of prscd: Auth
prscd->+AuthSvc: AuthSvc
AuthSvc-->-prscd: AuthResult
prscd->Alice: 4xx if auth fail, Upgrade if success

note over Alice, Bob, prscd: Step 2 - Join Channel
Alice->+prscd: `channel_join`
prscd-->-Alice: `channel_join` ACK
Alice->(2)Bob: `peer_online` with stat_obj payload
note right of Bob: add Alice to online users cache
note right of Bob: should respond `peer_state`
Bob->(2)Alice: `peer_state`
note left of Alice: `peer_state` should add Bob to cache

Alice->(2)Bob: `peer_state` at any time
note right of Bob: update Alice stat_obj

note over Alice, Bob, prscd: Step 3 - Broadcast custome data
Alice->(2)Bob: `data`

note over Alice, Bob, prscd: Step 4 - Leave Channel
Alice->Bob: `peer_offline`
prscd-->-Bob: `peer_offline`
note right of Bob: remove Alice from local cache
```

