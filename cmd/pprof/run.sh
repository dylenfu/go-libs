#!/bin/bash

go build
./pprof --heap=heap.txt
#./pprof --cpu=cpu.prof --heap=heap.prof
