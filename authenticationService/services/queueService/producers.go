package queueservice

import (
	"context"
	"fmt"
	"time"

	globalutilities "github.com/AFORANURAG/microServices/authenticationService/utilityFunctions/globalUtilities"
	"github.com/rabbitmq/amqp091-go"
)

type Producer struct{
	q *QueueService
}

func (p *Producer) SendEmail(msg string)(error){
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
ch,err:=p.q.CreateEmailServiceChannel()
globalutilities.FailOnError(err,"Error while calling CreateEmailServiceChannel")

err = ch.PublishWithContext(ctx,
  "logs", // exchange
  "",     // routing key
  false,  // mandatory
  false,  // immediate
  amqp091.Publishing{
          ContentType: "text/plain",
          Body:        []byte(msg),
  })
  globalutilities.FailOnError(err,fmt.Sprintf("Error while publishing"))
  return nil
}

func NewProducerServiceProvider(url string)(*Producer,error){
	queueService,err:=NewQueueServiceProvider(url)
globalutilities.FailOnError(err,"Error instantiating queueService Provider")
	return &Producer{q:queueService},nil
}