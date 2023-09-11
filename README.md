<h1><img src="./assets/gopher.png" width="32px"/> Caesar Cipher</h1>

A CLI tool for encrypting text using the Caesar Cipher, written in Go.  

The Caesar Cipher is a fundamental cryptographic technique in which letters within a word undergo precise shifts forward or backward by a specified amount, resulting in the creation of a new word with altered letter positions. It proves particularly valuable for safeguarding uncomplicated passwords.

## Compiling
> [!IMPORTANT]
> You need the [Go compiler](https://go.dev/dl/) installed to compile the script.  

```go
go fmt ./...
go build -o .\bin\ ./...
```

> [!NOTE]
> You can batch execute those commands using make (you need make installed)
> ```bash
> make
> ```

## Usage

Say, we decide to encrypt the word "password" forward by eight shifts, we simply just pass the word and the shift number as arguments:  
```go
.\bin\caesar-cipher -e password 8
```

### Result

-  Original: `p` `a` `s` `s` `w` `o` `r` `d` 
- Shifted (by 8): `x` `i` `a` `a` `e` `w` `z` `l`


The `-e` flag writes the result to a text file which is stored in a directory named "exported".
