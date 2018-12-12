package main 

import (
	"github.com/33cn/libbyz-go/replica"
)

func main() {
	replica.ByzInitReplica("./bft/config", "./bft/config_private/template")
}
