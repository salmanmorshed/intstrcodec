# intstrcodec

`intstrcodec` encodes positive integers into *seemingly* random strings. These strings can be decoded back into
the original integer values. The intended use case of the codec is to generate URL keys that are derived from 
numeric values (serial primary keys of a database table, for example) and have them *look* like random strings.

The transformation process has some important properties:

1. The encode and decode process is deterministic.
2. There are no collisions between encoded string values.
3. Strings derived from sequential numbers exhibit no common tail patterns.

Thanks to these properties, it is possible to keep the internal numeric values secret and use the string values
as public facing keys, given that the alphabet and blockSize are not exposed.

It's important to note that this library does not employ any cryptographically sane technique and not a good choice
for any application that requires confidentiality. It is essentially an obfuscation technique.

## Usage

### Install
```bash
go get github.com/salmanmorshed/intstrcodec
```

### Basic Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/salmanmorshed/intstrcodec"
)

func main() {
	alphabet := "bkus2y8ng9dch5xam3t6r7pqe4zfwjv"
	blockSize := 24

	codec, err := intstrcodec.New(alphabet, blockSize)
	if err != nil {
		log.Fatalln("failed to initialize codec:", err)
	}

	original := 123
	encoded := codec.Encode(original)
	decoded := codec.Decode(encoded)

	fmt.Printf("original=%d, encoded=%s, decoded=%d\n", original, encoded, decoded)
}
```

## Acknowledgement
The transformation technique this codec employs is directly taken from this python recipe by Michael Fogleman: 
http://code.activestate.com/recipes/576918/

## License
This project is licensed under the [MIT License](https://github.com/git/git-scm.com/blob/main/MIT-LICENSE.txt). 
The MIT License is a permissive open-source license that allows you to freely use, modify, and distribute this 
software for both commercial and non-commercial purposes, provided you include the original copyright notice and 
disclaimer. Feel free to explore, contribute, and build upon this project with confidence under the terms of the 
MIT License.