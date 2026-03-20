# go-todolist-rest-api

CreateTask
pattern: /tasks
method:  POST
info:    JSON in HTTP request body

GetTask
pattern: /tasks/{title}
method:  GET
info:    pattern

GetAllTask
pattern: /tasks
method:  GET
info:    -

GetAllUnconpletedTask
pattern: /tasks?completed=false
method:  GET
info:    query params

CompleteTask
pattern: /tasks/{title}
method:  PATCH
info:    pattern + JSON in request body

DeleteTask
pattern: /tasks/{title}
method:  DELETE
info:    pattern
