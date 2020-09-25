using System.Net;

namespace TodoService.Models
{
    public class Error
    {
        public string Message { get; }
        public int StatusCode { get; }

        public Error(string message, HttpStatusCode statusCode)
        {
            Message = message;
            StatusCode = (int)statusCode;
        }
    }
}