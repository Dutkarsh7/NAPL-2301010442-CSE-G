module PrimeFs

open System

let rec private checkDivisor number divisor =
    if divisor * divisor > number then
        true
    elif number % divisor = 0 then
        false
    else
        checkDivisor number (divisor + 1)

let isPrime number =
    number > 1 && checkDivisor number 2

printf "Enter an integer: "
let number = Int32.Parse(Console.ReadLine())

if isPrime number then
    printfn "%d is prime." number
else
    printfn "%d is not prime." number
