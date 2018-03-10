using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace rooms_ms.Models
{
    public class Room 
    {
        [Key][DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id {get; set;}
        [Required]
        public string NameRoom {get; set;}
        public string DescriptionRoom {get; set;}
        [Required]
        public int IdOwner {get; set;}
        public int[] participants {get; set;}
    }
}