# cosmosProject

基于Cosmos SDK和Tendermint的区块链游戏.

## 开始

```shell
ignite chain serve
```

### 配置

```shell
vim config.yml
```

### 创建棋局

```shell
export alice=$(checkersd keys show alice -a) $ export bob=$(checkersd keys show bob -a)
checkersd tx checkers create-game $alice $bob --from $alice 
```

### 查看棋盘

```shell
checkersd query checkers show-stored-game 1 --output json | jq ".storedGame.board" | sed 's/"//g' | sed 's/|/\n/g'
```

### 移动棋子

```shell
checkersd tx checkers play-move 1 1 2 2 3 --from $alice
```

### 查看所有棋局

```shell
checkersd query checkers list-stored-game
```
