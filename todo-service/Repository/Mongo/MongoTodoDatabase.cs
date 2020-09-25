using System.Collections.Generic;
using MongoDB.Driver;
using TodoService.Models;

namespace TodoService.Repository.Mongo
{
    public interface ITodoDatabaseSettings
    {
        string ConnectionString { get; }
        string DatabaseName { get; }
        string TodoCollectionName { get; }
    }

    public class TodoDatabaseSettings : ITodoDatabaseSettings
    {
        public string ConnectionString { get; set; }
        public string DatabaseName { get; set; }
        public string TodoCollectionName { get; set; }
    }

    public class MongoTodoRepository : ITodoRepository
    {
        private readonly IMongoCollection<TodoItem> _todos;

        public MongoTodoRepository(ITodoDatabaseSettings settings)
        {
            MongoClient client = new MongoClient(settings.ConnectionString);
            IMongoDatabase db = client.GetDatabase(settings.DatabaseName);
            _todos = db.GetCollection<TodoItem>(settings.TodoCollectionName);
        }

        public TodoItem CreateTodo(TodoItem newTodo)
        {
            newTodo.Id = MongoDB.Bson.ObjectId.GenerateNewId().ToString();
            _todos.InsertOne(newTodo);
            return newTodo;
        }

        public TodoItem DeleteTodo(string id)
        {
            TodoItem todo = GetTodo(id);
            _todos.DeleteOne(t => t.Id == id);
            return todo;
        }

        public TodoItem GetTodo(string id)
        {
            TodoItem todo = _todos.Find<TodoItem>(t => t.Id == id).FirstOrDefault();
            return todo;
        }

        public TodoItem UpdateTodo(string id, TodoItem updatedTodo)
        {
            _todos.ReplaceOne(t => t.Id == id, updatedTodo);
            return updatedTodo;
        }
    }
}