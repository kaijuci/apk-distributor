package catalog

import (
	"log"

	"github.com/kaijuci/apk-distributor/provider/firebase"
	"github.com/spf13/cobra"
)

func NewFirebaseCmd(firebaseCredentialsFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "firebase",
		Short: "Distribute APKs using Firebase App Distribution",
		Run: func(cmd *cobra.Command, args []string) {
			f, err := firebase.NewFirebaseDistributor(firebaseCredentialsFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Distributor: %v", f)
		},
	}
}
