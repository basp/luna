module Luna.Inference

open System.Collections
open Types
open AST
open StackEffects
    
let typeOfLiteral = function
    | Literal.Int _ -> Type.Int
    | Literal.Bool _ -> Type.Bool
    
let rec infer (env: EffectEnv) (init: Stack) (term: Term)
    : Result<Stack, string> =
        
    let folder (stackRes: Result<Stack, string>) (factor: Factor) =
        result {
            let! stack = stackRes
            match factor with
            | Factor.Literal lit ->
                return (typeOfLiteral lit) :: stack
            | Factor.Quote _ ->
                // Recursively check the quotation here.
                return Type.Quote :: stack
            | Factor.Symbol name ->
                match env.TryFind name with
                | None ->
                    return! Error $"Unknown symbol: %s{name}"
                | Some effect ->
                    return! applyEffect stack effect
        }
        
    List.fold folder (Ok init) term
    