# Ojichat With WASM

Ojichat will running on https://stdll00.github.io

build
```
sh build.sh
```

host 
```
python3 -m http.server
# or 
goexec 'http.ListenAndServe(`:8000`, http.FileServer(http.Dir(`.`)))'

```