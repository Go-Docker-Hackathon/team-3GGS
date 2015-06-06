package  mate
import(
    "time"
	"math/rand"
)

func  IdentifyCode() int {
	 r := rand.New(rand.NewSource(time.Now().UnixNano()))
	 return r.Intn(100)
}
