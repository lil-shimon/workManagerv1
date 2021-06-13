type Task struct {
	Name string `json:"name"`
}

func CreateType (c echo.Context) error {
	type := new(Task)
	if err := c.Bind(type); err != nil {
		return err
	}
	return c.Json(http.StatusCreated, type)
}