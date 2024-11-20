package db

    type Artist struct {
        Id          int      `json:"id"`
        Name         string   `json:"name"`
        Members      []string `json:"members"`
        Image        string   `json:"image"`
        CreationDate int      `json:"creationDate"`
        FirstAlbum   string   `json:"firstAlbum"`
        Locations    string   `json:"locations"`     
        ConcertDates string   `json:"concertDates"`  
        Relations    string  `json:"relations"`
        Relation map[string][]string 
    }


type Location struct {
        Id       int    `json:"id"`
        Location []string `json:"locations"`
}

type Relation struct {
        Id            int       `json:"id"`
        DatesLocations map[string][]string `json:"datesLocations"`

} 

type Date struct {
        Id int `json:"id"`
        Dates []string `json:"dates"`
}
