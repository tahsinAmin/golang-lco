## Production ready file (Before you write any code, you have to write)

'go mod init hello'

## Run

'go run <file_name>.go'

## Packages

### fmt

#### Printf

INPUT
'''
var username string = "Hitesh"
fmt.Printf("Variables is a of type: %T \n", username)
var isLoggedIn bool = false
fmt.Printf("Variables is a of type: %T \n", isLoggedIn)
'''
OUTPUT:
'''
Variables is a of type: string
Variables is a of type: bool
'''


## Numeric types
An integer, floating-point, or complex type represents the set of integer, floating-point, or complex values, respectively. They are collectively called numeric types. The predeclared architecture-independent numeric types are:

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

Source: https://go.dev/ref/spec#Lexical_elements

## Setting GOOS
$env:GOOS = "linux"
go build
