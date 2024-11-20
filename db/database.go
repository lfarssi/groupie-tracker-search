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
    Rela Relation
    
}


type Location struct {
    Index []struct {
        Id       int    `json:"id"`
        Location []string `json:"locations"`
        Dates string `json:"dates"`
    } `json:"index"`
    Locations string `json:"locations"`
}

type Relation struct {
    Index         []struct {
        Id            int       `json:"id"`
        DatesLocations map[string]interface{} `json:"datesLocations"`
    } `json:"index"`
} 

type Date struct {
    Index []struct {
        Id int `json:"id"`
        Dates []string `json:"dates"`
    } `json:"index"`
}
