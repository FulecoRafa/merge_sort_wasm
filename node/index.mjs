import fs from "fs";

import Go from './wasm_exec.js'

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

const arr = loadExample()
console.time("GO")
const ordered = GoMergeSort(arr)
console.timeEnd("GO")
console.log(ordered)
