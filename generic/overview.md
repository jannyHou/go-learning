# Overview

Generic types.

- Types can be `any` or `comparable` or any specific type.

    ```go
    type Pair[T fmt.Stringer] struct {
    Val1 T
    Val2 T
    }
    ```
- Create intrface with generic types.

    ```go
    type Differ[T any] interface {
    fmt.Stringer
    Diff(T) float64
    }
    ```