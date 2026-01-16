module rec Luna.Interpreter

type Value =
    | Literal of AST.Literal
    | Quote of AST.Term

type Stack = Value list

type Builtin = Stack -> Stack

type Env = Map<string, Builtin>

let dup : Stack -> Stack = function
    | x :: xs -> x :: x :: xs
    | _ -> failwith "dup: stack underflow"

let swap : Stack -> Stack = function
    | x :: y :: xs -> y :: x :: xs
    | _ -> failwith "swap: stack underflow"

let ``+`` : Stack -> Stack = function
    | Value.Literal (AST.Int x) :: Value.Literal (AST.Int y) :: xs ->
        Value.Literal (AST.Int (y + x)) :: xs
    | _ -> failwith "+: expected two integers"

let i : Stack -> Stack = function
    | Value.Quote q :: xs -> evalTerm builtins xs q
    | _ -> failwith "i: expected quotation"

let builtins : Env =
    Map.empty
    |> Map.add "+" ``+``
    |> Map.add "dup" dup
    |> Map.add "swap" swap
    |> Map.add "i" i

let evalTerm (env: Env) (stack: Stack) (term: AST.Term): Stack =
    List.fold (evalFactor env) stack term    

let evalFactor (env: Env) (stack: Stack) (factor: AST.Factor): Stack =
    match factor with
    | AST.Factor.Literal lit ->
        Value.Literal lit :: stack
    | AST.Factor.Quote term ->
        Value.Quote term :: stack
    | AST.Factor.Symbol name ->
        match env.TryFind name with
        | Some fn -> fn stack
        | None -> failwithf $"Unknown symbol: %s{name}"
