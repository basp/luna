module Luna.StackEffects

open Luna.Types

type StackEffect =
    {
        Pop : Stack
        Push : Stack
    }       
    
type EffectEnv = Map<string, StackEffect>

let applySubst (s: Subst) (eff: StackEffect) : StackEffect =
    {
        Pop = eff.Pop |> List.map (applySubst s)
        Push = eff.Push |> List.map (applySubst s)
    }
    
let dupEffect =
    {
        Pop = [Var "a"]
        Push = [Var "a"; Var "a"]
    }
    
let swapEffect =
    {
        Pop = [Var "a"; Var "b"]
        Push = [Var "b"; Var "a"]
    }
    
let iEffect =
    {
        Pop = [Quote]
        Push = []
    }

let plusEffect =
    {
        Pop = [Int; Int]
        Push = [Int]
    }

let effects : EffectEnv =
    Map.empty
    |> Map.add "dup" dupEffect
    |> Map.add "swap" swapEffect
    |> Map.add "+" plusEffect

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