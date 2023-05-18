#!/usr/bin/gnuplot --persist

data = ARG1
out = ARG2

set terminal png size 1200,800
set output out

plot data with lp title 'Sol' lw 4
