package tools

import (
    "math/rand"
    "time"
    "strings"
)

func GenerateRandomPassword(length int) string {
    var lettersAndDigits = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    rand.Seed(time.Now().UnixNano())
    var sb strings.Builder
    for i := 0; i < length; i++ {
        sb.WriteRune(lettersAndDigits[rand.Intn(len(lettersAndDigits))])
    }
    return sb.String()
}


