package domain

type Code struct {
	codeID string
	userID string
	code   string
	result string
	time   string
	Type   string
}

func (c *Code) CodeID() string {
	return c.codeID
}

func (c *Code) SetCodeID(codeID string) {
	c.codeID = codeID
}

func (c *Code) UserID() string {
	return c.userID
}

func (c *Code) SetUserID(userID string) {
	c.userID = userID
}

func (c *Code) Code() string {
	return c.code
}

func (c *Code) SetCode(code string) {
	c.code = code
}

func (c *Code) Result() string {
	return c.result
}

func (c *Code) SetResult(result string) {
	c.result = result
}

func (c *Code) Time() string {
	return c.time
}

func (c *Code) SetTime(time string) {
	c.time = time
}

func (c *Code) GetType() string {
	return c.Type
}

func (c *Code) SetType(Type string) {
	c.Type = Type
}
