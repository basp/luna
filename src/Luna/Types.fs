module Luna.Types

type Type =
    | Int
    | Bool
    | Quote
    | Var of string
    
type Stack = Type list

type Subst = Map<string, Type>

let rec resolve (s: Subst) (t: Type) : Type =
    match t with
    | Var a ->
        match s.TryFind a with
        | Some t' -> resolve s t'
        | None -> t
    | Int -> t
    | Bool -> t
    | Quote -> t

let resolveStack s stack =
    List.map (resolve s) stack

let unify (s: Subst) (t1: Type) (t2: Type) : Result<Subst, string> =
    let t1 = resolve s t1
    let t2 = resolve s t2
    match t1, t2 with
    | Int, Int -> Ok s
    | Bool, Bool -> Ok s
    | Var a, t -> Ok (s.Add(a, t))
    | t, Var a -> Ok (s.Add(a, t))
    | Quote, Quote -> Ok s // placeholder for now
    | _ -> Error $"cannot unify %A{t1} with %A{t2}"