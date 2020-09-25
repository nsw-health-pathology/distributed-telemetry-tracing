using System.Collections.Generic;
using System.Net;
using Microsoft.AspNetCore.Mvc;
using TodoService.Models;
using TodoService.Services;

namespace TodoService.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class TodosController : ControllerBase
    {
        private readonly ITodosService _svc;
        public TodosController(ITodosService svc)
        {
            this._svc = svc;
        }

        [HttpGet("{id}")]
        public ActionResult<IEnumerable<TodoItem>> Get(string id)
        {
            TodoItem t = _svc.GetTodo(id);
            if (t == null)
            {
                var error = new Error("Todo Not found", HttpStatusCode.NotFound);
                return new JsonResult(error) { StatusCode = error.StatusCode };
            }
            return Ok(t);
        }

        [HttpPost]
        public ActionResult<IEnumerable<TodoItem>> Post(TodoItem newTodo)
        {
            TodoItem t = _svc.CreateTodo(newTodo);
            return Created(t.Id, t);
        }

        [HttpPut("{id}")]
        public ActionResult<IEnumerable<TodoItem>> Update(string id, TodoItem updatedTodo)
        {
            TodoItem t = _svc.UpdateTodo(id, updatedTodo);
            if (t == null)
            {
                var error = new Error("Todo Not found", HttpStatusCode.NotFound);
                return new JsonResult(error) { StatusCode = error.StatusCode };
            }
            return Ok(t);
        }

        [HttpDelete("{id}")]
        public ActionResult<IEnumerable<TodoItem>> Delete(string id)
        {
            TodoItem t = _svc.DeleteTodo(id);
            if (t == null)
            {
                var error = new Error("Todo Not found", HttpStatusCode.NotFound);
                return new JsonResult(error) { StatusCode = error.StatusCode };
            }
            return Ok(t);
        }
    }
}