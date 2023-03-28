# Oga
![oga at the top](docs/oga-at-the-top.png)

**Oga** is an esoteric language inspired by [this](https://youtu.be/h1xva9WC4G0?t=188). This is my way of learning what it takes to implement a language. Enjoy!

## How far world
Every oga script needs a `mehn` function:

![mehn function](docs/mehn-function.png)

## Installing Oga
Oga is still in development, but you can download the latest experimental release on the [releases page](https://github.com/tobiade/oga/releases).

Alternatively, if you have Go installed, you can install the latest version of oga:
```sh
go install github.com/tobiade/oga@latest
```

## Using Oga
Create your script e.g `my_script.oga`:
![sample script](docs/ex1.png)

To run your oga script:
```sh
oga abeg run my_script.oga
```
However, there is no guarantee that oga will actually run your code. If you want to ensure your code is run, you have to beg by using the `pls` flag:
```sh
oga abeg run my_script.oga --pls
```

## Testing?
Testimony.