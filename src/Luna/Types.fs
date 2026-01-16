module Luna.Types

type Type =
    | Int
    | Bool
    | Quote
    
type Stack = Type list

type StackEffect =
    {
        Pop : Stack
        Push : Stack
    }       
    
type EffectEnv = Map<string, StackEffect>

let applyEffect (stack: Stack) (effect: StackEffect) : Result<Stack, string> =
    let rec consume pops stack =
        match pops, stack with
        | [], _ -> Ok stack
        | p :: ps, s :: ss when p = s -> consume ps ss
        | p :: _, [] ->
            Error $"Stack underflow, expected %A{p}"
        | p :: _, s :: _ ->
            Error $"Type mismatch, expected %A{p} but got %A{s}"
    result {
        let! rest = consume effect.Pop stack
        return effect.Push @ rest
    }

let plusEffect =
    {
        Pop = [Int; Int]
        Push = [Int]
    }
    
let dupIntEffect =
    {
        Pop = [Int ]
        Push = [Int; Int]
    }

let swapIntEffect =
    {
        Pop = [Int; Int]
        Push = [Int; Int]
    }
    
let iEffect =
    {
        Pop = [Quote]
        Push = []
    }

let effects : EffectEnv =
    Map.empty
    |> Map.add "+" plusEffect
    |> Map.add "dup" dupIntEffect
    |> Map.add "swap" swapIntEffect
    |> Map.add "i" iEffect
