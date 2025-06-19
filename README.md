# TodoGo
A simple yet effective CLI Todo app.

## How to use
- Adding a task
    - Only the title and status parameters are required.

```go run . addTask --title="I am a new task!" --goal="I will teach you how to add a task!" --goalStatus="1" --goalNote="You can write more notes here!"```

- Listing all tasks

``` go run . list ```

- Updating a task
    - Only ONE parameter is required.

``` go run . updateTask --id="YOUR_TASK_ID" --title="New title" --goal="Updated goal" --goalStatus="4" --goalNote="Updated Note" ```

- Removing a task
    - Id is required.

``` go run . --id="YOUR_TASK_ID" ```
