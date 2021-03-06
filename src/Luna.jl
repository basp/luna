module Luna

using Colors
using Images
using LinearAlgebra
using PolynomialRoots
using StaticArrays

export I
export Intersection
export Matrix4x4
export Object
export Ray
export Transform
export Vector4
export cross
export cube
export direction
export dot
export ispoint, isvector
export lerp
export matrix
export norm
export normalat
export normalize
export origin
export point
export reflect
export rotx, roty, rotz
export scale
export shear
export sphere
export transform
export translate
export vector

"""
    Vector4{T} <: FieldVector{4,T}

A 4-vector with elements of type `T`.
"""
struct Vector4{T} <: FieldVector{4,T}
    x::T
    y::T
    z::T
    w::T
end

"""
    Matrix4x4{T} = SMatrix{4,4,T,16} where T<:Real

A 4x4 matrix with elements of type `T`.
"""
const Matrix4x4{T} = SMatrix{4,4,T,16} where T<:Real

struct Ray{T}
    origin::Vector4{T}
    direction::Vector4{T}
end

struct Transform{T}
    matrix::Matrix4x4{T}
    matrix⁻¹::Matrix4x4{T}
end

"""
    Object{S,T}

An object is a thing that can be rendered where `S` is
expected to be a *symbol* tag and `T` is the numeric type
of the object's transformation. 

# Remarks
Objects are always defined in unit space and depend on
their `transform` to move them into world space.

# Examples
```julia-repl
julia> t = Matrix4x4{Float64}(I) |> scale(0.5, 0.75, 0.25) |> Transform
Transform([0.5 0.0 0.0 0.0; 0.0 0.75 0.0 0.0; 0.0 0.0 0.25 0.0; 0.0 0.0 0.0 1.0])

julia> Object{:foo, Float64}(t)
Object{:foo,Float64}(Transform([0.5 0.0 0.0 0.0; 0.0 0.75 0.0 0.0; 0.0 0.0 0.25 0.0;
0.0 0.0 0.0 1.0]))
```
"""
struct Object{S,T}
    transform::Transform{T}
    children::Array{Object{S,T}, 1}     
end

struct Intersection{S,T}
    t::T
    object::Object{S,T}
end

StaticArrays.similar_type(::Type{V}, ::Type{F}, size::Size{(4,)}) where {V<:Vector4, F<:Real} = Vector4

Base.promote_rule(::Type{Vector4{Int64}}, ::Type{Vector4{Rational{Int64}}}) = Vector4{Rational{Int64}}
Base.promote_rule(::Type{Vector4{Rational{Int64}}}, ::Type{Vector4{Int64}}) = Vector4{Rational{Int64}} 
Base.promote_rule(::Type{Vector4{Int64}}, ::Type{Vector4{Float64}}) = Vector4{Float64}
Base.promote_rule(::Type{Vector4{Float64}}, ::Type{Vector4{Int64}}) = Vector4{Float64}

function Base.show(io::IO, u::Vector4)
    if ispoint(u)
        print(io, "point($(u.x), $(u.y), $(u.z))")
    elseif isvector(u)
        print(io, "vector($(u.x), $(u.y), $(u.z))")
    else
        print(io, u)
    end
end

"""
    isless(a::Intersection, b::Intersection)

Return `true` if `a` is closer than `b`.

# Remarks
Intersections are compared using their `t` value.
"""
Base.isless(a::Intersection, b::Intersection) = isless(a.t, b.t)

"""
    lerp(a, b, t)

Linearly interpolates between `a` and `b`.

# Examples
```julia-repl
julia> lerp(0, 1, 0.5)
0.5

julia> lerp(0, 2, 0.5)
1.0

julia> lerp(-2, 2, 0.9)
1.6
```
"""
lerp(a, b, t) = (1 - t) * a + t * b

"""
    point(x, y, z)

Creates a position vector in 3-dimensional space.
"""
point(x, y, z) = Vector4(promote(x, y, z, one(x))...)

"""
    vector(x, y, z)

Creates a direction vector in 3-dimensional space.
"""
vector(x, y, z) = Vector4(promote(x, y, z, zero(x))...)

"""
    point(u)

Converts 4-vector `u` into a position vector.
"""
point(u) = point(u.x, u.y, u.z)

"""
    vector(u)

Converts 4-vector `u` into a direction vector.
"""
vector(u) = vector(u.x, u.y, u.z)

"""
    ispoint(u)

Returns `true` if 4-vector `u` is a position vector.
"""
ispoint(u) = isone(u.w)

"""
    isvector(u)

Returns `true` if 4-vector `u` is a direction vector.
"""
isvector(u) = iszero(u.w)

"""
    cross(u::Vector4, v::Vector4)

Computes the cross product of two 4-vectors. 

# Remarks
The `w` component is ignored and the result is returned as
a new direction vector.

# Examples
```julia-repl
julia> u = vector(1,1,0)
4-element Vector4{Int64} with indices SOneTo(4):
 1
 1
 0
 0

julia> v = vector(0,0,1)
4-element Vector4{Int64} with indices SOneTo(4):
 0
 0
 1
 0
 
julia> cross(u, v)
4-element Vector4{Int64} with indices SOneTo(4):
  1
 -1
  0
  0
```
"""
function LinearAlgebra.cross(u::Vector4, v::Vector4)
    x = u.y * v.z - u.z * v.y
    y = u.z * v.x - u.x * v.z
    z = u.x * v.y - u.y * v.x
    vector(x, y, z)
end

"""
    reflect(u, n)

Reflects vector `u` over vector `n`.
"""
reflect(u, n) = u - 2n * dot(u, n)

"""
    Ray(o, d)

Creates a ray with origin `o` and direction `d`.
"""
Ray(o, d) = Ray(promote(o, d)...)

"""
    origin(r)

Returns the origin position vector for ray `r`.
"""
origin(r) = r.origin

"""
    direction(r)

Returns the direction vector for ray `r`.
"""
direction(r) = r.direction

"""
    (r::Ray)(t)

Calculates a position vector at `t` along ray `r`.

# Examples
```julia-repl
julia> r(0)
4-element Vector4{Int64} with indices SOneTo(4):
 0
 0
 0
 1

julia> r(1)
4-element Vector4{Int64} with indices SOneTo(4):
 0
 0
 1
 1

julia> r(0.5)
4-element StaticArrays.SArray{Tuple{4},Float64,1,4} with indices SOneTo(4):
 0.0
 0.0
 0.5
 1.0
```
"""
(r::Ray)(t) = origin(r) + t * direction(r)

"""
    *(m, r::Ray)
    *(r::Ray, m)

Returns a new ray with `m` applied to origin and direction of `r`.

# Examples
```
julia> m = Matrix4x4{Float64}(I) |> scale(0.5, 0.75, 0.3)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 0.5  0.0   0.0  0.0
 0.0  0.75  0.0  0.0
 0.0  0.0   0.3  0.0
 0.0  0.0   0.0  1.0

julia> r = Ray(point(0,0,0), vector(0,0,1))
Ray{Int64}(point(0, 0, 0), vector(0, 0, 1))

julia> r * m
Ray{Float64}(point(0.0, 0.0, 0.0), vector(0.0, 0.0, 0.3))

julia> m * r
Ray{Float64}(point(0.0, 0.0, 0.0), vector(0.0, 0.0, 0.3))
```
"""
Base.:*(m, r::Ray) = Ray(m * origin(r), m * direction(r))
Base.:*(r::Ray, m) = m * r

"""
    translate(dx, dy, dz)

Creates a translation operation.

# Examples
```
julia> Matrix4x4{Float64}(I) |> translate(1, -2, 3)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 1.0  0.0  0.0   1.0
 0.0  1.0  0.0  -2.0
 0.0  0.0  1.0   3.0
 0.0  0.0  0.0   1.0
```
"""
translate(dx, dy, dz) = m -> m * _translate(dx, dy, dz)

"""
    scale(sx, sy, sz)

Creates a scaling operation.

# Examples
```julia-repl
julia> Matrix4x4{Float64}(I) |> scale(0.5, 0.5, 0.75)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 0.5  0.0  0.0   0.0
 0.0  0.5  0.0   0.0
 0.0  0.0  0.75  0.0
 0.0  0.0  0.0   1.0
```
"""
scale(sx, sy, sz) = m -> m * _scale(sx, sy, sz)

"""
    rotx(rad)

Creates a rotation operation along the x-axis.

# Examples
```julia-repl
julia> Matrix4x4{Float64}(I) |> rotx(π)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 1.0   0.0           0.0          0.0
 0.0  -1.0          -1.22465e-16  0.0
 0.0   1.22465e-16  -1.0          0.0
 0.0   0.0           0.0          1.0

 julia> Matrix4x4{Float64}(I) |> rotx(π/5)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 1.0  0.0        0.0       0.0
 0.0  0.809017  -0.587785  0.0
 0.0  0.587785   0.809017  0.0
 0.0  0.0        0.0       1.0
```
"""
rotx(rad) = m -> m * _rotx(rad)

"""
    roty(rad)

Creates a rotation operation along the y-axis.

See also [`rotx`](@ref) for examples. 
"""
roty(rad) = m -> m * _roty(rad)

"""
    rotz(rad)

Creates a rotation operation along the z-axis.

See also [`rotx`](@ref) for examples. 
"""
rotz(rad) = m -> m * _rotz(rad)

"""
    shear(xy, xz, yx, yz, zx, zy)

Creates a shearing operation.

# Examples
```julia-repl
julia> Matrix4x4{Float64}(I) |> shear(1, 2, 3, -1, -2, -3)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
  1.0   1.0   2.0  0.0
  3.0   1.0  -1.0  0.0
 -2.0  -3.0   1.0  0.0
  0.0   0.0   0.0  1.0
```
"""
shear(xy, xz, yx, yz, zx, zy) = m -> m * _shear(xy, xz, yx, yz, zx, zy)

"""
    Transform(m)

Creates a transform for matrix `m`.

# Remarks
A `Transform` is basically just a memoized inversion matrix so that calling
`inv(t::Transform)` returns the cached inversion instead of recalculating it. 
This implies `inv(t) == inv(matrix(t))`.

# Examples
```julia-repl
julia> m = Matrix4x4{Float64}(I) |> scale(0.5, 0.25, 0.75)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 0.5  0.0   0.0   0.0
 0.0  0.25  0.0   0.0
 0.0  0.0   0.75  0.0
 0.0  0.0   0.0   1.0

 julia> m |> Transform
Transform([0.5 0.0 0.0 0.0; 0.0 0.25 0.0 0.0; 0.0 0.0 0.75 0.0; 0.0 0.0 0.0 1.0])

julia> inv(t) == inv(matrix(t))
true
```
"""
Transform(m) = Transform(promote(m, inv(m))...)

function Base.show(io::IO, t::Transform)
    print(io, "Transform($(matrix(t)))")
end

"""
    matrix(t)

Returns the transform matrix associated with transformation `t`.

# Examples
```julia-repl
julia> t = Matrix4x4{Float64}(I) |> scale(0.5, π/5, -1) |> Transform
Transform([0.5 0.0 0.0 0.0; 0.0 0.6283185307179586 0.0 0.0; 0.0 0.0 -1.0 0.0; 0.0 0.0 0.0 1.0])

julia> matrix(t)
4×4 StaticArrays.SArray{Tuple{4,4},Float64,2,16} with indices SOneTo(4)×SOneTo(4):
 0.5  0.0        0.0  0.0
 0.0  0.628319   0.0  0.0
 0.0  0.0       -1.0  0.0
 0.0  0.0        0.0  1.0
```
"""
matrix(t) = t.matrix

"""
    inv(t::Transform)

Returns the inverse transform matrix.
"""
Base.inv(t::Transform) = t.matrix⁻¹

"""
    eltype(::Transform{T})

Returns the element type of the matrices wrapped by a transform.

# Remarks
Transforms depend on the `inv` method for matrices which creates 
matrices with element type `Float64` or `Float32` depending on 
your system architecture. This means that in practice this method 
will always return `Float64` (or `Float32`).
"""
Base.eltype(::Transform{T}) where T = T

"""
    transform(obj)

Returns the transformation associated with object `obj`.
"""
transform(obj) = obj.transform

"""
    sphere(t)

Returns a unit sphere with transformation `t`.
"""
sphere(t) = Object{:sphere,eltype(t)}(t, [])

"""
    cube(t)

Returns a unit cube with transformation `t`.
"""
cube(t) = Object{:cube,eltype(t)}(t, [])

"""
    group(t, children)

Returns a group with transformation `t` and `children`.
"""
group(t, children) = Objectr{:group,eltype(t)}(t, children)

"""
    Intersection(t, obj::Object{S,T}) where {S,T}

Creates a new intersection.
"""
Intersection(t, obj::Object{S,T}) where {S,T} = Intersection(convert(T, t), obj)

"""
    intersect(obj::Object{:sphere,T}, ray) where T

Returns the intersections (if any) between a sphere and ray.
"""
function Base.intersect(obj::Object{:sphere,T}, ray) where T
    ray *= inv(transform(obj))
    oc = origin(ray) - point(0, 0, 0)
    a = dot(direction(ray), direction(ray))
    b = 2 * dot(direction(ray), oc)
    c = dot(oc, oc) - 1
    d = b * b - 4 * a * c
    if iszero(d)
        t = -b / 2a
        return [Intersection(t, obj)]
    elseif d > 0
        sqrtd = √d
        t1 = (-b - sqrtd) / 2a
        t2 = (-b + sqrtd) / 2a
        return [Intersection(t1, obj), Intersection(t2, obj)]
    end
    return empty([], Intersection{:sphere,T})
end

"""
    normalat(obj::Object{:sphere,T}, worldpoint)

Returns the world normal of `obj` at `worldpoint`.
"""
function normalat(obj::Object{:sphere}, worldpoint)
    minv = inv(transform(obj))
    objectpoint = minv * worldpoint
    objectnormal = objectpoint - point(0, 0, 0)
    worldnormal = transpose(minv) * objectnormal
    return normalize(vector(worldnormal))
end

_translate(x, y, z) =
    @SMatrix[ 1 0 0 x ;
              0 1 0 y ;
              0 0 1 z ;
              0 0 0 1 ]

_scale(x, y, z) =
    @SMatrix[ x 0 0 0 ;
              0 y 0 0 ;
              0 0 z 0 ;
              0 0 0 1 ]

_rotx(r) =
    @SMatrix[ 1      0       0 0 ;
              0 cos(r) -sin(r) 0 ;
              0 sin(r)  cos(r) 0 ;
              0      0       0 1 ]

_roty(r) =
    @SMatrix[  cos(r) 0 sin(r) 0 ;
                    0 1      0 0 ;
              -sin(r) 0 cos(r) 0 ;
                    0 0      0 1 ]

_rotz(r) =
    @SMatrix[ cos(r) -sin(r) 0 0 ;
              sin(r)  cos(r) 0 0 ;
                   0       0 1 0 ;
                   0       0 0 1 ]

_shear(xy, xz, yx, yz, zx, zy) =
    @SMatrix[  1 xy xz 0 ;
              yx  1 yz 0 ;
              zx zy  1 0 ;
               0  0  0 1 ]

function trygethit(xs)
    filter!(x -> x.t > 0, xs)
    if isempty(xs)
        return false, nothing
    end
    return true, first(xs).t
end

function shade(obj::Object{:sphere}, r)
    xs = intersect(obj, r)
    ok, t = trygethit(xs)
    if ok
        worldpoint = r(t)
        n = normalat(obj, worldpoint)
        return 0.5 * RGB(n.x+1, n.y+1, -n.z+1)
    end
    uv = normalize(direction(r))
    t = 0.5 * (uv.y + 1.0)
    lerp(RGB(1.0, 1, 1), RGB(0.5, 0.7, 1.0), t)
end

function test()
    obj = I |> scale(0.5, 0.5, 0.5) |> Transform |> sphere
    
    ar = 16.0 / 9.0
    cols = 1920
    rows = floor(Int, cols / ar)
    
    vpheight = 2.0
    vpwidth = ar * vpheight
    focal_lenght = 1.0

    origin = point(0, 0, -1)
    horizontal = vector(vpwidth, 0, 0)
    vertical = vector(0, vpheight, 0)

    focal_aspect = vector(0, 0, focal_lenght)
    lowerleftcorner = origin - horizontal/2 - vertical/2 + focal_aspect

    data = fill(RGB(0,0,0), rows, cols)
    Threads.@threads for j in reverse(0:rows - 1)
        for i in 0:cols - 1
            u = i / (cols - 1)
            v = j / (rows - 1)
            direction = lowerleftcorner + u*horizontal + v*vertical - origin
            ray = Ray(origin, direction)
            color = shade(obj, ray)
            data[rows - j, i + 1] = color
        end
    end

    save("test.png", data)
end

end # module
