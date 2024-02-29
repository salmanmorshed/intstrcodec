# intstrcodec

`intstrcodec` provides a standalone codec that encodes/decodes integers to/from strings using a custom alphabet and 
bit-shuffling approach. It's a Golang port of the encoder used in the popular `short_url` Python package. Explanation 
of the original algorithm can be found here: http://code.activestate.com/recipes/576918/


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
		log.Fatal("failed to initialize codec:", err)
	}

	original := 123
	encoded := codec.IntToStr(original)
	decoded := codec.StrToInt(encoded)

	fmt.Printf("Original: %d, Encoded: %s, Decoded: %d\n", original, encoded, decoded)
}
```

### Custom Options
```go
package main

import "github.com/salmanmorshed/intstrcodec"

func main() {
	_, _ = intstrcodec.New(
		"bkus2y8ng9dch5xam3t6r7pqe4zfwjv", 24,

		// Set the minimum length of output string to 5
		intstrcodec.WithMinLength(5),

		// Use a custom implementation of integer power calculation
		intstrcodec.WithIntPowerFn(intstrcodec.CustomIntPower),
	)
}
```

#### Note on `WithIntPowerFn(CustomIntPower)`
The native golang math.Pow function introduces floating point error in the calculations which causes the codec to 
break for inputs beyond 2^55. This implementation manually calculates the power value using int only which allows 
the codec to successfully decode to the max value of int64. Measured performance difference is approximately 12%.

## License
This project is licensed under the [MIT License](https://github.com/git/git-scm.com/blob/main/MIT-LICENSE.txt). 
The MIT License is a permissive open-source license that allows you to freely use, modify, and distribute this 
software for both commercial and non-commercial purposes, provided you include the original copyright notice and 
disclaimer. Feel free to explore, contribute, and build upon this project with confidence under the terms of the 
MIT License.