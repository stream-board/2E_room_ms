using System.Collections.Generic;
using Microsoft.AspNetCore.Mvc;
using rooms_ms.Models;
using System.Linq;
using System;

namespace rooms_ms.Controllers
{
    [Route("room")]
    public class RoomController : Controller
    {
        private readonly RoomContext _context;

        public RoomController(RoomContext context)
        {
            _context = context;

            if (_context.Rooms.Count() == 0)
            {
                _context.Rooms.Add( new Room { Id = 0, NameRoom = "Default" , IdOwner = -1 , participants = new int[0]});
                _context.SaveChanges();
            }

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
            /*
            cuando el usuario va a crear una sala => solo se pasa el idOwner y el dueño, 
                no existe id sala o sea idsala = 0

            cuando el usuario es invitado y quiere unirse a sala => viene con un ID de sala:
                    esa sala existe?
                     si => añade a participants
                     no => badrequest            
            
            
             */
            if(item.Id == 0 ){//no le estoy pasando ningun id de sala
                //crear sala
                _context.Rooms.Add(item);
                _context.SaveChanges();            
                return CreatedAtRoute("GetRoom", new { id = item.Id }, item);
            }else{ //se le está pasando un IDSala
                //verificar si esa sala existe 
                var it = _context.Rooms.FirstOrDefault(t => t.Id == item.Id);
                //return new ObjectResult(it);
            
                if(it != null && it.IdOwner!=-1){ // es un invitado intentando acceder a una sala existente
                    //añado el usuario a la lista
                    int idPart = item.IdOwner;
                    int[] lista = {idPart};
                    int[] au = {idPart};
                    Room aux;
                    if(it.participants != null){
                        lista = it.participants.ToList().ToArray();
                        int len = lista.Length;
                        var z = new int[len+1];
                        lista.CopyTo(z,0);
                        au.CopyTo(z,len);
                        aux = new Room { participants = z };
                    }else{
                        aux = new Room { participants = lista };
                    }                    

                    it.participants = aux.participants;
                    aux.DescriptionRoom = "añadiendo uno a la lista";
                    _context.Rooms.Update(it);
                    _context.SaveChanges();                
                    
                    return new ObjectResult(aux);
                }else{//un invitado con un ID de sala no existente
                    return BadRequest(); //no hay sala a la que quiere acceder
                }
                 
            }
        }

        [HttpPut("{id}")]
        public IActionResult Update(long id, [FromBody] Room item)
        {
            /* 
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
            */
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