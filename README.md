# go-sss
Shamir's Secret Sharing (SSS) in Go

## Usage

```
go get github.com/Tverous/go-sss
```

## Example

```
package main

import "github.com/Tverous/go-sss/sss"

func main()  {
	n := 3
	m := 5
	secret := 123456789
	sss.GenerateShares(n, m, secret)
}
```
