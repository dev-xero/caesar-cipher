<h1><img src="./assets/gopher.png" width="32px"/> Caesar Cipher</h1>

Simple Go implementation of a Caesar Cipher Algorithm.  

The Caesar Cipher is a fundamental cryptographic technique in which letters within a word undergo precise shifts forward or backward by a specified amount, resulting in the creation of a new word with altered letter positions. It proves particularly valuable for safeguarding uncomplicated passwords.

## Compiling
> [!IMPORTANT]
> You need the [Go compiler](https://go.dev/dl/) installed to compile the script.  

Compiling:  
```go
go fmt ./...
go build -o .\bin\ ./...
```

> [!NOTE]
> You can batch execute those commands using make (you need make installed)
> ```bash
> make
> ```

## Using

Say, we decide to encrypt the word "password" forward by eight shifts, we just simply supply the word and the shift number:  
```go
.\bin\caesar-cipher -e password 8
```

-  Original: `p` `a` `s` `s` `w` `o` `r` `d` 
- Shifted (by 8): `x` `i` `a` `a` `e` `w` `z` `l`


The `-e` flag exports the encrypted word in a text file stored in a directory called "exported".
