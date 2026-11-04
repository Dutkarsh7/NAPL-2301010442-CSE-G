module Pattern

open System
open System.Text.RegularExpressions

let reverseString (text: string) =
    text |> Seq.rev |> String.Concat

let isPalindrome (text: string) =
    let normalized = text.ToLowerInvariant()
    normalized = reverseString normalized

let transformString (text: string) =
    match text with
    | value when value |> Seq.forall Char.IsLetter -> value.ToUpperInvariant()
    | value when value |> Seq.forall Char.IsDigit -> value + value
    | value -> value

let isValidEmail (email: string) =
    let pattern = "^[^@\\s]+@[^@\\s]+\\.[^@\\s]+$"
    Regex.IsMatch(email, pattern)

printfn "Palindrome checks:"
[ "madam"; "hello"; "Racecar" ]
|> List.iter (fun text ->
    printfn "%s -> reversed: %s, palindrome: %b" text (reverseString text) (isPalindrome text))

printfn "\nString transformations:"
[ "hello"; "1234"; "abc123" ]
|> List.iter (fun text -> printfn "%s -> %s" text (transformString text))

printfn "\nEmail validation:"
[ "student@example.com"; "invalid-email"; "user@college.org" ]
|> List.iter (fun email -> printfn "%s -> valid: %b" email (isValidEmail email))
