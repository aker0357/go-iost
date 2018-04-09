package core

import (
	"fmt"

	"github.com/iost-official/prototype/common"
)

//go:generate mockgen -destination mocks/mock_tx_pool.go -package core_mock github.com/iost-official/prototype/core TxPool
type TxPool interface {
	Add(tx Tx) error
	Del(tx Tx) error
	Find(txHash []byte) (Tx, error)
	GetSlice() ([]Tx, error)
	Has(txHash []byte) (bool, error)
	Size() int
}

type TxPoolImpl struct {
	TxPoolRaw
	txMap map[string]Tx
}

func (tp *TxPoolImpl) Add(tx Tx) error {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
	}
	tp.txMap[common.Base58Encode(tx.Hash())] = tx
	return nil
}

func (tp *TxPoolImpl) Del(tx Tx) error {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
	}
	delete(tp.txMap, common.Base58Encode(tx.Hash()))
	return nil
}

func (tp *TxPoolImpl) Find(txHash []byte) (Tx, error) {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
		return Tx{}, fmt.Errorf("not found")
	}

	tx, ok := tp.txMap[common.Base58Encode(txHash)]
	if !ok {
		return tx, fmt.Errorf("not found")
	}
	return tx, nil
}

func (tp *TxPoolImpl) Has(txHash []byte) (bool, error) {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
		return false, nil
	}
	_, ok := tp.txMap[common.Base58Encode(txHash)]
	return ok, nil
}

func (tp *TxPoolImpl) GetSlice() ([]Tx, error) {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
		return []Tx{}, nil
	}
	var txs []Tx
	for _, v := range tp.txMap {
		txs = append(txs, v)
	}

	return txs, nil
}

func (tp *TxPoolImpl) Size() int {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
		return 0
	}
	return len(tp.txMap)
}

func (tp *TxPoolImpl) Encode() []byte {
	for k, v := range tp.txMap {
		tp.TxHash = append(tp.TxHash, common.Base58Decode(k))
		tp.Txs = append(tp.Txs, v)
	}
	bytes, err := tp.Marshal(nil)
	if err != nil {
		panic(err)
	}
	tp.TxHash = [][]byte{}
	tp.Txs = []Tx{}
	return bytes
}

func (tp *TxPoolImpl) Decode(a []byte) error {
	_, err := tp.Unmarshal(a)
	if err != nil {
		return err
	}
	for i, v := range tp.TxHash {
		tp.txMap[common.Base58Encode(v)] = tp.Txs[i]
	}
	return nil
}

func (tp *TxPoolImpl) Hash() []byte {
	return common.Sha256(tp.Encode())
}
