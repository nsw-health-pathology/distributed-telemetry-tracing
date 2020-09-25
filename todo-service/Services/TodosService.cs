using TodoService.Models;
using TodoService.Repository;

namespace TodoService.Services
{

    public interface ITodosService
    {
        TodoItem GetTodo(string id);
        TodoItem CreateTodo(TodoItem newTodo);
        TodoItem UpdateTodo(string id, TodoItem updatedTodo);
        TodoItem DeleteTodo(string id);
    }

    public class TodosService : ITodosService
    {

        private readonly ITodoRepository _repo;
        public TodosService(ITodoRepository repo)
        {
            this._repo = repo;
        }

        public TodoItem CreateTodo(TodoItem newTodo)
        {
            return this._repo.CreateTodo(newTodo);
        }

        public TodoItem DeleteTodo(string id)
        {
            TodoItem existingTodo = this._repo.GetTodo(id);
            if (existingTodo == null) return null;

            return this._repo.DeleteTodo(id);
        }

        public TodoItem GetTodo(string id)
        {
            return this._repo.GetTodo(id);
        }

        public TodoItem UpdateTodo(string id, TodoItem updatedTodo)
        {
            TodoItem existingTodo = this._repo.GetTodo(id);
            if (existingTodo == null) return null;

            return this._repo.UpdateTodo(id, updatedTodo);
        }
    }
}