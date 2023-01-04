# PMK-WASM

A WPA2 PMK Generator in WebAssembly written in Go


## Build
```bash
GOOS=js GOARCH=wasm go build -o ./pkm-wasm.wasm
```


## Usage
```html
<!-- https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js --> 
<script src="wasm_exec.js"></script>
<script>
    const goWasm = new Go()
    
    WebAssembly.instantiateStreaming(fetch("pmk-wasm.wasm"), goWasm.importObject).then((result) => {
        goWasm.run(result.instance);
        
        var ssid = "Tp-link";
        var passphrase = "password123";
        var pmk = generateWpa2Pmk(passphrase, ssid);
        
        console.log(pmk); // db40321e8801cfa7300366d5919a5fec6e7ca7ed0d8d79738bd9791e7c46ad32
    });
</script>
```

