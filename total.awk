/Benchmark/ { total += $3 } END { printf("%d ns, %d μs, %d ms, %d s\n",  total, total/1000, total/1000/1000, total/1000/1000/1000) }
