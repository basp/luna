# luna
This is a Go implementation of the ray tracer described in the fantastic book [The Ray Tracer Challenge](http://raytracerchallenge.com/) by Jamis Buck.

This project has no explict goal apart from me getting more familiar with Go and doing a cool project in the process. This project would have taken a lot longer was it not for the awesome folks at  [go-gl](https://github.com/go-gl/mathgl) who made the `mathgl/mgl64` package which forms the foundation for this ray tracer.

## primitives
In the **Luna** domain there are three primitives which are super important: points, vectors and colors.

### points
Points are `Vec4` values that have their `W` component set to `1`. It is recommended to always use the `luna.Point` function to create a position value.

### vectors
Vectors are `Vec4` values that have their `W` component set to `0`. It is recommended to always use the `luna.Vector` function to create a direction value.

### colors
Colors are represented by `Vec3` values. They have no *alpha* component and they do not clamp to `[0..1]`. This means that upon rendering the client is responsible for clamping these color values into a valid color range as well as providing a fitting *alpha* value.

## transformations
A Luna `Transform` is a cached value consisting of three 4x4 matrices. The transformation matrix itself, the inverse of that transform matrix and the transposed inverse of the transform matrix.

The `NewImplicitTransform` function takes a `Mat4x4` argument and calculates the inverse `inv` and the inverse-transpose `invt` matrices. 

```
// creating a transform implicitly
m := luna.RotateY(PiOver4)
t := luna.NewImplicitTransform(m)
```

The `NewExplicitTransform` function takes an explicit inverse matrix `inv` and the inverse-transpose `invt` is calculated using the explicit inverse.

```
// creating a transform explicitly
m := luna.RotateY(PiOver4)
// for example purposes, usually uses specialized `Inv` instead
inv := m.Inv() 
t := luna.NewExplicitTransform(m, inv)
```

