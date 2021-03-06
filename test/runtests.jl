using Luna
using Test

@testset "Luna" begin
    @testset "Subtracting two points" begin
        p1 = point(3, 2, 1)
        p2 = point(5, 6, 7)
        @test p1 - p2 == vector(-2, -4, -6)
    end

    @testset "Subtracting a vector from a point" begin
        p = point(3, 2, 1)
        v = vector(5, 6, 7)
        @test p - v == point(-2, -4, -6)
    end

    @testset "Subtracting two vectors" begin
        v1 = vector(3, 2, 1)
        v2 = vector(5, 6, 7)
        @test v1 - v2 == vector(-2, -4, -6)
    end

    @testset "Multiplying by a translation matrix" begin
        t = Luna._translate(5, -3, 2) |> Transform
        p = point(-3, 4, 5)
        @test matrix(t) * p == point(2, 1, 7)
    end

    @testset "Individual transformations are applied in sequence" begin
        A = Luna._rotx(π / 2)
        B = Luna._scale(5, 5, 5)
        C = Luna._translate(10, 5, 7)
        p1 = point(1, 0, 1)
        p2 = A * p1
        p3 = B * p2
        p4 = C * p3
        @test p4 == point(15, 0, 7)
    end

    @testset "Chained transformations are applied in reverse order" begin
        A = Luna._rotx(π / 2)
        B = Luna._scale(5, 5, 5)
        C = Luna._translate(10, 5, 7)
        T = C * B * A
        p = point(1, 0, 1)
        @test T * p == point(15, 0, 7)
    end

    @testset "Fluent API transformations can be chained in order" begin
        T = eye() |> rotx(π / 2) |> scale(5, 5, 5) |> translate(10, 5, 7)
        p = point(1, 0, 1)
        @test T * p == point(15, 0, 7)
    end
end
