package wly
 
import "github.com/twpayne/go-geom"

func generateStorageBins(rack Rack, layout LayoutParameters) ([]StorageBin, string) {

	var bins []StorageBin = make([]StorageBin, 2 * layout.BinsPerRack)
	var binCollection = geom.NewGeometryCollection()

	var axis1 = 0;
	var axis2 = 1;

	if(!layout.Horizontal){
		axis1 = 1
		axis2 = 0
	}

	var sides = SplitPolygonEqually(&rack.Outline, axis1, 2)
	
	var racksA = SplitPolygonEqually(&sides[0], axis2, layout.BinsPerRack)
	var racksB = SplitPolygonEqually(&sides[1], axis2, layout.BinsPerRack)

	for i := 0; i < len(racksA); i++{

		var newBinA = StorageBin{}
		newBinA.Id = rack.Id + "-A-" + string(i*2)
		newBinA.Outline = racksA[i]
		newBinA.AsGeoJSON = marshallPolygon(&racksA[i])
		bins[i*2] = newBinA

		binCollection.Push(&racksA[i])
	}
	
	for i := 0; i < len(racksB); i++{
		var newBinB = StorageBin{}
		newBinB.Id = rack.Id + "-B-" + string(i*2+1)
		newBinB.Outline = racksA[i]
		newBinB.AsGeoJSON = marshallPolygon(&racksB[i])
		bins[i*2+1] = newBinB
		
		binCollection.Push(&racksB[i])
	}

	return bins, mustMarshallToGeoJSON(binCollection)

}
	  
