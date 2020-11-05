/* models.go
// Acuity API interface
*/

package acuity

type (

	// AddOn - model for appt
	AddOn struct {
		ID       int    `json:"id"`
		Name     string `json:"string"`
		Duration int    `json:"duration"`
		Price    string `json:"price"`
		Private  bool   `json:"private"`
	}

	// Appointment Structure
	Appointment struct {
		ID                  int     `json:"id,omitempty"`
		FirstName           string  `json:"firstName"`
		LastName            string  `json:"lastName"`
		Email               string  `json:"email"`
		Phone               string  `json:"phone"`
		Date                string  `json:"date"`
		Time                string  `json:"time"`
		EndTime             string  `json:"endTime"`
		DateCreated         string  `json:"dateCreated"`
		DateTime            string  `json:"datetime"`
		Price               string  `json:"price"`
		Paid                string  `json:"paid"`
		AmountPaid          string  `json:"amountPaid"`
		Type                string  `json:"type"`
		ApptTypeID          int     `json:"appointmentTypeID"`
		AddOnIDs            []int   `json:"addonIDs"`
		ClassID             string  `json:"classID"`
		Duration            string  `json:"duration"`
		Calendar            string  `json:"calendar"`
		CalendarID          int     `json:"calendarID"`
		CanClientCancel     bool    `json:"canClientCancel"`
		CanClientReschedule bool    `json:"canClientReschedule"`
		Location            string  `json:"location"`
		Certificate         string  `json:"certificate"`
		ConfirmationPage    string  `json:"confirmationPage"`
		Notes               string  `json:"notes"`
		Timezone            string  `json:"timezone"`
		Forms               []Form  `json:"forms"`
		Labels              []Label `json:"labels"`
	}

	// AppointmentBody Structure
	AppointmentBody struct {
		DateTime    string      `json:"datetime"`
		ApptTypeID  int         `json:"appointmentTypeID"`
		CalendarID  int         `json:"calendarID,omitempty"`
		FirstName   string      `json:"firstName"`
		LastName    string      `json:"lastName"`
		Email       string      `json:"email"`
		Phone       string      `json:"phone,omitempty"`
		Timezone    string      `json:"timezone,omitempty"`
		Certificate string      `json:"certificate,omitempty"`
		Fields      []FieldPost `json:"fields,omitempty"`
		Notes       string      `json:"notes,omitempty"`
		AddOnIDs    []int       `json:"addonIDs,omitempty"`
		Labels      []LabelPost `json:"labels,omitempty"`
		Admin       bool        `json:"admin"`
	}

	// AppointmentType - model for type response
	AppointmentType struct {
		ID            int    `json:"id,omitempty"`
		Active        string `json:"active"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Duration      int    `json:"duration"`
		Price         string `json:"price"`
		Category      string `json:"category"`
		Color         string `json:"color"`
		Private       bool   `json:"private"`
		Type          string `json:"type"`
		ClassSize     int    `json:"classSize"`
		PaddingAfter  int    `json:"paddingAfter"`
		PaddingBefore int    `json:"paddingBefore"`
		CalendarIDs   []int  `json:"calendarIDs"`
	}

	// AvailabilityCheckTime - Response Obj
	AvailabilityCheckTime struct {
		DateTime   string `json:"datetime"`
		ApptTypeID int    `json:"appointmentTypeID"`
		CalendarID int    `json:"calendarID"`
		Valid      bool   `json:"valid"`
	}

	// AvailabilityCheckTimeBody - Submission Body
	AvailabilityCheckTimeBody struct {
		DateTime   string `json:"datetime"`
		ApptTypeID int    `json:"appointmentTypeID"`
		CalendarID int    `json:"calendarID"`
	}

	// Block - calendar blocks
	Block struct {
		ID          int    `json:"id,omitempty"`
		CalendarID  int    `json:"calendarID"`
		Description string `json:"description"`
		Start       string `json:"start"`
		End         string `json:"end"`
		Until       string `json:"until"`
		Recurring   string `json:"reurring"`
		Notes       string `json:"notes"`
	}

	// BlockBody - calendar block body
	BlockBody struct {
		CalendarID int    `json:"calendarID"`
		Start      string `json:"start"`
		End        string `json:"end"`
		Notes      string `json:"notes"`
	}

	// Calendar - calendar obj
	Calendar struct {
		ID          int    `json:"id,omitempty"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		ReplyTo     string `json:"replyTo"`
		Description string `json:"description"`
		Location    string `json:"location"`
		Timezone    string `json:"timezone"`
	}

	// Class - Appt Classes
	Class struct {
		ID             int    `json:"id,omitempty"`
		ApptTypeID     int    `json:"appointmentTypeID"`
		CalendarID     int    `json:"calendarID"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Time           string `json:"time"`
		Calendar       string `json:"calendar"`
		Duration       int    `json:"duration"`
		IsSeries       bool   `json:"isSeries"`
		Slots          int    `json:"slots"`
		SlotsAvailable int    `json:"slotsAvailable"`
	}

	// Client - customers
	Client struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Notes     string `json:"notes"`
	}

	// ClientBody - same body structure
	ClientBody Client

	// DateTime - date reponse
	DateTime struct {
		Date string `json:"date,omitempty"`
		Time string `json:"time,omitempty"`
	}

	// FieldPost Structure
	FieldPost struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	}

	// Form Structure
	Form struct {
		ID     int     `json:"id"`
		Name   string  `json:"name"`
		Values []Value `json:"values"`
	}

	// Label model for appt
	Label struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	// LabelPost model for appt
	LabelPost struct {
		ID int `json:"id"`
	}

	// Me - user obj
	Me struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		Timezone       string `json:"timezone"`
		FirstDayOfWeek int    `json:"firstDayOfWeek"`
		TimeFormat     string `json:"timeFormat"`
		Currency       string `json:"currency"`
		SchedulingPage string `json:"schedulingPage"`
		EmbedCode      string `json:"embedCode"`
		Plan           string `json:"plan"`
		Description    string `json:"description"`
	}

	// Payment model for appt
	Payment struct {
		TransactionID string `json:"transactionID"`
		Created       string `json:"created"`
		Processor     string `json:"processor"`
		Amount        string `json:"amount"`
	}

	// Product model for appt
	Product struct {
		ID             int            `json:"id"`
		Name           string         `json:"name"`
		Description    string         `json:"description"`
		Price          string         `json:"price"`
		Type           string         `json:"type"`
		Hidden         bool           `json:"hidden"`
		Expires        int            `json:"expires"`
		ApptTypeIDs    []int          `json:"appointmentTypeIDs"`
		ApptTypeCounts map[string]int `json:"appointmentTypeCounts"`
		Minutes        int            `json:"minutes"`
	}

	// Order model for appt
	Order struct {
		ID        int    `json:"id"`
		Total     string `json:"total"`
		Status    string `json:"status"`
		Time      string `json:"time"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Title     string `json:"title"`
		Notes     string `json:"notes"`
	}

	// Value model for forms
	Value struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		FieldID int    `json:"fieldID"`
		Value   string `json:"value"`
	}

	// WebHook model for forms
	Webhook struct {
		ID                string `json:"id"`
		Action            string `json:"action"`
		CalendarID        string `json:"calendarID"`
		AppointmentTypeID string `json:"appointmentTypeID"`
	}
)
