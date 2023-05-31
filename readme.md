# checkers

基于Cosmos SDK和Tendermint的区块链游戏.

### 配置

```shell
vim config.yml
```

### 开始

```shell
checkersd start
```

### 查看系统信息

```shell
checkersd query checkers show-system-info
```

### 创建玩家

```shell
checkersd keys add zwz
```

### 查看余额
```shell
export alice=$(checkersd keys show alice -a) 
export bob=$(checkersd keys show bob -a)
checkersd query bank balances $alice
checkersd query bank balances $bob
checkersd query bank balances cosmos16xx0e457hm8mywdhxtmrar9t09z0mqt9x7srm3
```

### 创建棋局

```shell
checkersd tx checkers create-game \
    $alice $bob 1 token \
    --from $alice -y
```

### 估算gas费

```shell
checkersd tx checkers create-game $alice $bob 1000000 --from $alice --dry-run -y
```

### 查看棋局

```shell
checkersd query checkers show-stored-game 1
```

### 查看棋盘

```shell
checkersd query checkers show-stored-game 1 --output json | jq ".storedGame.board" | sed 's/"//g' | sed 's/|/\n/g'
```

### 查看是否可执行

```shell
checkersd query checkers can-play-move 1 w 1 2 2 3
checkersd query checkers can-play-move 1 r 0 5 1 4
checkersd query checkers can-play-move 1 b 0 5 1 4
checkersd query checkers can-play-move 1 b 1 2 2 3
```

### 移动棋子

```shell
checkersd tx checkers play-move 1 1 2 2 3 --from $alice -y
```

### 查看所有棋局

```shell
checkersd query checkers list-stored-game
```

