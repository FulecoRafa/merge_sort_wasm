import fs from "node:fs";

import Go from './wasm_exec.js'
import { mergesort } from './merge_sort.mjs'
import { fibonacci } from "./fibb.mjs";

const go = new Go();
const buf = fs.readFileSync('./main.wasm');
const wasm = await WebAssembly.instantiate(new Uint8Array(buf), go.importObject);
go.run(wasm.instance)

const readFileLines = filename =>
  fs
    .readFileSync(filename)
    .toString('UTF8')
    .split('\n');

function loadExample() {
    return readFileLines('../input/input.txt')
        .map(str => parseInt(str))
}

const fibb = 25

console.time("Node Fibonacci")
const resultFibb = fibonacci(fibb)
console.timeEnd("Node Fibonacci")
console.time("GoFibonacci")
let resultFibbGo = GoFibonacci(fibb)
console.timeEnd("GoFibonacci")

const arr = loadExample()
console.time("Node MergeSort")
const orderedNode = mergesort(arr)
console.timeEnd("Node MergeSort")
console.time("GO parallel MergeSort")
const orderedParallel = GoMergeSort(arr)
console.timeEnd("GO parallel MergeSort")
console.time("GO MergeSort")
const ordered = GoMergeSort(arr)
console.timeEnd("GO MergeSort")
console.log(resultFibb)
console.log(resultFibbGo)
console.log(orderedNode)
console.log(ordered)
console.log(orderedParallel)
