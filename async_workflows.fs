module AsyncWorkflows

open System

let runTask taskName (delay: int) result = async {
    printfn "%s started" taskName
    do! Async.Sleep delay
    printfn "%s completed after %d ms" taskName delay
    return result
}

let tasks =
    [ runTask "Task 1" 1200 "Result from Task 1"
      runTask "Task 2" 400 "Result from Task 2"
      runTask "Task 3" 800 "Result from Task 3" ]

printfn "Starting parallel tasks..."
let results = tasks |> Async.Parallel |> Async.RunSynchronously

printfn "All tasks completed. Results:"
results |> Array.iter (printfn "%s")
