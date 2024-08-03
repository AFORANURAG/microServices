package queueservice

import (
	"context"
	"fmt"
	"time"

	queueConstants "github.com/AFORANURAG/microServices/authenticationService/constants/queue"
	globalutilities "github.com/AFORANURAG/microServices/authenticationService/utilityFunctions/globalUtilities"
	"github.com/rabbitmq/amqp091-go"
)

type Producer struct{
	q *QueueService
}

func (p *Producer) SendEmail(msg string)(bool,error){
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
  globalutilities.FailOnError(err,"Error while publishing")
  return true,nil
}


func (p *Producer) SendOtp(msg []byte)(bool,error){
  fmt.Println(msg)
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
ch,err:=p.q.CreateOTPServiceChannel()
globalutilities.FailOnError(err,"Error while calling CreateEmailServiceChannel")

err = ch.PublishWithContext(ctx,
  queueConstants.OTPServiceMessageBrokerValues["exchange"], // exchange
  "",     // routing key
  false,  // mandatory
  false,  // immediate
  amqp091.Publishing{
          ContentType: "text/plain",
          Body:        []byte(msg),
  })
  globalutilities.FailOnError(err,"Error while publishing")
  return true,nil
}

func NewProducerServiceProvider( q *QueueService)(*Producer){
	// queueService,err:=NewQueueServiceProvider(url)
// globalutilities.FailOnError(err,"Error instantiating queueService Provider")
	return &Producer{q}
}