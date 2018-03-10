using System.Collections.Generic;
using Microsoft.AspNetCore.Mvc;
using rooms_ms.Models;
using System.Linq;

namespace rooms_ms.Controllers
{
    [Route("api/room")]
    public class RoomController : Controller
    {
        private readonly RoomContext _context;

        public RoomController(RoomContext context)
        {
            _context = context;
/*
            if (_context.TodoItems.Count() == 0)
            {
                _context.TodoItems.Add(new TodoItem { Name = "Item1" });
                _context.SaveChanges();
            }
*/
        }

        [HttpGet]
        public IEnumerable<Room> GetAll()
        {
            return _context.Rooms.ToList();
        }

        [HttpGet("{id}", Name = "GetRoom")]
        public IActionResult GetById(long id)
        {
            var item = _context.Rooms.FirstOrDefault(t => t.Id == id);
            if (item == null)
            {
                return NotFound();
            }
            return new ObjectResult(item);
        }

        [HttpPost]
        public IActionResult Create([FromBody] Room item)
        {
            if (item == null)
            {
                return BadRequest();
            }

            _context.Rooms.Add(item);
            _context.SaveChanges();

            return CreatedAtRoute("GetRoom", new { id = item.Id }, item);
            //return new ObjectResult(_context.Rooms.FirstOrDefault(t => t.Id == item.Id));
        }

        [HttpPut("{id}")]
        public IActionResult Update(long id, [FromBody] Room item)
        {
            if (item == null || item.Id != id)
            {
                return BadRequest();
            }

            var todo = _context.Rooms.FirstOrDefault(t => t.Id == id);
            if (todo == null)
            {
                return NotFound();
            }

            todo.NameRoom = item.NameRoom;
            todo.IdOwner = item.IdOwner;
            todo.DescriptionRoom = item.DescriptionRoom;
            todo.participants = item.participants;

            _context.Rooms.Update(todo);
            _context.SaveChanges();
            return new NoContentResult();
        }

        [HttpDelete("{id}")]
        public IActionResult Delete(long id)
        {
            var todo = _context.Rooms.FirstOrDefault(t => t.Id == id);
            if (todo == null)
            {
                return NotFound();
            }

            _context.Rooms.Remove(todo);
            _context.SaveChanges();
            return new NoContentResult();
        }       
    }
}