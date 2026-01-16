module Luna.Types

type Type =
    | Int
    | Bool
    | Quote
    | Var of string
    
type Stack = Type list

type Subst = Map<string, Type>

let rec applySubstType (s: Subst) (t: Type) : Type =
    match t with
    | Var a ->
        match s.TryFind a with
        | Some t' -> applySubstType s t'
        | None -> t
    | Int -> t
    | Bool -> t
    | Quote -> t

