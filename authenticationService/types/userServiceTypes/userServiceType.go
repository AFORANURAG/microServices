package userservicetypes
type UserServicePhrase string
func ProvideUserServicePhrase(phrase UserServicePhrase) UserServicePhrase {
    return phrase
}