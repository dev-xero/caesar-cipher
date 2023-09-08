<h1><img src="./assets/gopher.png" width="32px"/> Caesar Cipher</h1>

<img src="./assets/ui.webp" width="720px" />

Simple Go implementation of a Caesar Cipher Algorithm.  

The Caesar Cipher is a fundamental cryptographic technique in which letters within a word undergo precise shifts forward or backward by a specified amount, resulting in the creation of a new word with altered letter positions. It proves particularly valuable for safeguarding uncomplicated passwords.

## Running the script
> [!IMPORTANT]
> You need the [Go compiler]("https://go.dev/dl/") installed to compile the script.  

Compiling:  
```go
go fmt ./...
go build -o .\bin\ ./...
```
<details>
  <summary>Changing the compile location</summary>
  <p>You can change the build location to somewhere else using the -o flag</p>
</details>

Running:
```go
.\bin\caesar-cipher
```


> [!NOTE]
> Better still, you can batch run those commands using make
> ```make
> make
> ```

## Ciphering

Say, we decide to encrypt the word "password" forward by eight shifts, using the Caesar Cipher we'd have:  
-  Original: `p` `a` `s` `s` `w` `o` `r` `d` 
- Shifted (by 8): `x` `i` `a` `a` `e` `w` `z` `l`
