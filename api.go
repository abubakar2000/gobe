package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

// Blockstructure
type BlockData struct {
	Title             string
	Description       string
	Owners            []string
	Problem           string
	Domain            []string
	Technologies_used []string
	Viewing_price     float32
	Ownership_price   float32
	Pricing_history   []float32
}

// Blockchain structure
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

// For proposedIdeas cache
type ProposedIdea struct {
	BlockData BlockData
	SimIdea   string
	SimScore  float32
}

type User struct {
	Username    string
	Password    string
	Balance     float32
	Email       string
	PhoneNumber string
}

// expect head pointer which will now point to the tail of the blockchain
func InsertBlock(chainHead *Block, newBlock BlockData) *Block {
	if chainHead != nil {
		newBlock := &Block{
			Data:        newBlock,
			PrevHash:    "ph2",
			CurrentHash: "ch2",
			PrevPointer: chainHead,
		}
		chainHead = newBlock
	} else {
		chainHead = &Block{
			Data:        newBlock,
			PrevHash:    "ph1",
			CurrentHash: "ch1",
			PrevPointer: nil,
		}
	}
	return chainHead
}

func PrintChain(chainHead *Block) {
	cursor := chainHead
	for {
		if cursor != nil {
			println(cursor.Data.Title)
			println(cursor.Data.Viewing_price)
			println(cursor.Data.Description)
			cursor = cursor.PrevPointer
		} else {
			return
		}
	}
}

func main() {

	var ChainHead *Block

	blockd := BlockData{
		Title:             "ME",
		Description:       "IM",
		Problem:           "Ima mess",
		Owners:            []string{},
		Domain:            []string{},
		Technologies_used: []string{},
		Viewing_price:     1.4,
		Ownership_price:   0,
		Pricing_history:   []float32{},
	}
	ChainHead = InsertBlock(ChainHead, blockd)
	// PrintChain(ChainHead)

	// -------------SERVER SIDE PROGRAM----------
	fastergoding.Run()
	server := fiber.New()
	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to GOBE")
	})
	server.Listen(":8081")

	// API ENDPOINTS

}
