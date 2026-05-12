
```
broker, err := custombroker.NewCBroker(cfg.BucketLog, cfg.maxClientCount, cfg.multCachesMemory)
if err != nil {
    panic(err)
}
```
---
