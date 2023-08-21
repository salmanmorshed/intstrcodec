# intstrcodec

`intstrcodec` provides a standalone codec that encodes/decodes integers to/from strings using a custom alphabet and bit-shuffling approach. It's a Golang port of the encoder used in the popular `short_url` Python package. Explanation of the original algorithm can be found here: http://code.activestate.com/recipes/576918/


## Usage
- Install:
```bash
go get github.com/salmanmorshed/intstrcodec
```

- Example:
```go
package main

import (
	"fmt"
	"github.com/salmanmorshed/intstrcodec"
)

func main() {
	alphabet := "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	blockSize := 24

	codec, err := CreateCodec(alphabet, blockSize)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialize codec:", err)
		os.Exit(1)
	}

	original := 123
	encoded := cc.IntToStr(original)
	decoded := cc.StrToInt(encoded)

	fmt.Printf("Original: %d, Encoded: %s, Decoded: %d\n", original, encoded, decoded)
}
```

## License
This project is licensed under the GNU General Public License v2 (GPL-2.0), which is a copyleft open-source license. This means you are free to use, modify, and distribute the software for any purpose, including commercial purposes. However, if you modify and distribute the software, you must make your modified source code available under the same GPL-2.0 license.
