package requests

type Text struct {
	MessageTool				string  `json:"MessageTool"`
	Language          		string  `json:"Language"`
	MessageToolName			string  `json:"MessageToolName"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
