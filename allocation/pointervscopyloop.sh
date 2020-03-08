#!/bin/bash
go test ./... -bench=BenchmarkMemoryHeapLoop -benchmem -run=^$ -count=10 > headloop.txt && benchstat headloop.txt
go test ./... -bench=BenchmarkMemoryStackLoop -benchmem -run=^$ -count=10 > stackloop.txt && benchstat stackloop.txt