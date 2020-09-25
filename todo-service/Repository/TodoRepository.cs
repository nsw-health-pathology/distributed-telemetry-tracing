using TodoService.Models;

namespace TodoService.Repository
{

    public interface ITodoRepository
    {
        TodoItem GetTodo(string id);
        TodoItem CreateTodo(TodoItem newTodo);
        TodoItem UpdateTodo(string id, TodoItem updatedTodo);
        TodoItem DeleteTodo(string id);
    }
}