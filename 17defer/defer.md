# Defer

- In Go, defer makes a function run just before the current function ends.
- It allows you to ensure that certain clean-up tasks, such as closing a file or releasing a lock, are performed no matter how the function exits (whether normally or due to an error).

### Key Points to Remember
- Execution Timing: Deferred functions are executed in LIFO (Last In, First Out) order, just before the surrounding function returns.
- Common Use Cases: Resource management, such as closing files, releasing locks, or cleaning up temporary resources.
- Readability: Using defer makes the code more readable and maintainable by keeping resource management close to resource acquisition.

