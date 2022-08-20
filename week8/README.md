# week8

#### 测试指令
```
# 性能统计
redis-benchmark  -d 10  -r 10000 -n 50000 -t get,set --csv
-d 设置value大小
-r 设置随机key的数量
-t 设置测试的方法
--csv 以.csv的格式输出统计结果

# 内存使用统计
redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
```

#### 测试内容
##### 统计性能
内容：使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能

1. 10 字节
```
root@dev-host redis# redis-benchmark  -d 10  -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","53763.44","0.715","0.088","0.583","1.567","2.647","14.799"
"GET","58072.01","0.683","0.136","0.567","1.479","2.343","8.271"

root@dev-host redis# redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
"keys.bytes-per-key",99,"dataset.bytes",462352
```

2. 20 字节
```
root@dev-host redis# redis-benchmark  -d 20  -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","60240.96","0.648","0.144","0.527","1.375","2.159","11.543"
"GET","58343.06","0.681","0.128","0.543","1.543","2.295","8.767"

root@dev-host redis# redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
"keys.bytes-per-key",107,"dataset.bytes",541840
```

3. 50 字节
```
root@dev-host redis# redis-benchmark  -d 50  -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","56306.30","0.713","0.080","0.575","1.511","2.519","9.335"
"GET","56497.18","0.703","0.136","0.567","1.535","2.375","9.855"

root@dev-host redis# redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
"keys.bytes-per-key",139,"dataset.bytes",860112
```


4. 100 字节
```
root@dev-host redis# redis-benchmark  -d 100  -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","58754.41","0.680","0.104","0.575","1.415","2.111","23.071"
"GET","57142.86","0.685","0.120","0.567","1.487","2.167","17.295"

root@dev-host redis# redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
"keys.bytes-per-key",194,"dataset.bytes",1418832
```


5. 1K 字节
```
root@dev-host redis# redis-benchmark  -d 1000  -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","53879.31","0.751","0.128","0.639","1.575","2.431","9.743"
"GET","55617.35","0.696","0.152","0.575","1.471","2.351","7.071"

root@dev-host redis# redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
"keys.bytes-per-key",1101,"dataset.bytes",10491288
```


6. 5K 字节
```
root@dev-host redis# redis-benchmark  -d 5000  -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","44247.79","0.996","0.192","0.895","1.895","2.991","11.543"
"GET","54824.56","0.741","0.152","0.631","1.495","2.263","10.623"

root@dev-host redis# redis-cli  --csv MEMORY STATS  |grep -o '"keys.bytes-per-key",[0-9]*,"dataset.bytes",[0-9]*'
"keys.bytes-per-key",5170,"dataset.bytes",51199256
```

#### 测试结论
1. 随机写key总数不变的情况下，value越大，按p95统计，性能基本平稳。
2. 随机写key总数不变的情况下，value越大，平均key大小所占空间越多