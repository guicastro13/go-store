package httperr

type RestErr struct {
  Message string `json:"message"`
  Err string `json:"error,omitempty"`
  Code int `json:"code"`
  Fields []Fields `json:"fields,omitempty"`
}

type Fields struct {
  Field string `json:"field"`
  Value interface{} `json:"value, omitempty"`
  Message string `json:"message"`
}

