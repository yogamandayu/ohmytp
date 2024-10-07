package entity

// Requestor is a requester of all action.
type Requestor struct {
	Metadata RequestorMetadata
}

// RequestorMetadata is a metadata of requestor.
type RequestorMetadata struct {
	RequestID string
	IPAddress string
	UserAgent string
}
