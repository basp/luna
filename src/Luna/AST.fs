module rec Luna.AST

type Literal =
    | Int of int
    | Bool of bool

type Factor =
    | Literal of Literal
    | Symbol of string
    | Quote of Term

type Term = Factor list
