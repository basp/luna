namespace Luna

[<AutoOpen>]
module Common =
    type ResultBuilder () =
        member _.Bind(r, f) =
            match r with
            | Ok x -> f x
            | Error e -> Error e
            
        member _.Return x = Ok x
        
        member _.ReturnFrom(r: Result<_, _>) = r

    let result = ResultBuilder()    
