package geo

//
// @Author yfy2001
// @Date 2024/8/23 22 41
//

// Coordinate 坐标
type Coordinate struct {
	Longitude        float64 `json:"longitude"`        //经度
	Latitude         float64 `json:"latitude"`         //纬度
	LongitudeRadians float64 `json:"longitudeRadians"` //弧度制经度
	LatitudeRadians  float64 `json:"latitudeRadians"`  //弧度制纬度
}

func NewCoordinate(longitude, latitude float64) *Coordinate {
	return &Coordinate{
		Longitude:        longitude,
		Latitude:         latitude,
		LongitudeRadians: DegreeToRadians(longitude),
		LatitudeRadians:  DegreeToRadians(latitude),
	}
}

// AzimuthOffset 方位偏移
type AzimuthOffset struct {
	Azimuth  float64 `json:"azimuth"`  //方位角
	Distance float64 `json:"distance"` //偏移距离
}

func NewAzimuthOffset(azimuth, distance float64) *AzimuthOffset {
	return &AzimuthOffset{
		Azimuth:  azimuth,
		Distance: distance,
	}
}

// CoordinateOffset 坐标偏移
type CoordinateOffset struct {
	LongitudeOffset float64 `json:"longitudeOffset"` //经度偏移
	LatitudeOffset  float64 `json:"latitudeOffset"`  //纬度偏移
}

func NewCoordinateOffset(longitudeOffset, latitudeOffset float64) *CoordinateOffset {
	return &CoordinateOffset{
		LongitudeOffset: longitudeOffset,
		LatitudeOffset:  latitudeOffset,
	}
}
