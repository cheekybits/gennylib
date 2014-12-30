gennylib
========

Library of generic stuff for [Genny](https://github.com/metabition/genny).


## Contributing

If you'd like to contribute a generic template for use with `genny get`, please fork and send us a pull request! We're trying to build up a large collection of templates for anyone to use.

## Maps

#### Concurrent map

```
wget -q -O - "https://github.com/cheekybits/gennylib/raw/master/maps/concurrentmap.go" | genny gen "KeyType=BUILTINS ValueType=BUILTINS" >> conmaps.go
```
