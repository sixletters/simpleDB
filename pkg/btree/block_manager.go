package btree

import (
	"fmt"
	"os"
	"sixletters/kv-store/pkg/consts"
)

type BlockManager struct {
	file *os.File
}

func NewBlockManager(file *os.File) *BlockManager {
	return &BlockManager{file: file}
}

func (bm *BlockManager) GetLastID() (int64, error) {
	fileData, err := bm.file.Stat()
	if err != nil {
		return -1, err
	}
	fileSizeBytes := fileData.Size()

	// Empty file, no latest ID
	if fileSizeBytes == 0 {
		return -1, nil
	}

	// for e.g, if only the root block exists the fileSizeBytes will be equal to one block size
	// which means the latest one has an ID of 0
	return (fileSizeBytes / consts.BlockSize) - 1, nil
}

func (bm *BlockManager) RootBlockExists() bool {
	lastID, err := bm.GetLastID()
	if err != nil {
		// if last ID cannot be retrieved means, root block definitely does not exist
		// file cannot be opened
		return false
	}
	// Empty file
	if lastID == -1 {
		return false
	}
	return true
}

func (bm *BlockManager) AddBlock() (*Block, error) {
	latestID, err := bm.GetLastID()
	block := &Block{}
	if err != nil {
		block.Id = 0
	} else {
		block.Id = uint64(latestID + 1)
	}
	err = bm.WriteBlockToDisk(block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (bm *BlockManager) WriteBlockToDisk(block *Block) error {
	buffer := block.IntoBytes()
	offset := consts.BlockSize * block.Id
	_, err := bm.file.WriteAt(buffer, int64(offset))
	return err
}

func (bm *BlockManager) GetRootBlock() (*Block, error) {
	if bm.RootBlockExists() {
		return bm.GetBlockByID(0)
	}
	return bm.AddBlock()
}

func (bm *BlockManager) GetBlockByID(index int64) (*Block, error) {
	if index < 0 {
		return nil, fmt.Errorf("index is less than 0: %d", index)
	}
	offset := index * consts.BlockSize

	byteBuffer := make([]byte, consts.BlockSize)
	_, err := bm.file.ReadAt(byteBuffer, offset)
	if err != nil {
		return nil, err
	}
	return NewBlock().FromBytes(byteBuffer), nil
}
