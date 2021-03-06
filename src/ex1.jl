using Images
using Luna

function test()
    function closest!(xs)
        filter!(x -> x.t > 0, xs)
        if isempty(xs)
            return false, nothing
        end
        return true, first(xs).t
    end

    function shade(obj::Object{:sphere}, r)
        xs = intersect(obj, r)
        ok, t = closest!(xs)
        if ok
            worldpoint = r(t)
            n = normalat(obj, worldpoint)
            return 0.5 * RGB(n.x + 1, n.y + 1, -n.z + 1)
        end
        uv = normalize(direction(r))
        t = 0.5 * (uv.y + 1.0)
        lerp(RGB(1.0, 1, 1), RGB(0.5, 0.7, 1.0), t)
    end

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
    lowerleftcorner = origin - horizontal / 2 - vertical / 2 + focal_aspect

    img = fill(RGB(0, 0, 0), rows, cols)
    Threads.@threads for j in reverse(0:rows-1)
        for i = 0:cols-1
            u = i / (cols - 1)
            v = j / (rows - 1)
            direction = lowerleftcorner + u * horizontal + v * vertical - origin
            ray = Ray(origin, direction)
            color = shade(obj, ray)
            img[rows-j, i+1] = color
        end
    end

    save("test.png", img)
end
