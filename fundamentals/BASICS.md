###### Go is statically-typed, which means all variables must have a defined type at compile-time.
```
var explicit int // Explicitly typed
implicit := 10 // Implicitly typred as an int
```

###### Once declared variables can be assigned values with the = operator

```
count := 1 // Assign initial value
count = 2  // Update/Change to new value

count = false // This throws a compiler error due to assigning a non `int` type
```

###### Functions accepts zero or more parameters

```
// Hello is a public function.
func Hello (name string) string {
    return hi(name)
}

// hi is a private function.
func hi (name string) string {
    return "hi " + name

private = lowercase
public = uppercase
```