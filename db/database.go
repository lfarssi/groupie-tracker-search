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
    Relations    string   `json:"relations"`
}


type Location struct {
    Locations []string `json:"locations"`
}

type Relation struct {
    Index []struct {
        Location []string `json:"location"`
        Date     []string `json:"date"`
    } `json:"index"`
}

type Dates struct {
    Dates []string `json:"dates"`
}
