# methods

- Go does not have classes. However, you can define methods on types.
- If you want to export the methods or if you want to make a public methods, the first letter of the Methods must be in CAPITAL. e.g. GetStatus()
- A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type. 
- Example func (t Type) methodName(parameter list) \
  Here, (t Type) is a receiver.