package block

import (
	"fmt"
	"os"
	"sixletters/simple-db/pkg/consts"
)

type BlockManager interface {
	// Function fetches the last ID assigned in the contiguous block volume
	GetLastID() (int64, error)
	// Fetches and returns the underlying file
	GetBlockFile() *os.File
	// Checks if the rootblock already exists
	RootBlockExists() bool
	// Adds a block and returns it with the latest generated ID
	AddBlock() (*Block, error)
	// Writes and commits the block to disk
	WriteBlock(block *Block) error
	// Returns the root block with ID 0
	GetRootBlock() (*Block, error)
	// Fetches and returns a block by ID
	GetBlockByID(index int64) (*Block, error)
	// Returns the max number of items that can be stored in a block
	GetMaxItemsSize() int
	// Generates the next block ID
	GenerateBlockID() (uint64, error)
}

type blockManager struct {
	file *os.File
}

func NewBlockManager(file *os.File) BlockManager {
	return &blockManager{file: file}
}

func (bm *blockManager) GetLastID() (int64, error) {
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

func (bm *blockManager) GetBlockFile() *os.File {
	return bm.file
}

func (bm *blockManager) RootBlockExists() bool {
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

func (bm *blockManager) AddBlock() (*Block, error) {
	latestID, err := bm.GetLastID()
	block := &Block{}
	if err != nil {
		block.Id = 0
	} else {
		block.Id = uint64(latestID + 1)
	}
	err = bm.WriteBlock(block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (bm *blockManager) WriteBlock(block *Block) error {
	buffer := block.IntoBytes()
	offset := consts.BlockSize * block.Id
	_, err := bm.file.WriteAt(buffer, int64(offset))
	return err
}

func (bm *blockManager) GetRootBlock() (*Block, error) {
	if bm.RootBlockExists() {
		return bm.GetBlockByID(0)
	}
	return bm.AddBlock()
}

func (bm *blockManager) GetBlockByID(index int64) (*Block, error) {
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

func (bm *blockManager) GetMaxItemsSize() int {
	return consts.MaxLeafSize
}

func (bm *blockManager) GenerateBlockID() (uint64, error) {
	lastID, err := bm.GetLastID()
	if err != nil {
		return 0, err
	}
	return uint64(lastID + 1), nil
}
