package queueservicetypes
type QueueServicePhrase string

func ProvideQueueServicePhrase(phrase QueueServicePhrase) QueueServicePhrase {
    return phrase
}