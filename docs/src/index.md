# Luna

## Types

```@docs
Vector4{T}
Matrix4x4{T}
Ray{T}
Transform{T}
Object{S,T}
Intersection{S,T}
```

## Functions

```@docs
lerp(a, b, t)
point(x, y, z)
vector(x, y, z)
point(u)
vector(u)
ispoint(u)
isvector(u)
cross(u::Vector4, v::Vector4)
origin(r)
direction(r)
*(m::Matrix4x4, r::Ray)
*(r::Ray, m::Matrix4x4)
translate(dx, dy, dz)
scale(sx, sy, sz)
rotx(rad)
roty(rad)
rotz(rad)
shear(xy, xz, yx, yz, zx, zy)
matrix(t)
inv(t::Transform)
eltype(t::Transform)
transform(obj)
sphere(t)
cube(t)
```