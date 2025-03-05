logic 文件夹内存放rpc handler和http handler的具体业务实现, 因为这两个handler可能复用同样的业务逻辑

| echo http | kitex rpc |
    |              |
    v              v
|         logic         |
    |              |
    v              v
|   model   | other rpc server|