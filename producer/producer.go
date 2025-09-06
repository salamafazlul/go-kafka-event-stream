package main

import(
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

func init() {
    godotenv.Load()
}

type Comment struct{
	Text string `form:"text" json:"text"`
}

func main(){
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	app.Listen(":3000")
}

func ConnectProducer(brokersUrl []string)(sarama.SyncProducer, error){
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil{
		return nil, err
	}
	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error{
	brokersUrl := []string{os.Getenv("KAFKA_BROKERS")}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil{
		return err
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err!= nil{
		return err
	}

	fmt.Printf("Message is stored in tpoic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error{
	cmt := new(Comment)
	if err := c.BodyParser(cmt); err != nil{
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	cmtInBytes, err := json.Marshal(cmt)
	PushCommentToQueue("comments", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})
	if err != nil{
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return err
}