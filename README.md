Todo: 
 - /certificates


 Example Code:
 ```golang
import (
	"github.com/brianpowell/go-acuity"
)

func main() {

  ac := acuity.Setup("userID", "apiKey")
  
  // Get the User information that is accessing the API
	me, err := ac.GetMe() 
  
  // Query the Appointments
  query := map[string]interface{}{
    // "appointmentTypeID": 000,
    // "calendarID": 000,
  }
  out, err := ac.GetAppointments() 
}
 ```
 