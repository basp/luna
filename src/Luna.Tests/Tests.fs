module Tests

open Luna.AST
open Luna.Interpreter
open Xunit

let run term =
    evalTerm builtins [] term

let litVal v = Value.Literal v

let intVal i = litVal (Int i)

let boolVal b = litVal (Bool b)

let quote t = Value.Quote t

[<Fact>]
let ``push literal integer`` () =
    let term : Term = [
        Factor.Literal (Int 3)
        Factor.Literal (Int 4)
    ]    
    let result = run term
    let expected = [intVal 4; intVal 3]
    Assert.Equal<Value list>(expected, result)

[<Fact>]
let ``push literal boolean`` () =
    let term : Term = [
        Factor.Literal (Bool true)
        Factor.Literal (Bool false)
    ]
    let result = run term
    let expected = [boolVal false; boolVal true]
    Assert.Equal<Value list>(expected, result)
    
[<Fact>]
let ``push various literals`` () =
    let term : Term = [
        Factor.Literal (Int 3)
        Factor.Literal (Bool true)
    ]
    let result = run term
    let expected = [litVal (Bool true); litVal (Int 3)]
    Assert.Equal<Value list>(expected, result)
    
[<Fact>]
let ``integer addition`` () =
    let term = [
        Factor.Literal (Int 2)
        Factor.Literal (Int 3)
        Factor.Symbol "+"
    ]
    let result = run term
    Assert.Equal<Value list>([intVal 5], result)    
    
[<Fact>]
let ``duplicate top value`` () =
    let term : Term = [
        Factor.Literal (Int 10)
        Factor.Symbol "dup"
    ]
    let result = run term
    let expected = [litVal (Int 10); litVal (Int 10)]
    Assert.Equal<Value list>(expected, result)
    
[<Fact>]
let ``swap top two values`` () =
    let term = [
        Factor.Literal (Int 2)
        Factor.Literal (Int 3)
        Factor.Symbol "swap"
    ]
    let result = run term
    let expected = [litVal (Int 2); litVal (Int 3)]
    Assert.Equal<Value list>(expected, result)
    
[<Fact>]
let ``push quotation`` () =
    let inner = [
        Factor.Literal (Int 2)
        Factor.Literal (Int 3)
    ]
    let term = [ Factor.Quote inner ]
    let expected = [quote inner]
    let result = run term
    Assert.Equal<Value list>(expected, result)

[<Fact>]
let ``execute quotation`` () =
    let q = [
        Factor.Literal (Int 4)
        Factor.Symbol "+"
    ]
    let term = [
        Factor.Literal (Int 3)
        Factor.Quote q
        Factor.Symbol "i"
    ]
    let result = run term
    Assert.Equal<Value list>([intVal 7], result)

[<Fact>]
let ``addition stack effect`` () =
    let term = [
        Factor.Literal (Int 2)
        Factor.Literal (Int 3)
        Factor.Symbol "+"
    ]
    let effects = Luna.StackEffects.effects
    let res = Luna.Inference.infer effects [] term
    match res with
    | Ok s -> Assert.Equal<Luna.Types.Stack>([Luna.Types.Int], s)
    | Error e -> Assert.Fail(e)
    
[<Fact>]
let ``detect type error`` () =
    let term = [
        Factor.Literal (Int 2)
        Factor.Literal (Bool true)
        Factor.Symbol "+"
    ]
    let res = Luna.Inference.infer Luna.StackEffects.effects [] term
    match res with
    | Ok _ -> Assert.Fail("Expected type error")
    | Error e -> Assert.Contains("Type mismatch", e)
    
[<Fact>]
let ``dup effect contains type variables`` () =
    let eff = Luna.StackEffects.dupEffect
    Assert.Equal<Luna.Types.Stack>([ Luna.Types.Var "a" ], eff.Pop)
    Assert.Equal<Luna.Types.Stack>(
        [ Luna.Types.Var "a"; Luna.Types.Var "a" ], eff.Push)
    
[<Fact>]
let ``swap effect contains type variables`` () =
    let eff = Luna.StackEffects.swapEffect
    Assert.Equal<Luna.Types.Stack>(
        [ Luna.Types.Var "a"; Luna.Types.Var "b" ], eff.Pop)
    Assert.Equal<Luna.Types.Stack>(
        [ Luna.Types.Var "b"; Luna.Types.Var "a" ], eff.Push)
    
[<Fact>]
let ``dup effect uses single type variable`` () =
    let a = Luna.Types.Var "a"
    Assert.Equal<Luna.Types.Stack>([ a ], Luna.StackEffects.dupEffect.Pop)
    Assert.Equal<Luna.Types.Stack>([ a; a ], Luna.StackEffects.dupEffect.Push)
    
[<Fact>]
let ``swap effect uses two distinct type variables`` () =
    let a = Luna.Types.Var "a"
    let b = Luna.Types.Var "b"
    Assert.Equal<Luna.Types.Stack>([ a; b ], Luna.StackEffects.swapEffect.Pop)
    Assert.Equal<Luna.Types.Stack>([ b; a ], Luna.StackEffects.swapEffect.Push)
    
[<Fact>]
let ``plus effect is monomorphic`` () =
    let int = Luna.Types.Int
    Assert.Equal<Luna.Types.Stack>([ int; int ], Luna.StackEffects.plusEffect.Pop)
    Assert.Equal<Luna.Types.Stack>([ int ], Luna.StackEffects.plusEffect.Push)
