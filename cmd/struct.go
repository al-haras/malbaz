package cmd

type Samples struct {
    Samples []Sample `json:"data"`
}

// Sample struct that contains relevant info for sample
type Sample struct {
	SHA256     	string    	`json:"sha256_hash"`
	SHA384	   	string    	`json:sha3_384_hash`
	SHA1		string		`json:sha1_hash`
	MD5		string		`json:md5_hash`
	FirstSeen	string		`json:first_seen`
	LastSeen	string		`json:last_seen`
	FileName	string		`json:file_name`
	FileSize	int		`json:file_size`
	FileType	string		`json:file_type`
	Reporter	string		`json:reporter`
    	Signature  	string    	`json:"signature"`
    	Tags       	[]string  	`json:"Tags"`
}
