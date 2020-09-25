using System;
using System.Collections.Generic;
using TodoService.Models;

namespace TodoService.Repository.InMemory
{

    public class InMemoryTodoDatabase : ITodoRepository
    {
        private IDictionary<string, TodoItem> _db = new Dictionary<string, TodoItem>();

        public TodoItem CreateTodo(TodoItem newTodo)
        {
            newTodo.Id = Guid.NewGuid().ToString();
            this._db.Add(newTodo.Id, newTodo);
            return newTodo;
        }

        public TodoItem DeleteTodo(string id)
        {
            if (!_db.ContainsKey(id)) return null;
            TodoItem existingItem = _db[id];
            _db.Remove(id);
            return existingItem;
        }

        public TodoItem GetTodo(string id)
        {
            if (!_db.ContainsKey(id)) return null;
            return _db[id];
        }

        public TodoItem UpdateTodo(string id, TodoItem updatedTodo)
        {
            updatedTodo.Id = id;

            if (!_db.ContainsKey(id)) return null;
            _db[id] = updatedTodo;
            return updatedTodo;
        }
    }
}