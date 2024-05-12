package support

import (
	"strconv"

	"Adlet/core/backend/container"
)

func Assemble() {
	var mapping_all map[string]map[string][]int = make(map[string]map[string][]int)
	mapping_all["Location"] = make(map[string][]int)
	mapping_all["Member"] = make(map[string][]int)
	mapping_all["FirstAlbumDate"] = make(map[string][]int)
	mapping_all["Artist/band"] = make(map[string][]int)
	mapping_all["CreationDate"] = make(map[string][]int)

	for i := 0; i < len(container.Mainobject.Location.Index); i++ {
		for j := 0; j < len(container.Mainobject.Location.Index[i].Locations); j++ {
			mapping_all["Location"][container.Mainobject.Location.Index[i].Locations[j]] = append(mapping_all["Location"][container.Mainobject.Location.Index[i].Locations[j]], i)
		}
	}

	for i := 0; i < len(container.Mainobject.Artist.Index); i++ {
		for j := 0; j < len(container.Mainobject.Artist.Index[i].Members); j++ {
			mapping_all["Member"][container.Mainobject.Artist.Index[i].Members[j]] = append(mapping_all["Member"][container.Mainobject.Artist.Index[i].Members[j]], i)
		}
	}

	for i := 0; i < len(container.Mainobject.Artist.Index); i++ {
		mapping_all["FirstAlbumDate"][container.Mainobject.Artist.Index[i].FirstAlbum] = append(mapping_all["FirstAlbumDate"][container.Mainobject.Artist.Index[i].FirstAlbum], i)
		mapping_all["Artist/band"][container.Mainobject.Artist.Index[i].Name] = append(mapping_all["Artist/band"][container.Mainobject.Artist.Index[i].Name], i)
		mapping_all["CreationDate"][strconv.Itoa(container.Mainobject.Artist.Index[i].CreationDate)] = append(mapping_all["CreationDate"][strconv.Itoa(container.Mainobject.Artist.Index[i].CreationDate)], i)

	}


	container.MapAll = mapping_all
}
