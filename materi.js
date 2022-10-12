console.log("Hello world!")

// Variable

// var
// let
// const - constant

// var - bisa di deklarasikan ulang
var name = "Abel Dustin"
var name = "Faiq Ulul"

// let - gak bisa di deklarasikan ulang
let food = "Seblak"
food = "Martabak"

// const - gak bisa di deklarasikan ulang dan gak bisa di ubah value nya
// object
const drink = "es teh manis"
// drink = "es teh manis panas"

// Data Type

// string - "Tere Liye" | "Bpk Budi"
// number - 100 | 10.00
// boolean - true/false

var school = "SMP Negeri 17 Jakarta" //string

var age = 24 // number

var isStudying = false

// camelcase
var veggies = "kangkung"

const foodGradeQuality = "A"

// console.log(foodGradeQuality)

// console.log("Hi, my name is Abel Dustin")
// console.log("Hi, my name is "+name)
// console.log("Hi, my name is", name, "I like to eat", veggies)
// console.log(`Hi, my name is ${name}, I like to eat ${veggies}`)

// operator
let x = 5.5
let y = 4.5

let result = x + y
// console.log(result)

// condition
// if - else if - else
let word = "60"

if (word == 60) {
    // console.log("Kok kenal?")
} else {
    // console.log("Maaf gak kenal")
}

// function
function Aritmatika() {
    let bilanganPertama = 50
    let bilanganKedua = 70

    let result = bilanganPertama + bilanganKedua

    console.log(result)
}

Aritmatika()

function Aritmatika2(bilanganPertama, bilanganKedua) {
    let result  = bilanganPertama + bilanganKedua

    console.log(result)
}

Aritmatika2(40, 80)

function myName(name) {
    console.log(name)
}

myName("Achmad Rifai")