# Koding på norsk

Denne kodebrønnen inneholder det norske programmeringsspråket Pytonskript. Språket er basert på _Monkey_ som er beskreveti boken [Writing an Interpreter in Go](https://interpreterbook.com/).

Pytonskript er laget for å vise hvordan programmeringsspråk kan se ut på norsk.

## Installering

```go
go install github.com/solbero/pytonskriptskript@latest
```

## Bruk

### Interaktivt
```bash
$ go run main.go
>> la svar = 6 * 7;
>> svar;
42
```

### Kjøring fra fil
```bash
$ go run main.go ./examples/variabler.pytonskript
HelloStavanger!
```

## Lisens

MIT License
