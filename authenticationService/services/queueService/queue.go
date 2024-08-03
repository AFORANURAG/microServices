package queueservice

import (
	"fmt"

	queueConstants "github.com/AFORANURAG/microServices/authenticationService/constants/queue"
	queueservicetypes "github.com/AFORANURAG/microServices/authenticationService/types/queueServiceTypes"
	globalUtilities "github.com/AFORANURAG/microServices/authenticationService/utilityFunctions/globalUtilities"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

// the work of this service is to expose channel(logical tcp connections) through which publish and subscription  would work.
// we can create an interface but
type QueueService struct{
c *amqp.Connection
// that will be provided externally 
}
/*
The below function basically opens a new channel and declare email service queue and exchange
*/
func (q *QueueService) CreateEmailServiceChannel()(*amqp.Channel,error){
	ch,err:=q.c.Channel()
	globalUtilities.FailOnError(err, "Failed to open a channel")
	queue,err:=ch.QueueDeclare(queueConstants.EmailServiceMessageBrokerValues["queue"],true,false,false,false,nil)
	fmt.Println(queue)
	globalUtilities.FailOnError(err,fmt.Sprintf("Failed to declare queue %s",queueConstants.EmailServiceMessageBrokerValues["queue"]))
	
    err=ch.ExchangeDeclare(queueConstants.EmailServiceMessageBrokerValues["exchange"],"fanout",true,false,false,false,nil)
	globalUtilities.FailOnError(err,"Failed to declare exchange")
return ch, nil
	// declare context with timeout so that the process will be cancelled after 5 second to stop it running till timout or indefinetely	
}


func (q *QueueService) CreateOTPServiceChannel()(*amqp.Channel,error){
	ch,err:=q.c.Channel()
	globalUtilities.FailOnError(err, "Failed to open a channel")
	queue,err:=ch.QueueDeclare(queueConstants.OTPServiceMessageBrokerValues["queue"],true,false,false,false,nil)
	// fmt.Println(queue)
	globalUtilities.FailOnError(err,fmt.Sprintf("Failed to declare queue %s",queueConstants.OTPServiceMessageBrokerValues["queue"]))
	
    err=ch.ExchangeDeclare(queueConstants.OTPServiceMessageBrokerValues["exchange"],"fanout",true,false,false,false,nil)
	globalUtilities.FailOnError(err,"Failed to declare exchange")

	err=ch.QueueBind(queue.Name,"",queueConstants.EmailServiceMessageBrokerValues["exchange"],false,nil)
    globalUtilities.FailOnError(err,"Failed to create queue bindings for otpService Channel")
return ch, nil
	// declare context with timeout so that the process will be cancelled after 5 second to stop it running till timout or indefinetely	
}





func NewQueueServiceProvider(url queueservicetypes.QueueServicePhrase)(*QueueService){
	conn,err:=amqp091.Dial((string)(url))
	globalUtilities.FailOnError(err,"Failed to open connection")
	return &QueueService{c:conn}
}