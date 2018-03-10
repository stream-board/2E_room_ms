namespace rooms_ms.Models
{
    public class Room 
    {
        public int Id {get; set;}
        public string NameRoom {get; set;}
        public string DescriptionRoom {get; set;}
        public int IdOwner {get; set;}
        public int[] participants {get; set;}
    }
}