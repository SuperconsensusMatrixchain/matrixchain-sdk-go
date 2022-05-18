package common

import (
	"github.com/SuperconsensusMatrixchain/matrixchain/service/common"
	"github.com/SuperconsensusMatrixchain/matrixchain/service/pb"
)

// MakeTransactionID 计算交易ID，包括签名。
func MakeTransactionID(tx *pb.Transaction) ([]byte, error) {
	return common.MakeTxId(tx)
}

// MakeTxDigestHash 计算交易哈希，不包括签名。
func MakeTxDigestHash(tx *pb.Transaction) ([]byte, error) {
	return common.MakeTxDigestHash(tx)
}
