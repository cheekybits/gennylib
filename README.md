gennylib
========

Library of generic stuff for [Genny](https://github.com/metabition/genny).


## Maps

#### Concurrent map

```
wget -q -O - "https://github.com/metabition/gennylib/raw/master/maps/concurrentmap.go" | genny gen "KeyType=BUILTINS ValueType=BUILTINS" >> conmaps.go
```
