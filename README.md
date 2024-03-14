# Simple json serialization/deserialization tests

On my machine (AMD Ryzen Threadripper 3970X 32-Core Processor):

```
$ go build -v . && ./encode-decode
encode-decode
Generating 1000000 samples
*** Starting single threaded test ***
Encoding 1000000 samples took 3.057588639s or 327055.113708 per second
Decoding 1000000 samples took 6.810411931s or 146833.996259 per second
Encoding/Decoding 1000000 samples took 9.86800057s or 101337.651220 per second
*** Starting multi threaded test ***
Multi-threaded test took 33.720688888s
Encoding/Decoding 64000000 samples took 33.720688888s or 1897944.618290 per second
```
