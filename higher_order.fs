module HigherOrder

let numbers1 = [1; 2; 3; 4; 5]
let numbers2 = [6; 7; 8; 9; 10]

let printList values =
    printfn "[%s]" (values |> List.map string |> String.concat "; ")

printfn "List.map - squares:"
[numbers1; numbers2]
|> List.iter (fun numbers ->
    let squares = numbers |> List.map (fun number -> number * number)
    printf "Input: "
    printList numbers
    printf "Output: "
    printList squares)

printfn "\nList.filter - even numbers:"
[numbers1; numbers2]
|> List.iter (fun numbers ->
    let evens = numbers |> List.filter (fun number -> number % 2 = 0)
    printf "Input: "
    printList numbers
    printf "Output: "
    printList evens)

printfn "\nList.fold - sums:"
[numbers1; numbers2]
|> List.iter (fun numbers ->
    let sum = numbers |> List.fold (fun total number -> total + number) 0
    printf "Input: "
    printList numbers
    printfn "Sum: %d" sum)
