package container

type Instance struct {
	Artist   *Artiststruct
	Location *Locationstruct
	Dates    *Datestruct
	Relation *Relationstruct
}

type Artiststruct struct {
	Index []Subartist
}

type Subartist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
}

type Relationstruct struct {
	Index []Subrelation
}

type Subrelation struct {
	Id             int
	DatesLocations map[string][]string
}

type Datestruct struct {
	Index []Subdate
}
type Subdate struct {
	Id    int
	Dates []string
}

type Locationstruct struct {
	Index []Sublocation
}

type Sublocation struct {
	Id        int
	Locations []string
}

var (
	Mainobject Instance = Instance{}
	MapAll     map[string]map[string][]int
	Array_Available []int 
)
