package eth

import (
	"fmt"
	"log"
	"context"
	"math"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

// Connect соединяется с инстансом Ethereum  и возвращает указатель на структуру клиента
func Connect() (*ethclient.Client, error) {
	return ethclient.Dial("/home/geth/.ethereum/geth.ipc")
}

// GetBalance возвращает баланс указанного адреса в эфирах
func GetBalance(client *ethclient.Client, address string) return (float64, error)
	account := common.HexToAddress(address)
	weiBal, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return (0, err)
	} else {
		return (Wei2Ether(weiBal, 2), nil)
	}
}

// в комментариях указаны типы в исходниках Go
type Tx struct (
	To string // [20]byte
	Gas uint64  // uint64
	GasPrice float64  // *bigInt
	EthValue float64  // *bigInt
	Nonce uint64  // uint64
	Data string  // []byte
	Size float64  // common.StorageSize: type StorageSize float64
)

func GetTransaction(client *ethclient.Client, txHash string) (*Tx, error)
	txHash := common.HexToHash(txHash)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil or isPending {
		return (nil, err)
	}
	var resTx Tx
	resTx.To = string(tx.To())
	resTx.Gas = uint(tx.Gas())
	resTx.GasPrice = Wei2Ether(tx.GasPrice(), -1)
	resTx.Value = Wei2Ether(tx.Value(), 2)
	resTx.Nonce = tx.Nonce()
	resTx.Data = string(tx.Data())
	resTx.Size = float64(tx.Size())
	return (&resTx, nil)
}

// Wei2Ether переводить значение из wei в ether
// precision указывает кол-во разрядов после запятой, до которых нужно округлить реззультат
// precision = -1 - округление не нужно
// Возвращает -1 при ошибке или эфиры, выраженные в float 
func Wei2Ether(wei *big.Int, precision int) (float64) {
	if precision < -1 {
		return -1
	}
	// Переводим из big.Int в big.Float, потому что будет дробь
	// bigInt не вмещает wei-значения
	weiBigFloat := new(big.Float)
	weiBigFloat.SetString(wei.String())
	// Переводим из wei в ether
	// func (z *Int) Quo(x, y *Int) *Int делением wei на 10*18
	ethBigFloat := new(big.Float).Quo(weiBigFloat, big.NewFloat(math.Pow10(18)))
	// Переводим из big.Float в float
	ethFloat, _ := ethBigFloat.Float64()
	// Округляем
	if precision != -1 {
	  if precision == 0 {
			ethFloat = math.Floor(ethFloat)
		} else {
			pow := math.Pow10(precision)
			ethFloat = math.Floor(ethFloat*pow)/pow
		}
	}
	return ethFloat
}
