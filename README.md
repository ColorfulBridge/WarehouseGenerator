# WarehouseSimulatorService
## Overview
A multipart service that provides various convenience function to simulate a warehouse.
* Generation of a warehouse layout based on some input parameters
..* Generation of racks
..* Generation of storage bins
..* Generation of routes

## API

### API for Layout Generation 

API endpoint `/api/whlayout`

Input Message 
```json
{	
	"layout" : {
		"Horizontal" : false,
		"HLayout" : [1.075,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.075],
		"VLayout" : [1.1,0.35,1.1,0.35,1.1],
		"HDocks" : [0.05,1.9,0.05],
		"VDocks" : [1.1,0.2,1.1,0.2,1.1,0.2,1.1],
		"BinsPerRack" : 10
	},
	"outlineGeoJSON" : {
		"type": "Polygon",
		"coordinates": [
	    [
	      [
	        8.557853400707245,
	        49.664416151968375
	      ],
	      [
	        8.55984091758728,
	        49.66434497200099
	      ],
	      [
	        8.559948205947876,
	        49.66563487574149
	      ],
	      [
	        8.55795532464981,
	        49.66570605382166
	      ],
	      [
	        8.557853400707245,
	        49.664416151968375
	      ]
	    ]
	  ]
	}
	
}
```

