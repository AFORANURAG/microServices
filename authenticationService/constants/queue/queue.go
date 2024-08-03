package queueConstants
type ExchangeType int
type TutorTestSubjectUpdateType int


var EmailServiceMessageBrokerValues = map[string]string{
    "exchange":         "emailServiceExchange",
    "queue":            "emailServiceQueue",
}

var OTPServiceMessageBrokerValues=map[string]string{
    "exchange":"otpServiceExchange",
    "queue":"otpServiceQueue",
}