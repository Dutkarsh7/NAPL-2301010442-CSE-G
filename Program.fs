module Program

open System

let readMark subjectNumber =
    printf "Enter marks for subject %d: " subjectNumber
    Double.Parse(Console.ReadLine())

let marks = [1..4] |> List.map readMark
let percentage = List.sum marks / 4.0

let division =
    match percentage with
    | value when value >= 60.0 -> "First Division"
    | value when value >= 50.0 -> "Second Division"
    | value when value >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Percentage: %.2f%%" percentage
printfn "Division: %s" division
